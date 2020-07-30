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

// Benchmarking of the request time
var now = time.Now()

func main() {
	http.HandleFunc("/", CharactersListHandler)
	http.ListenAndServe(":8000", nil)
}

func CharactersListHandler(w http.ResponseWriter, r *http.Request) {
	results, err := rickAndMortyApi.GetAllCharacters()
	LogCharacters(results)
	tmpl, err := template.New("").ParseFiles("rickAndMortyApi/templates/characters.html", "rickAndMortyApi/templates/base.html")

	if err != nil {
		panic(err)
	}
	err = tmpl.ExecuteTemplate(w, "base", results)

	if err != nil {
		log.Fatalf("error while getting all characters: %v", err)
	}

}

func LogCharacters(results []rickAndMortyApi.Character) () {
	rosterFile, err := os.OpenFile("roasters.txt", os.O_RDWR|os.O_CREATE|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("error opening the file roasters.txt: %v", err)
	}

	defer rosterFile.Close()

	// log at the same time in terminal and in the roasterFile
	wrt := io.MultiWriter(os.Stdout, rosterFile)
	log.SetOutput(wrt)
	for _, character := range results {
		log.Println("------------------------------")
		log.Printf("Name %s", character.Name)
		log.Println("------------------------------")
	}
	log.Printf("took %v", time.Now().Sub(now).String())
}
