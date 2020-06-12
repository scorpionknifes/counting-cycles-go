package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		// Uncomment for testing multiple files with input
		createTxt("input/test1", "output/test1")
		createTxt("input/test2", "output/test2")
		os.Exit(1)
	*/

	/* default CONST used for running a single file
	 	inputFile - input path/file name e.g. test1.txt or test1
		outputFile - output path/file name e.g. output.txt or output
	*/
	var (
		inputFile  = ""
		outputFile = ""
	)

	if inputFile == "" || outputFile == "" {
		// if no CONST defined so using arguments
		if len(os.Args) < 3 {
			fmt.Println("Missing parameters")
			fmt.Println(" e.g. hw4.exe test1.txt output1.txt")
			fmt.Println("      hw4.exe test1 output1")
			fmt.Println("      hw4.exe input/test1 output/test1")
			os.Exit(1)
		}
		inputFile = os.Args[1]
		outputFile = os.Args[2]
	}

	// create output.txt using inputFile path and outputFile path
	createTxt(inputFile, outputFile)

}
