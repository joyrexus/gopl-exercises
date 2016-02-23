package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)  // counts of words

	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "wordcount: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	
	input := bufio.NewScanner(file)
	input.Split(bufio.ScanWords) // split on words instead of lines

	for input.Scan() {
		word := input.Text()
		counts[word]++
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "wordcount: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("word\tfreq\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
}
