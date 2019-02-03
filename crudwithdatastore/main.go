package main

import (
	"cloud.google.com/go/datastore"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Person struct {
	Key 	*datastore.Key 	`datastore:"__key__"json:"-"`	// JSONには出力しない
	ID		int64		`datastore:"-"json:"id"`			// JSONでは返すが、Save時には無視する
	Name	string		`json:"name"`
	Age		int			`json:"age"`
	Sex		string		`json:"sex"`
	Address	string		`json:"address"`
}

type Persons []Person

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
	dsClient, err := newClient(ctx)
	defer dsClient.Close()
	if err != nil {
		http.Error(w, "Internal server error. Please try again later.", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	var strId string
	if strId = vars["id"]; strId == "" {
		switch r.Method {
		case "GET":
			persons, err := getPersons(ctx, dsClient)
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

			var person Person
			if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
				http.Error(w, "Request Body is invalid.", http.StatusBadRequest)
				return
			}

			if err := insertPerson(ctx, dsClient, person); err != nil {
				http.Error(w, "Insert is failed.", http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusCreated)
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
			person, err := getPerson(ctx, dsClient, id);
			if err != nil {
				if err == datastore.ErrNoSuchEntity {
					http.Error(w, err.Error() + ". Param id is " + vars["id"], http.StatusNotFound)
					return
				}
				http.Error(w, "Internal server error. Please try again later.", http.StatusInternalServerError)
				return
			}
			person.ID = person.Key.ID
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(person)
			return

		case "PUT":
			if r.Header.Get("Content-Type") != "application/json; charset=UTF-8" {
				http.Error(w, "Request Content-Type is invalid. Expected application/json.", http.StatusBadRequest)
				return
			}

			var person Person
			if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
				http.Error(w, "Request Body is invalid.", http.StatusBadRequest)
				return
			}

			if err := updatePerson(ctx, dsClient, id, person); err != nil {
				http.Error(w, "Update is failed.", http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			return

		case "DELETE":
			_, err := getPerson(ctx, dsClient, id);
			if err != nil {
				if err == datastore.ErrNoSuchEntity {
					http.Error(w, err.Error() + ". Param id is " + vars["id"], http.StatusNotFound)
					return
				}
				http.Error(w, "Internal server error. Please try again later.", http.StatusInternalServerError)
				return
			}

			if err := deletePerson(ctx, dsClient, id); err != nil {
				http.Error(w, "Delete is failed.", http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			return
	}
}

func getPersons(ctx context.Context, dsClient *datastore.Client) ([]Person, error) {
	q := datastore.NewQuery("Person")
	var ps []Person
	if _, err := dsClient.GetAll(ctx, q, &ps); err != nil {
		return ps, err
	}
	for i := 0; i < len(ps); i++ {
		ps[i].ID = ps[i].Key.ID
	}
	return ps, nil
}

func getPerson(ctx context.Context, dsClient *datastore.Client, id int64) (Person, error) {
	var p Person
	key := datastore.IDKey("Person", id, nil)
	if err := dsClient.Get(ctx, key, &p); err != nil {
		return p, err
	}
	return p, nil
}

func insertPerson(ctx context.Context, dsClient *datastore.Client, p Person) error {
	key := datastore.IncompleteKey("Person", nil)
	_, err := dsClient.Put(ctx, key, &p)
	if err != nil {
		return err
	}
	return nil
}

func updatePerson(ctx context.Context, dsClient *datastore.Client, id int64, p Person) error {
	key := datastore.IDKey("Person", id, nil)
	_, err := dsClient.Put(ctx, key, &p)
	if err != nil {
		return err
	}
	return nil
}

func deletePerson(ctx context.Context, dsClient *datastore.Client, id int64) error {
	key := datastore.IDKey("Person", id, nil)
	if err := dsClient.Delete(ctx, key); err != nil {
		return err
	}
	return nil
}

func newClient(ctx context.Context) (*datastore.Client, error) {
	projectId := os.Getenv("PROJECT_ID")
	dsClient, err := datastore.NewClient(ctx, projectId)
	if err != nil {
		return nil, err
	}
	return dsClient, nil
}
