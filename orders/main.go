package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type song struct {
	title, artist, genre string
}

func songEntries(data [][]string) []song {
	var songs []song
	for _, row := range data {
		title := row[0]
		artist := row[1]
		genre := row[2]
		songs = append(songs, song{title, artist, genre})
	}
	return songs
}

func main() {
	// The top 50 most listened songs in 2019 in the world by Spotify.
	f, err := os.Open("orders/songs.csv")
	if err != nil {
		log.Fatalf("unable to open a file: %v", err)
	}
	defer f.Close()
	orders, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatalf("failed to parse a CSV file: %v", err)
	}
	songs := songEntries(orders)
	for _, song := range songs {
		fmt.Printf("%+v\n", song)
	}
}
