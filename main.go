package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json");
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json");
	params := mux.Vars(r);

	for _, item := range movies {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json");
	params := mux.Vars(r);

	for index, item := range movies {
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json");
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r);

	for index, item := range movies {
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)

			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			break
		}
	}
}


func main(){
	r := mux.NewRouter()

	movies = append(movies, Movie{ID:"1", Isbn:"1234", Title:"movie 1", Director: &Director{FirstName: "vansh", LastName: "Patel"}})
	movies = append(movies, Movie{ID:"2", Isbn:"14", Title:"movie 2", Director: &Director{FirstName: "vansh", LastName: "Patel"}})
	movies = append(movies, Movie{ID:"3", Isbn:"134", Title:"movie 3", Director: &Director{FirstName: "vansh", LastName: "Patel"}})
	movies = append(movies, Movie{ID:"4", Isbn:"24", Title:"movie 4", Director: &Director{FirstName: "vansh", LastName: "Patel"}})
	movies = append(movies, Movie{ID:"5", Isbn:"123", Title:"movie 5", Director: &Director{FirstName: "vansh", LastName: "Patel"}})
	movies = append(movies, Movie{ID:"6", Isbn:"34", Title:"movie 6", Director: &Director{FirstName: "vansh", LastName: "Patel"}})


	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

	fmt.Printf("staring server on 8000");
	log.Fatal(http.ListenAndServe(":8000",r))
}