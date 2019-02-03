package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Person struct {
	ID		int			`json:"id"`
	Name 	string		`json:"name"`
	Age 	int			`json:"age"`
	Sex 	string		`json:"sex"`
}

type Persons []Person

var persons = Persons{
	{1, "Tarou", 23, "male"},
	{2, "Hanako", 25, "female"},
	{3, "Kojiro", 29, "male"},
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/person", handler)
	router.HandleFunc("/person/{id}", handler)
	http.Handle("/", router)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	fmt.Println("My server is running. Port is 8080.")
}

func handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var strId string
	if strId = vars["id"]; strId == "" {
		switch r.Method {
			case "GET":
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				w.WriteHeader(http.StatusOK)
				_ = json.NewEncoder(w).Encode(persons)
				return

			case "POST":
				if r.Header.Get("Content-Type") != "application/json; charset=UTF-8" {
					http.Error(w, "Request Content-Type is invalid. Expected application/json.", http.StatusBadRequest)
					return
				}

				var person Person
				if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
					http.Error(w, "Request Body is invalid.", http.StatusBadRequest)
					return
				}

				// TODO Insert Entity
				w.WriteHeader(http.StatusCreated)
				return
		}
	}

	switch r.Method {
		case "GET":
			id, err := strconv.Atoi(strId)
			if err != nil {
				http.Error(w, "Request param is invalid. Param id is " + vars["id"], http.StatusBadRequest)
				return
			}

			var person Person
			if err := getPerson(id, &person); err != nil {
				http.Error(w, "Target person is not found. Param id is " + strId, http.StatusBadRequest)
				return
			}

			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(person)
			return

		case "PUT":
			// TODO Update

			w.WriteHeader(http.StatusOK)
			return

		case "DELETE":
			//TODO Delete

			w.WriteHeader(http.StatusOK)
			return
	}
}

func getPerson(index int, p *Person) error {
	for i := range persons {
		if (persons[i].ID == index) {
			p.ID = persons[i].ID
			p.Name = persons[i].Name
			p.Age = persons[i].Age
			p.Sex = persons[i].Sex
			return nil
		}
	}
	return fmt.Errorf("index " + strconv.Itoa(index) + " is not found.")
}
