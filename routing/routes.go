package routes

import (
	"blog/main/tmpl"
	"context"
	"fmt"
	"io"
	"net/http"
)

type Post struct {
	Title   string
	Content string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func StartHandlers() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", GetIndex)
	mux.HandleFunc("/test", GetTest)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	return mux
}

func GetIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
    http.ServeFile(w, r, "tmpl/index.html")
}

func GetTest(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("got /test request\n")
    component := templ.Test("working")
    component.Render(context.Background(), w)
}

func ServeStatic(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /static request\n")
	io.WriteString(w, `<div hx-put="/hello" hx-swap="outerHTML">hello</div>`)
}
