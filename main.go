package main
import (
	"fmt"
	"net/http"
)
func main() {
	handler := http.NewServeMux()
	///we create a new router to expose our api
	//to our users
	handler.HandleFunc("/api/hello", SayHello)
	//Every time a  request is sent to the endpoint ("/api/hello") the function SayHello will be invoked
	http.ListenAndServe("0.0.0.0:8080", handler)
	//we tell our api to listen to all request to port 8080.
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	//notice how this function takes two parameters
	//the first parameter is a ResponseWriter writer and this is where we wite the response we want to send back to the user
	//the second parameter is a pointer of type  http.Request this holds all inforamtion of the request sent by the used
	//this may include query parameters,path parameters and many more
	fmt.Fprintf(w, `Hello world`)

}
