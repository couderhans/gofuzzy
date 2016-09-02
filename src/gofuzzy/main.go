package main

import (
	"errors"
	"fmt"
	"gofuzzy/fuzzydata"
	"log"
	"os"
)

func parse_args() (zipfile string, err error) {
	args := os.Args[1:]
	if len(args) != 1 {
		return "", errors.New("Usage: ./run zipfile")
	}
	zipfile = args[0]
	return
}

func log_and_exit(err error) {
	log.Fatal(err)
	os.Exit(1)
}

func main() {

	zipfile, err := parse_args()

	fmt.Println("Reading fzPackage: ", zipfile)
	result, err := fuzzydata.ReadZip(zipfile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("zip file opened", result)

}
