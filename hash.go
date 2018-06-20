package main

import (
	"os"
	"fmt"
)

func main() {
	// TODO can the 'flags' package be used in lieu of os here?
	if len(os.Args) < 2 {
		// ensure that a filename's been provided
		fmt.Println("An input filename must be specified before a hash can be created.")
		os.Exit(1)
	}

	// use Stat in lieu of Open here 'cause we'll need to know if
	// we're operating on a file or directory
	fs, serr := os.Stat(os.Args[1])
	if serr != nil {
		fmt.Println(serr)
		os.Exit(1)
	} else if fs.IsDir() {
		// the operation was a success but we were provided a directory.
		// throw an error for now; later we can handle this explicitly
		fmt.Println("Can't operate on a directory. Please specify a specific file.")
		os.Exit(1)
	}

	n := getDetails(fs)
	fmt.Println(n)
	//fmt.Printf("%+v\n",fs)
}

func getDetails(f os.FileInfo) string {
	return "hello"
}
