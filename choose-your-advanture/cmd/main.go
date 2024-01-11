package main

import (
	"cyoa"
	"flag"
	"fmt"
	"os"
)

func main() {
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

	fmt.Printf("%+v\n", story)
}
