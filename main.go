package main

import (
	"io"
	"log"
	"os"
)

func main() {

	//os package has a function called Args returns args slice of type string string.
	//Grab second element in slice: name of textfile that we want to open
	file, err := os.Open(os.Args[1]) // For read access. *File is a pointer returned from os.Open
	//if error, log error
	if err != nil {
		log.Fatal(err)

		//exit programme if error
		os.Exit(1)
	}

	//call the Copy function from io package, passing into the standard output function from io package that will write
	//contents to terminal on operating system.
	//Second argument is the Reader function from io package and passing in the file that will by copied into the standard output function
	io.Copy(os.Stdout, io.Reader(file))

	//io.copy will take returned vlaue from the reader function and pipe into stdout channel to display in terminal.
	//Note: *File implements the Reader function from io package.
}
