package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Node stores in and out in a hash set
type Node struct {
	in  map[int]bool
	out map[int]bool
}

// createTxt creates output.txt from input.txt
func createTxt(inputFile string, outputFile string) {
	// Reading file names
	input, err := readLines(checkTxt(inputFile))
	if err != nil {
		fmt.Println("Reading file name: " + inputFile + " error")
		os.Exit(1)
	}

	// Write file names
	err = writeLines(getOutput(input), checkTxt(outputFile))
	if err != nil {
		fmt.Println("Writing to file name: " + outputFile + " error")
		os.Exit(1)
	}
	fmt.Println("Success: input " + inputFile + " has been output to " + outputFile)
}

// checkTxt check if filename has extension of .txt
func checkTxt(filename string) string {
	if filepath.Ext(strings.TrimSpace(filename)) != ".txt" {
		filename = filename + ".txt"
	}
	return filename
}

// divideInt divide two intergers as floats
func divideInt(i int, j int) float32 {
	return float32(i) / float32(j)
}

// printInts print []int without "[]"
func printInts(array []int) string {
	return strings.Trim(fmt.Sprint(array), "[]")
}
