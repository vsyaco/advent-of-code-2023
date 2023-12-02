package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func calibrationString(st string) string {

	var intsStrs []string

	for _, c := range st {
		if unicode.IsDigit(c) {
			intsStrs = append(intsStrs, string(c))
		}
	}

	if len(intsStrs) == 0 {
		return "0"
	}

	if len(intsStrs) == 1 {
		return intsStrs[0] + intsStrs[0]
	}

	return intsStrs[0] + intsStrs[len(intsStrs)-1]
}

func main() {

	var sum int64

	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal("Error while reading file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		text := scanner.Text()

		if text == "" {
			continue
		}

		calibrationString := calibrationString(text)

		integer, err := strconv.ParseInt(calibrationString, 10, 64)
		if err != nil {
			continue
		}

		sum += integer

	}

	fmt.Println(sum)
}
