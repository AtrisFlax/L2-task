package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	outputFile := getOutputFile()
	url := getUrl()

	// Send HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Can't perform GET request\n")
		os.Exit(1)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	// Create output file
	file, err := os.Create(*outputFile)
	if err != nil {
		fmt.Println("Can't create output file")
		os.Exit(1)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	// Copy response to output file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Can't copy response body to file")
		os.Exit(1)
	}

	fmt.Printf("Saving to: %s\n", *outputFile)
}

func getUrl() string {
	return flag.Arg(0)
}

func getOutputFile() *string {
	outputFile := flag.String("O", "index.html", "Output file")
	flag.Parse()
	return outputFile
}
