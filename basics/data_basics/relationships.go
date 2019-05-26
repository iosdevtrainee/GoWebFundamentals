package data 

import (
  "database/sql"
  "errors"
  "fmt"
  _ "github.com/lib/pq"
)

type Post struct {
  Id     int 
  Content string 
  Author  string 
  Comments []Comment
}

type Comment struct { 
  Id     int 
  Content string 
  Author string 
  Post   *Post 
}

var Db *sql.DB

func init() {
  var err error
  Db, err = sql.Open("postgres", "user=postgres dbname=postgres password=testdb sslmode=disable")
  if err != nil {
    panic(err)
  }
}

func (comment *Comment) Create() (err error) {
  if comment.Post == nil {
    err = errors.New("No Post associated with this comment")
    return 
  }
  err = Db.QueryRow("insert into comments (content, author, post_id) values ($1, $2, $3) returning id", comment.Content, comment.Author, comment.Post.Id).Scan(&comment.Id)
  return 
}

func GetPost(id int) (post Post, err error) {
	var post = Post{}
	rows, err := Db.QueryRow("select id, content, author from posts where id = $1, id").Scan(&post.Id, &post.Content, &post.Author)
	return
}

func (post *Post) Create() (err error) {
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

func main() {
  post := Post{Content: "Hello World!", Author: "Sau Sheong"}
  post.Create()
  comment := Comment{Content: "Good post!", Author: "Joe", Post: &post}
  comment.Create()
  readPost, _ := GetPost(post.Id)
  fmt.Println(readPost)
  fmt.Println(readPost.Comments)
  fmt.Println(readPost.Comments[0].Post) 
}
