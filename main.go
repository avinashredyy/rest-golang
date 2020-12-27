package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

//1. We have seperated each kind of request into its own individual function.
//2. One specific function will run depending on the tpe of request we get on our route
func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))
}

func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "put called"}`))
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "delete called"}`))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}

func main() {
	//1. We are using Gorilla mux here which is a third party http library which helps us with certain additional helper
	//functions that would have otherwise have to be coded explicitly
	//2. We initalize the gorilla mux like below
	r := mux.NewRouter()
	//1. Gorilla mux here provides us this helper function like Methods() in the example below which take in the type of request
	//as an argument and essentially says to http.HandleFunc() that, if the request is of type GET, execute line 52 or if the
	//request is of type POST, execute line 53, so forth and so on.
	//2. If we had not used gorilla mux we would have had to identify the type of request that we recieve by deconstructing
	//*http.Request and then exectuing lines of code or another function based on the type of request.
	//3. So basically now we can execute a different method based on the type of request for the same path("/" in this example)
	r.HandleFunc("/", get).Methods(http.MethodGet)
	r.HandleFunc("/", post).Methods(http.MethodPost)
	r.HandleFunc("/", put).Methods(http.MethodPut)
	r.HandleFunc("/", delete).Methods(http.MethodDelete)
	r.HandleFunc("/", notFound)
	http.ListenAndServe(":8000", r)
}
