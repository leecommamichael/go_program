package main

import "os"

func main() {
	var command = os.Args
	for _, arg := range command {
		println(arg)
	}
}
