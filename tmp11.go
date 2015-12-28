package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	var buf bytes.Buffer
	logger := log.New(&buf, "logger: ", log.Lshortfile)
	logger.SetOutput(os.Stderr)
	logger.Print("Hello, log file!")

	fmt.Print(&buf)

	fmt.Fprintln(os.Stderr, "KMZ")
}
