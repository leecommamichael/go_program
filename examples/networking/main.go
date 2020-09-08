package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type pokemon struct {
	Name string
	URL  string
}

type pokeAPIResponse struct {
	Count    int
	Next     string
	Previous string
	Results  []pokemon
}

func main() {
	resp, _ := http.Get("https://pokeapi.co/api/v2/pokemon")
	body, _ := ioutil.ReadAll(resp.Body)
	var decodedBody pokeAPIResponse

	err := json.Unmarshal(body, &decodedBody)
	if err != nil {
		println(err.Error())
		os.Exit(0)
	}
	println(fmt.Sprintf(`There are %d pokemon in total and %d on this page.`, decodedBody.Count, len(decodedBody.Results)))
	println(fmt.Sprintf(`The next page is at: %s`, decodedBody.Next))
	for i, pokemon := range decodedBody.Results {
		println(fmt.Sprintf(`	#%d: %s`, i+1, pokemon.Name))
	}
}
