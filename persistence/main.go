package main

import "io/ioutil"

func main() {
	bytes := []byte("I'd like to remember this moment.")
	ioutil.WriteFile("./hey", bytes, 0644)
}
