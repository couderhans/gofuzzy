package main

import (
	"archive/zip"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
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

	if err != nil {
		log_and_exit(err)
	}

	fmt.Printf("Loading fuzzy sets from zipfile: %v", zipfile)

	// Open the zip archive containing the fuzzy set data.
	r, err := zip.OpenReader(zipfile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("zip file opened")
	defer r.Close()

	// Iterate through the files in the archive,
	// If a member file, create member list of the contents
	for _, f := range r.File {
		fileName := f.Name
		extensions := strings.Split(fileName, ".")
		fmt.Printf("Contents of %s:\n", fileName)
		rc, err := f.Open()
		if err != nil {
			log_and_exit(err)
		}
		if strings.Contains(extensions[1], "members") {
			fmt.Println("Creating member group %v", extensions[0])

		}
		rc.Close()
		fmt.Println()
	}

	if err != nil {
		log_and_exit(err)
	}

	//	defer conn.Close()

}
