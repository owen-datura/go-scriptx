package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func main() {
	// TODO can the 'flags' package be used in lieu of os here?
	if len(os.Args) < 2 {
		// ensure that a filename's been provided
		fmt.Println("An input filename must be specified before a hash can be created.")
		os.Exit(1)
	}

	fn := os.Args[1]
	inputFile, oerr := os.Open(fn)
	if oerr != nil {
		fmt.Println(oerr)
		os.Exit(1)
	}
	defer inputFile.Close()

	// use Stat to determine if we're handling a directory
	fs, _ := inputFile.Stat()
	if fs.IsDir() {
		// throw an error until we revise this to handle dirs explicitly
		fmt.Println("Can't operate on a directory. Please specify a specific file.")
		os.Exit(1)
	}

	fmt.Println(generateHash(inputFile))
	//fmt.Printf("%+v\n",fs)
}

func generateHash(f *os.File) string {
	// first create the hashing object...
	hm := sha1.New()
	// then feed it the data present in the input file
	if _, err := io.Copy(hm, f); err != nil {
		// throw an error (exception?)
	}

	/*
		calling Sum forces the hash to be computed. note that the param
		supplied to the hasher (a byte array) can be used to populate,
		compute and output a hash in one operation.

		we then call EncodeToString to force the raw bytes into a format
		suitable for output.
	*/
	return hex.EncodeToString(hm.Sum(nil))
}
