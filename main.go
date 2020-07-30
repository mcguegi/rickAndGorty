package main

import (
	"github.com/macaguegi/rickAndGorty/rickAndMortyApi"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
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
	CollectLocations(results)
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
		log.Println("-------CHARACTER---------")
		log.Printf("Name %s", character.Name)
		log.Println("-------CHARACTER---------")
	}
	log.Printf("took %v", time.Now().Sub(now).String())
}

func CollectLocations(results []rickAndMortyApi.Character) () {
	var wg sync.WaitGroup

	wg.Add(len(results))

	resu := make(chan []rickAndMortyApi.LocationData)
	for _, character := range results {
		go func(character rickAndMortyApi.Character) {
			location, err := rickAndMortyApi.GetLocationData(character)
			if err != nil {
				log.Fatalf("error getting location data: %v", err)
			}

			resu <- location

			wg.Done()
		}(character)
	}

	go func() {
		wg.Wait()
		close(resu)
	}()

	display(resu)
	log.Printf("took %v", time.Now().Sub(now).String())

}

func display(resu chan []rickAndMortyApi.LocationData) {
	for r := range resu {
		log.Println("----------------LOCATIONS---------------")
		log.Printf("ID: %v", r[0].ID)
		log.Printf("Name: %s", r[0].Name)
		log.Printf("Dimension: %s", r[0].Dimension)
		log.Println("----------------LOCATIONS---------------")
	}
}