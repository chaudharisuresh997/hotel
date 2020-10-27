package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)



func init() {
	Init()
}
func hotelHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	//fmt.Printf("entered %v", vars)
	json.NewEncoder(w).Encode(GetRestaurant(vars["query"], vars["veg"]))
	//fmt.Fprintf(w, "Category: %v\n", GetRestaurant(vars["query"], vars["veg"]))
	//return )

}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/search", hotelHandler).
		Queries(
			"query", "{query}",
			"veg", "{veg}",
		).
		Methods("GET")
	//	r.HandleFunc("/user", UserByValueHandler).Methods("GET")
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:9000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
