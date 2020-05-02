package main

import (
	"fmt"
	"net/http"

	"./handler"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	setupStaticFiles(router)
	setupAPIEndpoints(router)
	setupTemplate(router)

	fmt.Println("The server is running on 8080.")
	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}

func setupAPIEndpoints(router *mux.Router) {
	rh := handler.NewHandler()

	router.HandleFunc("/", rh.RegisterReport).Methods("POST")
	router.HandleFunc("/", rh.Index).Methods("GET")
	router.HandleFunc("/t", rh.IndexT).Methods("GET")

}

func setupTemplate(router *mux.Router) {

}

func setupStaticFiles(router *mux.Router) {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	router.PathPrefix("/static/").
		Handler(http.StripPrefix("/static/", fs))

}
