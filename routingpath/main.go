package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hoge", handlerHoge)
	router.HandleFunc("/foo", handlerFoo)
	router.HandleFunc("/bar", handlerBar)
	http.Handle("/", router)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	fmt.Println("My server is running. Port is 8080.")
}

func handlerHoge(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "plain/text; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Routing Hoge")
}

func handlerFoo(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "plain/text; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Routing Foo")
}

func handlerBar(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "plain/text; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Routing Bar")
}
