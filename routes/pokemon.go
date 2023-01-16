package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

const apiUrl = "https://pokeapi.co/api/v2"

type species struct {
	Name string `json:"name"`
}

type pokemonInfo struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Height  int     `json:"height"`
	Weight  int     `json:"weight"`
	Species species `json:"species"`
}

func GetPokemon(w http.ResponseWriter, r *http.Request) {
	pokemonName := chi.URLParam(r, "name")
	url := fmt.Sprintf("%s/pokemon/%s", apiUrl, pokemonName)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	pokemonObj := pokemonInfo{}
	err = json.Unmarshal(body, &pokemonObj)
	if err != nil {
		log.Fatal(err)
	}

	pokemonJson, err := json.Marshal(pokemonObj)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Add("content-type", "application/json")
	io.WriteString(w, string(pokemonJson))
}
