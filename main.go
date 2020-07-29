package main

import (
	"github.com/macaguegi/rickAndGorty/rickAndMortyApi"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// Benchmarking of the request time
	now := time.Now()

	rosterFile, err := os.OpenFile("roasters.txt", os.O_RDWR|os.O_CREATE|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("error opening the file roasters.txt: %v", err)
	}

	defer rosterFile.Close()

	// log at the same time in terminal and in the roasterFile
	wrt := io.MultiWriter(os.Stdout, rosterFile)
	log.SetOutput(wrt)

	http.HandleFunc("/", CharactersListHandler)
	http.ListenAndServe(":8000", nil)

	log.Printf("took %v", time.Now().Sub(now).String())
}

func CharactersListHandler(w http.ResponseWriter, r *http.Request) {
	results, err := rickAndMortyApi.GetAllCharacters()

	tmpl, err := template.New("").ParseFiles("rickAndMortyApi/templates/characters.html", "rickAndMortyApi/templates/base.html")
	// check your err
	if err != nil {
		panic(err)
	}
	err = tmpl.ExecuteTemplate(w, "base", results)

	if err != nil {
		log.Fatalf("error while getting all characters: %v", err)
	}
}
