package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"

	"guess-it-2/calculation"
)

// Data Analysis Functions

func main() {
	window := flag.Int("window", 5, "Number of last numbers to consider for prediction")

	if *window <= 0 {
		fmt.Fprintln(os.Stderr, "Window size must be positive")
		os.Exit(1)
	}

	flag.Parse()

	var data []float64

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		value, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("Error parsing input")
			continue
		}
		data = append(data, value)
		if len(data) > *window {
			data = data[len(data)-*window:]
		}

		if len(data) > 1 {
			lower, upper := calculation.CalculateRange(data)
			fmt.Printf("%d %d\n", int(lower), int(upper))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing input: %v\n", err)
	}
}
