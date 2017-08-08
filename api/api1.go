package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Movie Struct
type Movie struct {
	Title  string `json:"title"`
	Rating string `json:"rating"`
	Year   string `json:"year"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/movies", handleMovies).Methods("GET")
	router.HandleFunc("/films", handleFilms).Methods("GET")
	fmt.Println("Server started")
	http.ListenAndServe(":8080", router)

}

func handleMovies(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var movies = map[string]*Movie{
		"tt0076759": &Movie{Title: "Star Wars: A New Hope", Rating: "8.7", Year: "1977"},
		"tt0082971": &Movie{Title: "Indiana Jones: Raiders of the Lost Ark", Rating: "8.6", Year: "1981"},
	}

	outgoingJSON, error := json.Marshal(movies)

	if error != nil {
		log.Println(error.Error())
		http.Error(res, error.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(res, string(outgoingJSON))
}

type Film struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Score       string `json:"rt_score"`
}

func handleFilms(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	respo, err := http.Get("https://ghibliapi.herokuapp.com/films")

	if err != nil {
		log.Fatal(err)
	}

	defer respo.Body.Close()

	body, err := ioutil.ReadAll(respo.Body)

	if err != nil {
		log.Fatal(err)
	}

	var films []Film
	jsonError := json.Unmarshal(body, &films)

	if jsonError != nil {
		log.Fatal(jsonError)
	}

	outgoingJSON, error := json.Marshal(films)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Fprint(res, string(outgoingJSON))
}
