package routes

import (
	"blog/main/tmpl"
	"context"
	"fmt"
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
    mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("/posts", GetPosts)
	return mux
}

func GetIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
    component := templ.Index()
    component.Render(context.Background(), w)
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("got /posts request\n")
    temp := []string{"one", "two", "three"}
    component := templ.Posts(temp)
    component.Render(context.Background(), w)
}

func ServeStatic(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /static request\n")
}
