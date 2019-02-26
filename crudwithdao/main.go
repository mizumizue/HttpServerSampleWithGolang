package main

import (
	"cloud.google.com/go/datastore"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopractices/crudwithdao/model"
	"log"
	"net/http"
	"strconv"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/person", handlerPerson)
	router.HandleFunc("/person/{id}", handlerPerson)
	http.Handle("/", router)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	fmt.Println("My server is running. Port is 8080.")
}

func handlerPerson(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var pd model.PersonDao = model.NewPersonDao(ctx, "Your-ProjectId")

	vars := mux.Vars(r)
	var strId string
	if strId = vars["id"]; strId == "" {
		switch r.Method {
		case "GET":
			persons, err := pd.GetAll()
			if err != nil {
				http.Error(w, "Internal server error. Please try again later.", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(persons)
			return

		case "POST":
			if r.Header.Get("Content-Type") != "application/json; charset=UTF-8" {
				http.Error(w, "Request Content-Type is invalid. Expected application/json.", http.StatusBadRequest)
				return
			}

			var person model.Person
			if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
				http.Error(w, "Request Body is invalid.", http.StatusBadRequest)
				return
			}

			person, err := pd.Create(person)
			if err != nil {
				http.Error(w, "Insert is failed.", http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusCreated)
			_ = json.NewEncoder(w).Encode(person)
			return
		}
	}

	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		http.Error(w, "Request param is invalid. Param id is " + vars["id"], http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "GET":
		person, err := pd.Get(id)
		if err != nil {
			if err == datastore.ErrNoSuchEntity {
				http.Error(w, err.Error() + ". Param id is " + vars["id"], http.StatusNotFound)
				return
			}
			http.Error(w, "Internal server error. Please try again later.", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(person)
		return

	case "PUT":
		if r.Header.Get("Content-Type") != "application/json; charset=UTF-8" {
			http.Error(w, "Request Content-Type is invalid. Expected application/json.", http.StatusBadRequest)
			return
		}

		var person model.Person
		if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
			http.Error(w, "Request Body is invalid.", http.StatusBadRequest)
			return
		}

		person, err := pd.Update(person)
		if err != nil {
			http.Error(w, "Update is failed.", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(person)
		return

	case "DELETE":
		if err := pd.Delete(id); err != nil {
			if err == datastore.ErrNoSuchEntity {
				http.Error(w, "Not found", http.StatusNotFound)
			} else {
				http.Error(w, "Delete is failed.", http.StatusInternalServerError)
			}
			return
		}

		w.WriteHeader(http.StatusOK)
		return
	}
}
