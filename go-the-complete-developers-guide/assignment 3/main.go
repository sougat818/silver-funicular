package main

import "os"

func main() {
	filename := os.Args[1]
	data, err := os.ReadFile(filename)
	if err != nil {
		println("Error reading file", filename, err.Error())
		os.Exit(1)
	}
	println(string(data))

}
