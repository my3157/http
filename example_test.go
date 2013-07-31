package http

import (
	"log"
	"os"
)

func ExampleGet() {
	// curl in 3 lines of code.
	if _, err := Get(os.Stdout, "http://www.gorillatoolkit.org/"); err != nil {
		log.Fatalf("could not fetch: %v", err)
	}
}

func ExampleClient_Post() {
	// send the contents of os.Stdin to a remote webserver.
	status, _, r, err := DefaultClient.Post("http://example.com", nil, os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	if r != nil {
		defer r.Close()
	}
	log.Printf("Post result: %v", status)
}
