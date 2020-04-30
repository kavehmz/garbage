package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

/*
type item struct {
	line string
	prev *item
}
*/

func main() {
	var fileName = flag.String("file", "", "filename")
	var nLine = flag.Int("n", 10, "number of lines")
	flag.Parse()

	if *nLine == 0 {
		return
	}

	log.Println("Parsing:", *fileName)

	file, err := os.Open(*fileName)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	index := -1
	for scanner.Scan() {
		index++
		if index >= *nLine {
			index = 0
		}
		if len(lines) < *nLine {
			lines = append(lines, scanner.Text())
		} else {
			lines[index] = scanner.Text()
		}
	}

	count := 0
	for count < len(lines) {
		index++
		if index >= len(lines) {
			index = 0
		}
		fmt.Println(lines[index])
		count++
	}

	// i := len(lines) - *nLine
	// if i < 0 {
	// 	i = 0
	// }
	// for ; i <= len(lines)-1; i++ {
	// 	fmt.Println(lines[i])
	// }

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
