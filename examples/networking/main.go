package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	resp, _ := http.Get("https://pokeapi.co/api/v2/pokemon")
	body, _ := ioutil.ReadAll(resp.Body)
	println(string(body))
}
