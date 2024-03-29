package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/promethean-tech/go-vs-node-api/tree/main/go/interal/handlers"
	log "github.com/sirupsen/logrus"
)

// the main go function that deploys the routes
func main() {

	// create an instance of chi
	r := chi.NewRouter()

	// create the handlers, this seems to be the standard found here
	// https://drstearns.github.io/tutorials/gomiddleware/
	handlers.Handler(r)

	// Define a handler function for the "/" route
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// Send the response "HELLO FROM GO"
		fmt.Fprint(w, "HELLO WORLD")
	})

	// Start the HTTP server on localhost:8000
	port := 8000
	fmt.Printf("Server is running on http://localhost:%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), r)

	if err != nil {
		log.Error((err))
	}
}
