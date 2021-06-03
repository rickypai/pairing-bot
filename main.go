package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
)

// It's alive! The application starts here.
func main() {

	// setting up database connection
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "pairing-bot-284823")
	if err != nil {
		log.Panic(err)
	}
	defer client.Close()

	fc := &firestoreClient{client: client, ctx: ctx}

	// Use a method on a struct for the handler, and set global state there?

	http.HandleFunc("/", nope)
	http.HandleFunc("/webhooks", fc.handle)
	http.HandleFunc("/match", fc.match)
	http.HandleFunc("/endofbatch", fc.endofbatch)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
