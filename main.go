package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sync"
)

type Author struct {
	Id   string
	Name string
}

type Event struct {
	Id     string
	Title  string
	Desc   string
	Author *Author
}

var Authors []Author
var Events []Event

var GlobalMutex sync.Mutex

func homePage(w http.ResponseWriter, r *http.Request) {
	GlobalMutex.Lock()
	defer GlobalMutex.Unlock()

	fmt.Fprintf(w, "helloWorld")
}

func returnAllEvents(w http.ResponseWriter, r *http.Request) {
	GlobalMutex.Lock()
	defer GlobalMutex.Unlock()

	json.NewEncoder(w).Encode(Events)
}

func returnSingleEvent(w http.ResponseWriter, r *http.Request) {
	GlobalMutex.Lock()
	defer GlobalMutex.Unlock()

	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Println("key: ", key)
	for _, event := range Events {
		if event.Id == key {
			json.NewEncoder(w).Encode(event)
		}
	}
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/events", returnAllEvents)
	router.HandleFunc("/event/{id}", returnSingleEvent)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	Authors = []Author{
		{
			Id:   "5",
			Name: "AuthorName",
		},
	}

	Auth := &Authors[0]

	Events = []Event{
		Event{Id: "1", Title: "Test", Desc: "Test Desc", Author: Auth},
		Event{Id: "2", Title: "2Test", Desc: "2Test Desc"},
	}
	handleRequests()
}
