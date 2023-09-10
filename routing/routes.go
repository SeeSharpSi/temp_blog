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
	mux.HandleFunc("/form", GetForm)
	mux.HandleFunc("/form_html", GetForm_html)
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

func GetForm_html(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /form_html request\n")
	if r.Header.Get("Hx-Request") == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.ServeFile(w, r, "tmpl/form.html")
	}
}

func GetForm(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /form request\n")
	if r.Header.Get("Hx-Request") == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		fname := r.URL.Query().Get("fname")
		lname := r.URL.Query().Get("lname")
		fmt.Printf("\nfname => %s\nlname => %s\n", fname, lname)
		_, err := w.Write([]byte("<button hx-get='/form_html' hx-swap='outerHTML'>New Form</button>"))
		check(err)
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

func ServeStatic(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /static request\n")
	io.WriteString(w, `<div hx-put="/hello" hx-swap="outerHTML">hello</div>`)
}
