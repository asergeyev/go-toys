package main

import (
	"compress/gzip"
	"io"
	"log"
	"os"
)

func main() {
	r, err := os.Open("./test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	w, err := os.Create("./test.txt.gz")
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()

	zw, err := gzip.NewWriterLevel(w, gzip.BestCompression)
	if err != nil {
		log.Fatal(err)
	}
	defer zw.Close()

	_, err = io.Copy(zw, r)
	if err != nil {
		log.Fatal(err)
	}
}
