package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var route = http.NewServeMux()
var db, _ = sql.Open("mysql", "root:12345@tcp(localhost:33066)/go-mysql-crud?charset=utf8")

type Post struct {
	id         string `json: "id"`
	title      string `json: "title"`
	content    string `json: "content"`
	created_at string `json: "created_at"`
}

// Logger return log message
func Logger() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now(), r.Method, r.URL)
		route.ServeHTTP(w, r) // dispatch the request
	})
}

// server starting point
func ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, `{
		"message": "Pong"
	}
	`)
}

//-------------- API ENDPOINT ------------------//

// AllPosts returns all post data
func AllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	fmt.Println("AllPosts working")

}

// CreatePost create a new post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {

	}
}

// UpdatePost update a  spesific post
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	fmt.Println("AllPosts working")

}

// DeletePost remove a spesific post
func DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	fmt.Println("AllPosts working")
}

func catch(err error) {
	if err != nil {
		panic(err)
	}
}

func response(msg string, code int) {

}

func main() {

	route.HandleFunc("/", ping)
	route.HandleFunc("/posts", AllPosts)
	route.HandleFunc("/post/create", CreatePost)
	route.HandleFunc("/post/update", UpdatePost)
	route.HandleFunc("/post/delete", DeletePost)

	http.ListenAndServe(":8005", Logger())
}
