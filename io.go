package main

import (
	"bufio"
	"fmt"
	"os"
)

// getOutput from nodes
func getOutput(input []string) []string {
	t := NewNodes(input)
	in, out := t.AverageDegree()
	output := []string{
		printInts(t.SameNode()),
		fmt.Sprint(in, out),
	}
	order, ordered := t.FindOrder()
	end := []string{}
	if ordered {
		end = []string{"Ordered:", printInts(order), "Yes"}
	} else {
		cycles, yes := t.FindAllCycle()
		end = []string{"Cycle(s):"}
		for _, cycle := range cycles {
			end = append(end, printInts(cycle))
		}
		end = append(end, yes)
	}
	return append(output, end...)
}

// readLines reads a whole file and scan lines
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines writes the lines to the given path.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}
