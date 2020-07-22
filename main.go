package main

import (
	"github.com/macaguegi/rickAndGorty/rickAndMortyApi"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	// Benchmarking of the request time
	now := time.Now()

	rosterFile, err := os.OpenFile("roasters.txt", os.O_RDWR|os.O_CREATE|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("error opening the file roasters.txt: %v", err)
	}

	defer rosterFile.Close()

	// log at the same time in terminal and in the roasterFile
	wrt := io.MultiWriter(os.Stdout, rosterFile)
	log.SetOutput(wrt)

	results, err := rickAndMortyApi.GetAllCharacters()

	if err != nil {
		log.Fatalf("error while getting all characters: %v", err)
	}

	for _,character := range results {
		log.Println("------------------------------")
		log.Printf("Name %s", character.Name)
		log.Println("------------------------------")
	}
	log.Printf("took %v", time.Now().Sub(now).String())
}
