package main

import "io/ioutil"

func main() {
	newFilename := "./memory.txt"
	poem := "Old pond\nA frog jumps in –\nThe sound of water"
	bytes := []byte(poem)

	ioutil.WriteFile(newFilename, bytes, 0644)
}
