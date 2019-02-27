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
	Key	*datastore.Key 	`datastore:"__key__"json:"-"`	// JSONには出力しない
	ID	int64	`datastore:"-"json:"id"`	// JSONでは返すが、Save時には無視する
	Name	string	`json:"name"`
	Age	int	`json:"age"`
	Sex	string	`json:"sex"`
	Address	string	`json:"address"`
	OrganizationKey	*datastore.Key `json:"-"`
	Organization Organization	`json:"organization,omitempty"`
}

type Persons []Person

type Organization struct {
	Key	*datastore.Key 	`datastore:"__key__"json:"-"`	// JSONには出力しない
	ID	int64	`datastore:"-"json:"id"`	// JSONでは返すが、Save時には無視する
	Name	string	`json:"name"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/organization/{id}/person", handlerOrganization)
	http.Handle("/", router)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	fmt.Println("My server is running. Port is 8080.")
}

func handlerOrganization(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "Request param is empty. Param id is " + vars["id"], http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		http.Error(w, "Request param is invalid. Param id is " + vars["id"], http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "GET":
		orgKey := datastore.IDKey("Organization", id, nil)
		persons, err := getPersonByOrganization(ctx, dsClient, orgKey)
		if err != nil {
			http.Error(w, "Internal server error. Please try again later.", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(persons)
		return
	}
}

func getPerson(ctx context.Context, dsClient *datastore.Client, id int64) (Person, error) {
	var p Person
	key := datastore.IDKey("Person", id, nil)
	if err := dsClient.Get(ctx, key, &p); err != nil {
		return p, err
	}
	p.ID = p.Key.ID
	return p, nil
}

func getOrganization(ctx context.Context, dsClient *datastore.Client, key *datastore.Key) (Organization, error) {
	var o Organization
	if err := dsClient.Get(ctx, key, &o); err != nil {
		return o, err
	}
	o.ID = o.Key.ID
	return o, nil
}

func getPersonByOrganization(ctx context.Context, dsClient *datastore.Client, key *datastore.Key) ([]Person, error){
	var ps []Person
	q := datastore.NewQuery("Person").Filter("OrganizationKey=", key)
	_, err := dsClient.GetAll(ctx, q, &ps)
	if err != nil {
		return ps, err
	}
	return ps, nil
}

func newClient(ctx context.Context) (*datastore.Client, error) {
	projectId := os.Getenv("PROJECT_ID")
	dsClient, err := datastore.NewClient(ctx, projectId)
	if err != nil {
		return nil, err
	}
	return dsClient, nil
}
