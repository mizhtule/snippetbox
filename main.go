package main

import (
	"log"
	"net/http"
)

func main() {
	// Use the http.NewServeMux() function to initialize a new serve mux, then
	// register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the serve mux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit.
	// Note that any error returned by http.ListenAndServe() is always non-nil.
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Create a new snippet . . . "))
}

// snippetView add a snippetView handler function.
func snippetView(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Display a specific snippet . . . "))
}

// snippetCreate add snippetCreate handler function.
func home(w http.ResponseWriter, r *http.Request) {
	// Check if the current request URL path exactly maches "/". If it doesn't, use
	// the http.NotFound() function to send a 404 response to the client.
	// Importantly, we then return from the handler. If we don't return the handler
	// would keep executing and also write the "Hello from Snippetbox!" message.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	_, _ = w.Write([]byte("Hello from Snippetbox!"))
}
