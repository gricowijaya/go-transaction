package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"          // for asynchronous functions
)

// for global counter
var counter int
// the mutex getting the synchronization 
var mutex = &sync.Mutex{} 

// homepage controller
func home(w http.ResponseWriter, r *http.Request) { 
    // return word "Homepage" to the web page
    fmt.Fprintf(w, "Homepage");
    // print to trace
    _, err := fmt.Println("GET / 200 OK")

    // error handling
    if(err != nil) { 
        fmt.Println("GET /increment 500 BAD")
    }
}

// asynchronous function with mutex
func increment(w http.ResponseWriter, r *http.Request) {
    // set before using the shard resource 
    mutex.Lock()
    counter++
    fmt.Fprintf(w, strconv.Itoa(counter))
    // print to trace
    _, err := fmt.Println("GET /increment 200 OK")
    // error handling
    if(err != nil) { 
        fmt.Println("GET /increment 500 BAD")
    }
    // unlock when the scheduled processes such as unlocking 
    mutex.Unlock()
}

// function to handle the ./index.html file
func about(w http.ResponseWriter, r *http.Request) {
    // getting the file from the ./index.html
    http.ServeFile(w, r, r.URL.Path[1:])
    fmt.Println("GET /static 200 OK")
}

// create the request handlers or routes 
func handleRequests() {
    // create the route of home
    http.HandleFunc("/", home)
    // create the route of increment function it is an asynchronous function.
    http.HandleFunc("/increment", increment)
    // router for serving the file in ./index.html
    http.Handle("/static", http.FileServer(http.Dir("./static")))
    // router for serving the static/html file
    http.HandleFunc("/about", about)
}

// serve function
func serve() { 
    // Trace the endpoint 
    fmt.Println("Listen on :3000");
    // listen and serve the endpoint
    log.Fatal(http.ListenAndServe(":3000", nil))
}

func main() {
    // call the handleRequest
    handleRequests()
    serve()
}


// ref https://tutorialedge.net/golang/creating-simple-web-server-with-golang/
