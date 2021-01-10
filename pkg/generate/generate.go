package generate

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"strings"
	"text/template"

	"github.com/georgemac/truce"
	"github.com/georgemac/truce/pkg/http"
)

type Generator struct {
	api truce.API

	context context
}

type context struct {
	Package    string
	Imports    string
	ClientName string
	ServerName string
	Types      map[string]truce.Type
	Bindings   http.Bindings
}

func New(api truce.API, opts ...Option) (Generator, error) {
	bindings, err := http.BindingsFrom(api)
	if err != nil {
		return Generator{}, err
	}

	g := Generator{
		api: api,
		context: context{
			Package:    "types",
			ClientName: "Client",
			ServerName: "Server",
			Types:      api.Types,
			Bindings:   bindings,
		},
	}

	Options(opts).Apply(&g)

	return g, nil
}

func (g Generator) WriteTypesTo(w io.Writer, opts ...Option) error {
	Options(opts).Apply(&g)
	return writeGo(w, typeTmpl, g.context)
}

func (g Generator) WriteClientTo(w io.Writer, opts ...Option) error {
	Options(opts).Apply(&g)
	return writeGo(w, clientTmpl, g.context)
}

func (g Generator) WriteServerTo(w io.Writer, opts ...Option) error {
	Options(opts).Apply(&g)
	return writeGo(w, serverTmpl, g.context)
}

func writeGo(w io.Writer, tmpl *template.Template, context interface{}) error {
	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, context); err != nil {
		return err
	}

	src, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	_, err = w.Write(src)
	return err
}

var tmplFuncs = template.FuncMap{
	"name": func(f truce.Field) (v string) {
		parts := strings.Split(f.Name, "_")
		for _, p := range parts {
			v += strings.Title(p)
		}
		return
	},
	"tags": func(f truce.Field) (v string) {
		return fmt.Sprintf("`json:%q`", f.Name)
	},
	"signature": func(f truce.Function) string {
		builder := &strings.Builder{}
		fmt.Fprintf(builder, "%s(ctxt context.Context, ", f.Name)
		for i, arg := range f.Arguments {
			if i > 0 {
				builder.Write([]byte(", "))
			}

			fmt.Fprintf(builder, "v%d %s", i, arg.Type)
		}

		builder.Write([]byte(") ("))
		if rtn := f.Return; rtn.Name != "" {
			fmt.Fprintf(builder, "rtn %s, ", rtn.Type)
		}

		builder.Write([]byte("err error)"))

		return builder.String()
	},
	"path": func(b http.Binding) string {
		if len(b.Path) == 0 {
			return fmt.Sprintf("%q", b.Path)
		}

		return fmt.Sprintf(`fmt.Sprintf(%q, %s)`, b.Path.FmtString(), b.Path.ArgString())
	},
	"args": func(f truce.Function) (v string) {
		for i := range f.Arguments {
			if i > 0 {
				v += ", "
			}

			v += fmt.Sprintf("v%d", i)
		}
		return
	},
	"method": func(b http.Binding) string {
		return strings.Title(strings.ToLower(b.Method))
	},
}

var typeTmpl = template.Must(
	template.
		New("types").
		Funcs(tmplFuncs).
		Parse(`package {{ .Package }}

import (
{{ .Imports }}
)

{{ range $type := .Types }}
type {{ $type.Name }} struct {
{{ range $field := $type.Fields }}  {{name .}}   {{.Type}} {{tags .}}
{{ end }}
}
{{ end }}`),
)

var clientTmpl = template.Must(
	template.
		New("client").
		Funcs(tmplFuncs).
		Parse(`package {{ .Package }}

import (
"fmt"
"io"
"io/ioutil"
"net/http"
"net/url"
"encoding/json"
"context"
{{ .Imports }}
)

type {{.ClientName}} struct {
    client *http.Client
    host *url.URL
}

func New{{.ClientName}}(host string) (*{{.ClientName}}, error) {
    u, err := url.Parse(host)
    if err != nil {
        return nil, err
    }

    return &{{.ClientName}}{client: http.DefaultClient, host: u}, nil
}

{{ $ctxt := . }}{{ range .Bindings }}
func (c *{{ $ctxt.ClientName }}) {{signature .Function}} {
    u, err := c.host.Parse({{path .}})
    if err != nil {
        return
    }

    {{if ne (len .Query) 0}}query := u.Query()
    {{range $k, $v := .Query}}query.Set("{{$k}}", {{$v}})
    {{end}}u.RawQuery = query.Encode(){{end}}

    var (
        body io.Reader
        resp *http.Response
    )

    {{if ne .BodyVar ""}}buf := &bytes.Buffer{}
    body = buf
    if err = json.NewEncoder(buf).Encode({{.BodyVar}}); err != nil {
        return
    }{{end}}

    req, err := http.NewRequest("{{.Method}}", u.String(), body)
    if err != nil {
        return
    }

    resp, err = c.client.Do(req.WithContext(ctxt))
    if err != nil {
        return
    }

    defer func(){
        _, _ = io.Copy(ioutil.Discard, resp.Body)
        _ = resp.Body.Close()
    }()

    {{if .HasReturn}}{{if .ReturnIsPtr}}rtn = &{{.ReturnType}}{}{{end}}
    err = json.NewDecoder(resp.Body).Decode({{if not .ReturnIsPtr}}&{{end}}rtn){{end}}

    return
}
{{ end }}`),
)

var serverTmpl = template.Must(
	template.
		New("server").
		Funcs(tmplFuncs).
		Parse(`package {{ .Package }}

import (
"context"
"encoding/json"
"net/http"

"github.com/go-chi/chi"
{{ .Imports }}
)

type Service interface {
    {{range .Bindings}}{{signature .Function}}
    {{end}}
}

type {{.ServerName}} struct {
    chi.Router
    srv Service
}
{{ $ctxt := . }}
func New{{.ServerName}}(srv Service) *{{.ServerName}} {
    s := &{{.ServerName}}{
      Router: chi.NewRouter(),
      srv: srv,
    }

    {{ range .Bindings }}s.Router.{{method .}}("{{.Path}}", s.handle{{.Function.Name}})
    {{end}}

    return s
}

{{ range .Bindings }}
func (c *{{ $ctxt.ServerName }}) handle{{.Function.Name}}(w http.ResponseWriter, r *http.Request) {
    {{range .Path}}{{if eq .Type "variable"}}{{.Var}} := chi.URLParam(r, "{{.Value}}")
    {{end}}{{end}}
    {{range $k, $v := .Query}}{{$v}} := r.URL.Query().Get("{{$k}}")
    {{end}}
    {{if ne .BodyVar ""}}var {{.BodyVar}} {{.BodyType}}
    if err := json.NewDecoder(r.Body).Decode(&{{.BodyVar}}); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    {{end}}
	r0, err := c.srv.{{.Function.Name}}(r.Context(), {{args .Function}})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(r0); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return
}
{{ end }}`),
)
