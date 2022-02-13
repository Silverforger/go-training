package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Contact struct {
	ID       int
	Last     string
	First    string
	Company  string
	Address  string
	Country  string
	Position string
}

type DataStruc struct {
	nextID   int
	mu       sync.Mutex
	contacts []Contact
}

// handler function checks if URL is valid and runs process and processID functions based on input url
func (db *DataStruc) handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ID int
		if r.URL.Path == "/contacts" {
			db.process(w, r)
		} else if n, _ := fmt.Sscanf(r.URL.Path, "/contacts/%d", &ID); n == 1 {
			db.processID(ID, w, r)
		} else {
			http.Error(w, "Invalid URL.", http.StatusBadRequest)
		}
	}
}

// process function handles all methods associated to /contacts, such as POST and GET for db.contacts
// PUT and DELETE is not allowed in this function.
func (db *DataStruc) process(w http.ResponseWriter, r *http.Request) {
	var contact Contact
	switch r.Method {
	case "POST":
		if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		for _, existing := range db.contacts {
			if contact.First == existing.First && contact.Last == existing.Last {
				http.Error(w, "That contact already exists on the database.", http.StatusConflict)
				return
			}
		}

		db.mu.Lock()
		contact.ID = db.nextID
		db.nextID++
		db.contacts = append(db.contacts, contact)
		db.mu.Unlock()
		fmt.Println(w, "Contact added to database.", http.StatusCreated)

		return

	case "GET":
		w.Header().Set("Content-Type", "application/json")
		if len(db.contacts) == 0 {
			http.Error(w, "Retrieved nothing. Database is empty.", http.StatusConflict)
			return
		}
		if err := json.NewEncoder(w).Encode(db.contacts); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	default:
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "Error 405: HTTP method not allowed.", http.StatusMethodNotAllowed)
	}
}

// processID function takes in an ID input and either deletes that ID or updates it based on the method applied (PUT/DELETE).
// GET method is also allowed here, but POST is not.
func (db *DataStruc) processID(id int, w http.ResponseWriter, r *http.Request) {
	var contact Contact
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")

		if len(db.contacts) == 0 {
			http.Error(w, "Retrieved nothing. Database is empty.", http.StatusConflict)
			return
		}

		for index, existing := range db.contacts {
			if id == existing.ID {
				if err := json.NewEncoder(w).Encode(db.contacts[index]); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				return
			}
		}

		http.Error(w, "Contact not found in database.", http.StatusNotFound)

	case "DELETE":
		db.mu.Lock()
		for index, existing := range db.contacts {
			if id == existing.ID {
				db.contacts = append(db.contacts[:index], db.contacts[index+1:]...)
				fmt.Fprintf(w, "Contact with ID %v successfully deleted.\n", id)
				db.mu.Unlock()
				return
			}
		}
		db.mu.Unlock()

		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "Contact not found in database.", http.StatusNotFound)

	case "PUT":
		if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		db.mu.Lock()

		for index, existing := range db.contacts {
			if id == existing.ID {
				db.contacts[index] = contact
				db.contacts[index].ID = index
				db.mu.Unlock()
				fmt.Fprintf(w, "Contact with ID %v successfully updated.\n", id)
				return
			}
		}
		db.mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "Contact not found in database.", http.StatusNotFound)

	default:
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "Error 405: HTTP method not allowed.", http.StatusMethodNotAllowed)
	}
}

// main function for running the server.
func main() {
	db := DataStruc{contacts: []Contact{}}
	http.ListenAndServe(":8080", db.handler())
}
