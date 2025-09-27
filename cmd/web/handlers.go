package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// snippetCreate add snippetCreate handler function.
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// Use r.Method to check whether the request is using POST or not.
	if r.Method != http.MethodPost {
		// If it's not, use the w.WriteHeader() method to send a 405 status
		// code and the w.Write() method to write a 405 status text
		// response body. We then return from the function so that the
		// subsequent code is not executed.
		w.Header().Set("Allow", http.MethodPost)
		// Use the http.Error() function to send a 405 status code and
		// "Method Not Allowed" string as the response body.
		statusCode := http.StatusMethodNotAllowed
		http.Error(w, http.StatusText(statusCode), statusCode)
		return
	}

	_, _ = w.Write([]byte("Create a new snippet . . . "))
}

// snippetView add a snippetView handler function.
func snippetView(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id parameter from the query string and try to
	// convert it to an integer using the strconv.Atoi() function. If it can't
	// be converted to an integer, or the value is less than 1, we return a 404 page
	// not found response.
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// Use the fmt.Fprintf() function to interpolate the id value with our response
	// and write it to the http.ResponseWriter.
	_, _ = fmt.Fprintf(w, "Display a specific snippet with ID %d . . . ", id)
}

func home(w http.ResponseWriter, r *http.Request) {
	// Check if the current request URL path exactly matches "/". If it doesn't, use
	// the http.NotFound() function to send a 404 response to the client.
	// Importantly, we then return from the handler. If we don't return the handler
	// would keep executing and also write the "Hello from Snippetbox!" message.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	_, _ = w.Write([]byte("Hello from Snippetbox!"))
}
