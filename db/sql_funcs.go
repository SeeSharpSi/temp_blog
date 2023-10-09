package sql_funcs

import (
	types "blog/main/types"
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func connect_db() *sql.DB {
	db, err := sql.Open("sqlite3", "./posts.db")
	check(err)
	return db
}

func Add_Post(title string, teaser string, content string) {
	db := connect_db()
	defer db.Close()
	title = strings.Replace(title, "'", "''", -1)
	teaser = strings.Replace(teaser, "'", "''", -1)
	content = strings.Replace(content, "'", "''", -1)
	sqlStmt := fmt.Sprintf("insert into post (title, teaser, content) values ('%s', '%s', '%s');", title, teaser, content)
	_, err := db.Exec(sqlStmt)
	check(err)
}

func Get_Posts() []types.Post {
	db := connect_db()
	defer db.Close()
	sqlStmt := fmt.Sprintf("select * from post;")
	result, err := db.Query(sqlStmt)
	defer result.Close()
	check(err)
	posts := []types.Post{}
	for result.Next() {
		temp_post := types.Post{}
		result.Scan(&temp_post.Id, &temp_post.Title, &temp_post.Teaser, &temp_post.Content)
		posts = append(posts, temp_post)
	}
	return posts
}

func Get_Post(id int) types.Post {
	var response types.Post
	db := connect_db()
	defer db.Close()
	sqlStmt := fmt.Sprintf("select * from post where post_id = %d", id)
	result, err := db.Query(sqlStmt)
	fmt.Printf("%+v\n", result)
	defer result.Close()
	check(err)
	result.Next()
	if _, err := result.Columns(); err != nil {
		return types.Post{
			Title:   fmt.Sprintf("Post %d does not exist", id),
			Content: "<h1>error 404</h1>",
			Id:      id,
		}
	}
	result.Scan(&response.Id, &response.Title, &response.Teaser, &response.Content)
	return response
}
