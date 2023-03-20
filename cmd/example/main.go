package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/junkjuk/architecture-lab-2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "Input file")
	outputFile      = flag.String("o", "", "Output file")
)

func main() {
	flag.Parse()

	var input io.Reader

	if *inputExpression != "" && *inputFile != "" {
		fmt.Fprintln(os.Stderr, "too many input flags")
	}

	if *inputExpression != "" {
		input = strings.NewReader(*inputExpression)
	} else if *inputFile != "" {
		file, err := os.Open(*inputFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, "input file error: "+err.Error())
		}
		defer file.Close()
		input = file
	} else {
		fmt.Fprintln(os.Stderr, "input error")
		return
	}

	var output io.Writer
	if *outputFile != "" {
		file, err := os.Create(*outputFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, "output file error: "+err.Error())
		}
		defer file.Close()
		output = file
	} else {
		output = os.Stdout
	}

	handler := lab2.ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "convertation error: "+err.Error())
	}
}
