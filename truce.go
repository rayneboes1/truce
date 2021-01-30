// Code generated by gocode.Generate; DO NOT EDIT.

package truce

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/encoding/gocode/gocodec"
)

var cuegenCodec, cuegenInstance = func() (*gocodec.Codec, *cue.Instance) {
	var r *cue.Runtime
	r = Runtime
	instances, err := r.Unmarshal(cuegenInstanceData)
	if err != nil {
		panic(err)
	}
	if len(instances) != 1 {
		panic("expected encoding of exactly one instance")
	}
	return gocodec.New(r, nil), instances[0]
}()

// cuegenMake is called in the init phase to initialize CUE values for
// validation functions.
func cuegenMake(name string, x interface{}) cue.Value {
	//lint:ignore SA1019 until FieldByName is produced by cue gocode https://github.com/cuelang/cue/pull/664
	f, err := cuegenInstance.LookupField(name)
	if err != nil {
		panic(fmt.Errorf("could not find type %q in instance", name))
	}
	v := f.Value
	if x != nil {
		w, err := cuegenCodec.ExtractType(x)
		if err != nil {
			panic(err)
		}
		v = v.Unify(w)
	}
	return v
}

// Data size: 2039 bytes.
var cuegenInstanceData = []byte("\x01\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\xbcY\xefn\u0738\x11_\xe5R\xa0\x12\xae\xfd\xde\x03\n\xd0\xdc \xd8u\u05b2\x1b\xf7\x0f*\xd4glm\xe7b\xe0\x1a/\x92\xed}\xa8\xac\x04\xf4.w\u035cD\xe9(\xca\x17'\xd9\x03\xda^\xaf}\xaa>T\x1f\xe0\\\fII\x94v\u05ce\x1b\xa3\xf9b\xeep\xe67?\xce\f\xc9\x11\xf3\xb3\xab\u007f\xdes\xee]\xfd\xab\xe3\\\xfd\xad\xd3\xf9\xdd\xd5_?q\x9cO\x19\xcf%\xe1\x13zH$\x01\xb9\xf3\x89s\xffy\x9aJ\xe7^\u01f9?\"\xf2\xdc\xf9\xb4\xe3\xfc\xe4\t\x8bi\xee\\\xfd\xd0\xe9t~y\xf5\x8f{\x8e\xf3\xf30\x9a\x14\u051f\xb1\xd8X\xfe\xd0q\xae\xbe\xeftzW\u007f\xff\xc4q~Z\u02ff\xef8\xf7\x9c\xfb\xcfHB\x01\xe8\xbe\x12z\x9dN\xe7\xc7\xcf\xfe\x03L\x1c\xc7\xf9\u015c\xc9\xf3\xe2\u031f\xa4\xc9\xf6X\x14\x13\xfa|t\xb0-a\xe08\x8e[\xf0\x84\x88\xfc\x9c\xc4\u038f\x9f\xfd;#\x93\xaf\u025c\"5\xeby,\xc9R!Q\xcfs\xb1\xa0s\xfa&\u00de\x8bs)\x18\x9f\xe7\xd8\xeb{\u07abn>9\xa7\t99{\x1d\xa0w\x9e\x1b\xea\xc9H\xfdp_\xc9\u02cc\x06\b!\xa4\xc5 b\xf9\x9fH\x0625\x89\xf6\xbeC\xf8eB\xb20\x8cz\xe1\xcb(\xeao\x86Q\xd4\xf37\xfb\x0f\xb0\xd6\x1e\nA.\x03t\x96\xa6\xb1\x16\x8c\x04K\x02T\t\xceH\xdev1KEBd`\x89\xd8\u030cs\u007f,X2\x12t\xc6\xde\xf4V\x88\xf0iO\xf1\xea\xe3\x01\xc2a\x84\xfb\x03\x847q\x1f\xed\xed!\u0338\xc4jY\xb5S\x90\xd19\x15XIK\xbf \xfd\xed\xafA\xb6\xb8+\u07f38%\x80\xd9\xf6\u03cb\xe4l\u027dR\xbeS\xf7\x10\xec%\xdf \xa4\x84\xb7\x9c\u07e9_\xc9\x12\x9aK\x92dK\xce5P\xcb\xf7\x94H\xba\x056wFb\xa3\xcc\xfb\u00c7w\x80\xa4\xc3x'PUA\xdc\t\x9a\x15\xe7\xbbY\u8964\xed\x94}\x14\xea\xda\x1a3\xa7\xc8\x1e(\xdb^\xab\x83cF\xe2\x9c\x1a\x919:\xa4(h\x83\xda\xearB\x18\xe6\x1b\x1bY{\xdbX\xef\xcd:\xd4\xc20\n\xa3\xc8\xdf|\x80\x9b\xde-\x95^\x18\x81N\u007f\xbf\x17nF\xfd\xfd\x90l\xbd-\rJ\x87\xea\xb84^\x12\x92\x8d\x88\x90y\x80\x90>\x8d\xfd'\x8cO\x87q\xfc\xa28K\x88\x9c\x9c\xf7\xd6\x1d\xa5\x03\xedu\x80v\xfb\x86\xccW$\xd6|*\xd4\xf0qt#+\xb7:\xd0!h\xe9\xd9k:Q\a\x8dK\xa6S&Y\xca\x015\u0368\x90\x8c\xe6\xfa\x0e(\x97a\x1c\x1a\x99\x01\xb2\x9d+\xf9\xa2\xb4\xd8X6y \xe8,@\xb8\xbb=I\x93,\xe5\x94\xcb|[\xdf>\xf9\xf6i\u03c6\xea\xe3\x1alQ\as\u00ce\xa6f\xa5\xd2f\x1chJ\x98\x80H\x030I\x93j\x19\xc6\xc2fT\xae\x02\xeaDK\x16\x95\xeaF[\xf7\x06\xfa\x80ax\x1b\x94\x8a~\tg\x93]&\xb3\xc4\x05Tt-\xab\x8a\xc5\x15\x93\xb2\xc0\xcd\xec\x92\xcfU\xf4?\x9c\xbd\x1d\xf5\x85\xb7\xf0\u048cr\x92\xb1]\x15\xc7Y*\x10\xc9\x18t+\x03\x18|EE\xceR\x9e#\xc6Q\x9e\xd1\t\x9b\xb1\t\x91J\x02\x8e\x8d\xba\xd1R\x16\x87t\x06\u02b6\xad\xa2\x88O{\x06\xb8\x8f\u02d4i\x99\u046b\u016e\xa1\x14 \xbc\xeb\xef\xf8\xbb&\xea\x8c\xcf\u04a0N-\x93\xb1*u\x83j\xc4\x17\x1a,\xb0\b42_\x87\xa7\x862q\xaa\x05j]\x90/\x1d\a\x18\xd5\xcb:\xa43\x1f$y\xad\x0e\v)\u056deXEkmE\xfd/[\u0686\xb5\xe7\x19\xa3\xf1T\xbbVC\xe3\xdb\xd0\xf0\x95,oX\x01\x81\xca\n\x18\xf4\xac\xc6\x0f=l\xea\x9a\xf5\x06-i\xd9\r\x96.\xd5\"\x1b\x1a\v\xfb\u05e2\xefk\x1co\xa5B=\xacF\xe5`\xd1HHF\xe4\xb9\x15\x03\xb5~n\x16\u03dbQ\x9f\x15|R\xd7^\xb5n\xae\xc8\n\xc2sh\x86s\xff\\\xca\xcc\a\xd8f*\xf0i}\x9f\xa5_\xa6\xdfR\xb1\xc64\xa1\xf2<\x9d\xf6[\x89\x84t\xa9\xca?\x9e\x06\x86a=\xf9*#\x82$TR\x91\a(Dj[\x88\xb9a\xaf\xbd\x101/\x12\xa8\xbbF\xd8\xd9\f\xad&Q\xa9\x87\x1a\xc8\xe7$\xa1\x11\x1c\x14\xbc\x88c\xe8\x02nc\xe7\xcfD\x9a\x98S\xe6\u007f6=K\xa7\x97\xb8Y4\xa0\xa1\xf6`m\xe0\xad(\xb4\x8f\xa8F\x83|\xfbZt\x19//B\x84n\xbfb\x1biJ\xf3\x89`\x99T\xe7\n\x9cY\xb5r\x1f\xa9\xf0\xac\xab\xc2k\x1c\xf4q\x83\xec\xed\xea@'\x05\xba*\xa8\xf3VR\\A\xbf)\x98\xa0\u04e0\xea\xa4VE\xcd\u07ac(\xb2j\x99\xa7\xfc(\xc9\xe4\xe5h\xa9\xa63(g\xab\xd6\x11\x9b\xa1\x98\xf2^\xd6G\x9f\xa3\x1d\xf4\x0ee\xa8\t\x05%\xb3|\xbe}\xd8\xd6\xf8\xff\ue37d\xd5\x05\xee^W\xb9\xebJ\xf7\xa6\xdam\x16o\xab\x94o:N]A\xf3,\xe5y\xfb\xe2\xc0\x8fwvp\x8bJ\xbbpuXt\xddV\xe7\x19z\xbc\xb3\x83J\xd0vQnh\x13Ae!\xb8\xcf\xd3\xe7j\xd0Z\xf0$\xe5\x92r\xb9\x14\aL\xb2,6\xfd\xc2\xf6\xeb<\xe5m~7\xc5w}\x84\xeb\xdb\xca\xe6\xb7\x1c\xe9v\xac\x97~\u007fT*L\xf9\xaf\xd81fCX\b\xf6\r\xb1\xc2\xe0Z|(\xcceD\xd8\xe64\x97\u007f\\\xdac\x1f\x98wc\x8fT\xd9\xdb\xe6k\xf2\xb9*\x9d\x8a\u06ad\x82v};\xd0nN\x9b=\xa7~\xc2\xe2{\x8dG\xac\xf0\xa2\xfe\xdd\x1d\x8e\x8e\xcb2*\xef&\xa4\x1a\xc0\xaa/\xbc\xb0:\xdfBf\x85\xbc\x19\xb5\tW\xbfZ\x95}\xea~\x80\xba'\x19\xe5\xc3\xd1\xf1\x89\x82l8D{\xdf\xe1\x97\xe1\xce\xd6\xef\xa3G\xfa\x1bm\x9e\ue6db\xa9\xfbEZ\x19,>\x9cwe\xa6\xa8\xa9^\x14\x10\xbbc\"\xe6TznN\xc5\x05\x15\xc0j|\x99\xd1i)\x9e\u010cr\xd9\x16/\xbc&w\x85\t\xb7\x8a\xf5h\x87\u07a3M\xeco\xe7\u07d2\xf9\x9c\n_\xe5\u07ab\xa9\xed\x02\x88\x86\xb3\xac-\xd3?\xe4r\x9a\x16\xf2s\xec\xb9\xd9\xd7\xf3\xa0\x01\xab\xe8c\x85P\xb3\xaa\x16\u0580y\xa1\u0585=\xb7[s\x1f\x8e\x8e\x95r\x15\xb52;\x15\xbbf\xf8\x95;\x93b+\xe5\xcay3\x03\xbc\f\xb8[\xf5\x9d+\uc798\xb9u\xb6\xf5\x1d\xb4\xaf\xad\xe1&\x82\x1c<\x1d\x8fG&\x9d0T\x93\x863\u0737\xbe\xef\x1b'\x9e\x9b\xa9\xa7\x97\xa0Z\x1a\x15\"\x15%\x99\x89E\x06\x80\x8e`\xb2d\x93K\"\x8b\xfc \x9d\xd2\x00M\xac\x02*Y[\xa1Sa\x1an\xfd%\n\xc9\xd6[\xf8\xab^\x14\xae\xa5_\xa2\xa8\x85V\x97\xab&\xdf}\x02\x9f\x11\x91\xe7\xea\xa39@\x9b`]^ eg\xb2@\xefQW\x8b\x94~Y\x067\xf3\u049fU*\x12X\x15G.E\xa1\xbe\xb0\xf4\xe7\u046aT\xc1\u011a<\x99,\xa8\xe0)S;r\x8a\u00af\xb6~\x13\xa92z\xf7xQ3@\xf6\x89\x00\x81\x05\x1fm\xf6d\xeb\xed*\xf6\xa9\xe8uI\x1c\xf7\xd1{P\xd3\x0fC\xe1f\xb4\x0f\xab\x05\x93Ga\xd4\xfc\x8d\xb5\xe6\xcbRh\x83\x96S\xeaQkYa_\xed0+\u049ac\x95\x8e\xb34\x8d!\x8c\xe65\xae[e\xc3\xces\xe3h(\x8bQ\u007f-\x05\b\u007fq4\x06\x16xt\xf2\xc2\f\xfel\xfe\x0e\xc7\aO\xd5\xe8d4>>y\xf6B\x8d\x0f\x8f\xbe<\x1a\x1f\xa9\xe1\u04e3\xe1\xa1\x1a\x1c\x9c<{vt\xa0\xad\xc6\u03c7\aG\xb8QWK\x19\x1d\x9a\xb9\xf5I\xcd\x04K\x98d\x17\xb0\xe7C\xfd\xbe;\xd0/\u0183\xfa\x8dv`\x1eD\a\xd5;\xe3\xa0\xfab\x1f\u060f\xaf\x91\xd7\xcdc6\x010dZ\xe27\xd0\xc6Zn\xd0;\x84\xc3\xe8\xb4\xf7\xa6\x8fU/\xdc\xcdR\xc6\xed\x1ez\x95\xc1\xa6\xa5O\u2e3c\x1d,\xa5G\xc8x\x86Q\t\xe9U\x01\xb0*\xae\xcc\vt\xb5\x81\xe9i=\xd8g\xeb4\xd4\x17\x84\xe7^\x10Q\xff\x8f\xccu\xfa\xdf\x14T\\^kP\x9f\xc1$.j\x80\x85\xd7\xe9\xfc7\x00\x00\xff\xff\x01\x80\xad[\x82\x1b\x00\x00")
