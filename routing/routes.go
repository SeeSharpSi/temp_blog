package routes

import (
	"blog/main/tmpl"
	types "blog/main/types"
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
	posts := []types.Post{
		{
			Title:   "title one",
			Content: "content one",
            Id: 1,
		},
		{
			Title:   "title two",
			Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
            Id: 2,
		},
		{
			Title:   "title three",
			Content: "content three",
            Id: 3,
		},
	}
	header := r.Header
	fmt.Println(header)
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
    intPostId, err := strconv.Atoi(postId)
    check(err)
    component := templ.ExpandedPost(intPostId)
    component.Render(context.Background(), w)
}
