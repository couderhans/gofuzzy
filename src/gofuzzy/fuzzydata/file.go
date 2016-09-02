package fuzzydata

import (
	"io"
	"os"
)

func ReadFile(source string) (n string, err error) {

	// open input file
	fi, err := os.Open(source)
	if err != nil {
		panic(err)
	}
	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	// make a buffer to keep chunks that are read
	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
	}
	n = string(buf)
	return n, err
}

func WriteToFile(destination string, buf []byte) {

	// open output file
	fo, err := os.Create(destination)
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	// write a chunk
	if _, err := fo.Write(buf); err != nil {
		panic(err)
	}
}
