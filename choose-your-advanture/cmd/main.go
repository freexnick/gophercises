package main

import (
	"cyoa"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the CYOA web application on")
	filename := flag.String("file", "gopher.json", "the JSON file with CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "problem opening then file %s, reason: %s.\n", *filename, err)
		os.Exit(0)
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't decode json, %s \n", err)
	}

	h := cyoa.NewHandler(story)

	fmt.Printf("Starting the server at: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
