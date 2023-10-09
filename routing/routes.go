package routes

import (
	"blog/main/tmpl"
	types "blog/main/types"
    sql_funcs "blog/main/db"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

func StartHandlers() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Get("/", GetIndex)
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.Get("/posts", GetPosts)
	r.Get("/home", GetHome)
	r.Get("/post/{postId}", GetPost)
	return r
}

func GetIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	http.Redirect(w, r, "/home", http.StatusMovedPermanently)
}

func GetHome(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /home request\n")
	header := r.Header
	if header.Clone().Get("Hx-Request") == "true" {
		component := templ.Home("Silas")
		component.Render(context.Background(), w)
	} else {
		component := templ.Index()
		component.Render(context.Background(), w)
		component = templ.Home("Silas")
		component.Render(context.Background(), w)
	}
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /posts request\n")
    var posts []types.Post
    posts = sql_funcs.Get_Posts()
	header := r.Header
	if header.Clone().Get("Hx-Request") == "true" {
		component := templ.Posts(posts)
		component.Render(context.Background(), w)
	} else {
		component := templ.Index()
		component.Render(context.Background(), w)
		component = templ.Posts(posts)
		component.Render(context.Background(), w)
	}
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	postId := chi.URLParam(r, "postId")
	header := r.Header
	intPostId, err := strconv.Atoi(postId)
    post := sql_funcs.Get_Post(intPostId)
	if header.Clone().Get("Hx-Request") == "true" {
		component := templ.ExpandedPost(post)
		component.Render(context.Background(), w)
	} else {
		component := templ.Index()
		component.Render(context.Background(), w)
		component = templ.ExpandedPost(post)
		component.Render(context.Background(), w)
	}
	check(err)
}
