package main

import (
	"database/sql"
	"encoding/json"
	"os"

	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB

type Todo struct {
	Id          int    `json: "id", db: "id"`
	Name        string `json:"name", db:"name"`
	Description string `json:"description", db:"description"`
	DueDate     string `json:"duedate", db:"duedate"`
	Status      string `json:"status", db:"status"`
}

func main() {
	initDB()
	initRouter()
}

/********************************/
/************ routes ************/
/********************************/

func initRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/", Home).Methods("GET")

	router.HandleFunc("/todos/", GetTodos).Methods("GET")

	router.HandleFunc("/todo/{id}/", GetTodo).Methods("GET")

	router.HandleFunc("/todo/", CreateTodo).Methods("POST")

	router.HandleFunc("/todo/{id}/", UpdateTodo).Methods("PUT")

	router.HandleFunc("/todo/{id}/", DeleteTodo).Methods("DELETE")

	// start server
	log.Fatal(http.ListenAndServe(":8000", router))
}

/********************************/
/*********** handlers ***********/
/********************************/

/*** home ***/
var Home = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Golang ToDo REST API")
}

/*** index ***/
var GetTodos = func(w http.ResponseWriter, r *http.Request) {
	var todos []Todo

	SqlStatement := `
        SELECT * FROM todos
    `

	rows, err := db.Query(SqlStatement)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var Id int
		var Name string
		var Description string
		var DueDate string
		var Status string

		rows.Scan(&Id, &Name, &Description, &DueDate, &Status)

		todos = append(todos, Todo{
			Id:          Id,
			Name:        Name,
			Description: Description,
			DueDate:     DueDate,
			Status:      Status,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

/*** show ***/
var GetTodo = func(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	params := mux.Vars(r)

	SqlStatement := `
        SELECT * FROM todos
        WHERE id = $1
    `

	err := db.QueryRow(
		SqlStatement,
		params["id"],
	).Scan(&todo.Id, &todo.Name, &todo.Description, &todo.DueDate, &todo.Status)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&todo)
}

/*** create ***/
var CreateTodo = func(w http.ResponseWriter, r *http.Request) {
	todo := &Todo{}

	err := json.NewDecoder(r.Body).Decode(todo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	SqlStatement := `
        INSERT INTO todos (name, description, duedate, status)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `

	id := 0
	err = db.QueryRow(SqlStatement, todo.Name, todo.Description, &todo.DueDate, "new").Scan(&id)
	if err != nil {
		panic(err)
	}

	todo.Id = id

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

/*** update ***/
var UpdateTodo = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var todo Todo

	r.ParseForm()

	SqlStatement := `
	UPDATE todos
	SET name = $2, description = $3, duedate = $4, status = $5
	WHERE id = $1
	RETURNING *
`

	err := db.QueryRow(
		SqlStatement,
		params["id"],
		r.FormValue("name"),
		r.FormValue("description"),
		r.FormValue("duedate"),
		r.FormValue("status"),
	).Scan(&todo.Id, &todo.Name, &todo.Description, &todo.DueDate, &todo.Status)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&todo)
}

/*** delete ***/
var DeleteTodo = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var todo Todo

	SqlStatement := `
        DELETE FROM todos
        WHERE id = $1
        RETURNING *
    `

	err := db.QueryRow(
		SqlStatement,
		params["id"],
	).Scan(&todo.Id, &todo.Name, &todo.Description, &todo.DueDate, &todo.Status)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&todo)
}

/********************************/
/************** db **************/
/********************************/

// hook up to postgres db
func initDB() {
	var err error
	db_pass := os.Getenv("DATABASE_PASSWORD")
	db, err = sql.Open("postgres", "user=postgres password="+db_pass+" dbname=gotodo sslmode=disable")

	if err != nil {
		panic(err)
	}
}
