package example

import (
	"fmt"
)

type NotAuthorized struct {
	Message string `json:"message"`
}

func (e NotAuthorized) Error() string {
	return fmt.Sprintf("error: message=%q", e.Message)
}

type NotFound struct {
	Message string `json:"message"`
}

func (e NotFound) Error() string {
	return fmt.Sprintf("error: message=%q", e.Message)
}

type PatchPostRequest struct {
	Body  []byte `json:"body"`
	Draft bool   `json:"draft"`
	Title string `json:"title"`
}

type PatchUserRequest struct {
	Age    int     `json:"age"`
	Height float64 `json:"height"`
	Name   string  `json:"name"`
}

type Post struct {
	Body  []byte `json:"body"`
	Draft bool   `json:"draft"`
	Id    string `json:"id"`
	Title string `json:"title"`
}

type PutPostRequest struct {
	Body  []byte `json:"body"`
	Draft bool   `json:"draft"`
	Title string `json:"title"`
}

type PutUserRequest struct {
	Age    int     `json:"age"`
	Height float64 `json:"height"`
	Name   string  `json:"name"`
}

type User struct {
	Age    int     `json:"age"`
	Height float64 `json:"height"`
	Id     string  `json:"id"`
	Name   string  `json:"name"`
}
