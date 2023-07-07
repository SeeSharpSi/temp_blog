package routes

import (
	"fmt"
	"io"
	"net/http"
	"text/template"
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
	mux.HandleFunc("/home", GetHome)
	mux.HandleFunc("/posts", GetPosts)
	mux.HandleFunc("/pictures", GetPictures)
	return mux
}

func GetIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	http.ServeFile(w, r, "tmpl/index.html")
}

func GetHome(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /home request\n")
	if r.Header.Get("Hx-Request") == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.ServeFile(w, r, "tmpl/home.html")
	}
}

func GetPictures(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /pictures request\n")
	if r.Header.Get("Hx-Request") == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.ServeFile(w, r, "tmpl/pictures.html")
	}
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /posts request\n")
	tmpl, err := template.ParseFiles("tmpl/posts.html")
	check(err)
	Posts := []Post{
		{
			Title:   "title1",
			Content: "content1",
		},
		{
			Title:   "title2",
			Content: "content2",
		},
		{
			Title:   "title3",
			Content: "content3",
		},
	}
	if r.Header.Get("Hx-Request") == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		tmpl.Execute(w, Posts)
	}
}

func ServeStatic(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /static request\n")
	io.WriteString(w, `<div hx-put="/hello" hx-swap="outerHTML">hello</div>`)
}
