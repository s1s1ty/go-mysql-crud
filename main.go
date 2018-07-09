package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
)

var router *chi.Mux
var db *sql.DB

const (
	dbName = "go-mysql-crud"
	dbPass = "12345"
	dbHost = "localhost"
	dbPort = "33066"
)

// Post type details
type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	// created_at time.Time `json:"created_at"`
}

func init() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)

	var err error

	dbSource := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8", dbPass, dbHost, dbPort, dbName)
	db, err = sql.Open("mysql", dbSource)
	catch(err)
}

func routers() *chi.Mux {
	router.Get("/", ping)

	router.Get("/posts", AllPosts)
	router.Get("/posts/{id}", DetailPost)
	router.Post("/posts/create", CreatePost)
	router.Put("/posts/update/{id}", UpdatePost)
	router.Delete("/posts/{id}", DeletePost)

	return router
}

// server starting point
func ping(w http.ResponseWriter, r *http.Request) {
	respondwithJSON(w, http.StatusOK, map[string]string{"message": "Pong"})
}

//-------------- API ENDPOINT ------------------//

// AllPosts returns all post data
func AllPosts(w http.ResponseWriter, r *http.Request) {
	errors := []error{}
	payload := []Post{}

	rows, err := db.Query("Select id, title, content From posts")
	catch(err)

	defer rows.Close()

	for rows.Next() {
		data := Post{}

		er := rows.Scan(&data.ID, &data.Title, &data.Content)

		if er != nil {
			errors = append(errors, er)
		}
		payload = append(payload, data)
	}

	respondwithJSON(w, http.StatusOK, payload)
}

// CreatePost create a new post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post Post
	json.NewDecoder(r.Body).Decode(&post)

	query, err := db.Prepare("Insert posts SET title=?, content=?")
	catch(err)

	_, er := query.Exec(post.Title, post.Content)
	catch(er)
	defer query.Close()

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
}

// DetailPost return a post by ID
func DetailPost(w http.ResponseWriter, r *http.Request) {
	payload := Post{}
	id := chi.URLParam(r, "id")

	row := db.QueryRow("Select id, title, content From posts where id=?", id)

	err := row.Scan(
		&payload.ID,
		&payload.Title,
		&payload.Content,
		// &payload.created_at,
	)

	if err != nil {
		respondWithError(w, http.StatusNotFound, "no rows in result set")
		return
	}

	respondwithJSON(w, http.StatusOK, payload)
}

// UpdatePost update a  spesific post
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	var post Post
	id := chi.URLParam(r, "id")
	json.NewDecoder(r.Body).Decode(&post)

	query, err := db.Prepare("Update posts set title=?, content=? where id=?")
	catch(err)
	_, er := query.Exec(post.Title, post.Content, id)
	catch(er)

	defer query.Close()

	respondwithJSON(w, http.StatusOK, map[string]string{"message": "update successfully"})

}

// DeletePost remove a spesific post
func DeletePost(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	query, err := db.Prepare("delete from posts where id=?")
	catch(err)
	_, er := query.Exec(id)
	catch(er)
	query.Close()

	respondwithJSON(w, http.StatusOK, map[string]string{"message": "successfully deleted"})
}

func main() {
	routers()
	http.ListenAndServe(":8005", Logger())
}
