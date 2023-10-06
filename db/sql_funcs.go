package sql_funcs

import (
	"database/sql"
	"fmt"
    types "blog/main/types"

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
        temp_post := types.Post{
        }
        result.Scan(&temp_post.Id, &temp_post.Title, &temp_post.Teaser, &temp_post.Content)
        posts = append(posts, temp_post)
    }
    return posts
}
