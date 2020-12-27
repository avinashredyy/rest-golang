package main

import (
	"net/http"
)

// Earlier we used this struct to attach it to our function so that we can access it in main(). But http.HandleFunc() takes care
//of it
//type server struct{}

//1. We implement the home method which has a switch case implemented depending on the type of request we get.
//2. The function signature of home is simalr to ServeHTTP. That is, we are using http.ResponseWriter, *http.Request just like
//ServeHTTP.
//3. So basically if you want to use your function with http.HandleFunc(), you need to pass http.ResponseWriter, *http.Request
//as parameters to your function.
//4. The http.HandleFunc() allows us to pass our function home that has the same signature as the ServeHTTP
//5. Since we will be passing this function to the http.HandleFunc() it automatically detects that our function satisfies
//the function signature condition of ServeHTTP
func home(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(`{"message": "get called"}`))
	case "POST":
		rw.WriteHeader(http.StatusCreated)
		rw.Write([]byte(`{"message": "post called"}`))
	case "PUT":
		rw.WriteHeader(http.StatusAccepted)
		rw.Write([]byte(`{"message": "put called"}`))
	case "DELETE":
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(`{"message": "delete called"}`))
	default:
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte(`{"message": "not found"}`))
	}
}

func main() {
	//Earlier we initalize the server struct so that we can use the functions associated with it in main(). We dont need it anymore.
	//s := &server{}

	//1. Handle registers the handler for the given pattern and the second argument is the function that we want to execute when
	//the mentioned path in the first parameter is hit("/" in this case).
	//2. Only http.HandleFunc() allows us to pass any function(home() in this case) that satisfes the function signature condition
	//of ServeHTTP
	//3. http.Handle() does not allow this. For http.Handle() we need to create a struct and then pass it as a reciever to our
	//function. Then initialize that struct in main so that we get access to the function.
	http.HandleFunc("/", home)
	//This starts the server and listens for requests
	http.ListenAndServe(":8000", nil)
}
