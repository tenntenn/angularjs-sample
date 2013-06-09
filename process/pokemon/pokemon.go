package pokemon

import (
	"appengine"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"log"
)

type Pokemon struct {
	Id int64 `json: "id"`
	Name string `json: "name"`
}

func init() {
	http.HandleFunc("/pokemon", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	val, err := handlePokemons(r)
	if err == nil {
		log.Printf("val=%s", val)
		err = json.NewEncoder(w).Encode(val)
		if err != nil {
			c.Errorf("%s", err)
		}
	}
	if err != nil {
		c.Errorf("pokemon error: %#v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handlePokemons(r *http.Request) (interface{}, error) {
	switch r.Method {
	case "GET":
		id, _ := strconv.ParseInt(r.FormValue("id"), 10, 64)
		return getPokemons(id)
	}
	return nil, fmt.Errorf("method not implemented")
}


func getPokemons(id int64) ([3]Pokemon, error) {
	pokemons := [3]Pokemon{}

	charmander := Pokemon{Id: 1, Name: "フシギダネ"}
	pokemons[0] = charmander
	pokemons[1] = charmander
	pokemons[2] = charmander
	return pokemons, nil
}
