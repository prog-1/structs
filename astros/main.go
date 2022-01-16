package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type (
	astro struct {
		Craft, Name string
	}
	astros struct {
		People []astro
	}
)

func main() {
	// A list of astronauts who are currently on the ISS.
	// Retrieved from http://api.open-notify.org/astros.json.
	f, err := os.Open("astros/astros.json")
	if err != nil {
		log.Fatalf("unable to open a file: %v", err)
	}
	defer f.Close()

	astronauts := astros{}
	if err := json.NewDecoder(bufio.NewReader(f)).Decode(&astronauts); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", astronauts)
}
