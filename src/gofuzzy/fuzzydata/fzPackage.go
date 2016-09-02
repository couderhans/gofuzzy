package fuzzydata

import (
	"archive/zip"
	"fmt"
	"log"
	"strings"
)

func ReadZip(zipfile string) (n string, err error) {

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
			log.Fatal(err)
		}
		if strings.Contains(extensions[1], "members") {
			fmt.Println("Creating member group: ", extensions[0])
			//line, err := fuzzydata.ReadFile(fileName)

		}
		rc.Close()
		fmt.Println()
	}

	//	defer conn.Close()

	return n, err

}
