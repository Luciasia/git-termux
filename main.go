package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Define a simple data structure
type Person struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Slice to hold data
var people []Person

// Get all people
func getPeople(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

// Get a person by ID
func getPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get URL parameters
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

// Create a new person
func createPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = "123" // In a real application, you would generate a unique ID
	people = append(people, person)
	json.NewEncoder(w).Encode(person)
}

// Delete a person by ID
func deletePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main() {
	// Initialize router
	r := mux.NewRouter()

	// Mock data
	people = append(people, Person{ID: "1", Name: "John Doe", Age: 30})
	people = append(people, Person{ID: "2", Name: "Jane Doe", Age: 25})

	// Route handlers / endpoints
	r.HandleFunc("/people", getPeople).Methods("GET")
	r.HandleFunc("/people/{id}", getPerson).Methods("GET")
	r.HandleFunc("/people", createPerson).Methods("POST")
	r.HandleFunc("/people/{id}", deletePerson).Methods("DELETE")

	// Start server
	http.ListenAndServe(":8000", r)
}
