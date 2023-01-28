package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	collection := []Data{}

	index := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineArray := strings.Split(line, "           ")
		if index != 0 {
			s, err := strconv.ParseFloat(lineArray[1], 64)
			if (err != nil) {
				fmt.Println(err)
			}
			collection = append(collection, Data{name: lineArray[0], score: s})
		}
		index += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Original: ", collection)

	Sorting(collection)
}
