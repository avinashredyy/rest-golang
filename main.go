package main

import (
	"fmt"
	"net/http"
)

// We create a server struct here to pass it as a reciver so that we could use it in main()
type server struct{}

//1. We implement the ServeHTTP method provided by net/http package which is essentially a Handler interface.
//2. ServeHTTP method satisfies the Handler interface's signature. Hence, the compiler automaticall detects that and lets
//us work with ServeHTTP directly.
//3. ServeHTTP recieves server struct so that we can access this method in main(). ServeHTTP also takes 2 parameters which are
//essentially request and response objects.
func (s *server) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(`{"message": "Hello World!"}`))
}

func main() {
	//We initalize the server struct so that we can use the functions associated with it in main()
	s := &server{}
	fmt.Println(s)
	//Handle registers the handler for the given pattern and the second argument is the function that we want to execute when
	//the mentioned path in the first parameter is hit("/" in this case).
	http.Handle("/", s)
	//This starts the server and listens for requests
	http.ListenAndServe(":8000", nil)
}
