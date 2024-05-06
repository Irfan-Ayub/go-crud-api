package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Todo struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var todos []Todo

func main() {
	router := mux.NewRouter()

	// Initialize some sample todos
	todos = append(todos, Todo{ID: 1, Title: "Todo 1", Content: "Todo 1 content"})
	todos = append(todos, Todo{ID: 2, Title: "Todo 2", Content: "Todo 2 content"})

	// Define API endpoints
	router.HandleFunc("/todos", getTodos).Methods("GET")
	router.HandleFunc("/todos/{id}", getTodo).Methods("GET")
	router.HandleFunc("/todos", createTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", updateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", deleteTodo).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	log.Println("Get all todos")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("Get todo by id")

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range todos {
		id, _ := strconv.Atoi(params["id"])
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Todo{})
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("Create todo")

	w.Header().Set("Content-Type", "application/json")
	var todo Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todo.ID = len(todos) + 1
	todos = append(todos, todo)
	json.NewEncoder(w).Encode(todo)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("Update todo")

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range todos {
		if strconv.Itoa(item.ID) == params["id"] {
			var todo Todo
			_ = json.NewDecoder(r.Body).Decode(&todo)
			todo.ID = item.ID
			todos[index] = todo
			json.NewEncoder(w).Encode(todo)
			return
		}
	}
	log.Println("Todo not found")
	json.NewEncoder(w).Encode(&Todo{})
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete todo")

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range todos {
		if strconv.Itoa(item.ID) == params["id"] {
			todos = append(todos[:index], todos[index+1:]...)
			json.NewEncoder(w).Encode(todos)
			return
		}
	}
	log.Println("Todo not found")
	json.NewEncoder(w).Encode(&Todo{})
}
