package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Operation struct {
	InputPath  string
	OutputPath string
	Operation  string
	Array      []int
}

func (o *Operation) SortArray() {
	sort.Ints(o.Array)
}

func (o *Operation) ReverseArray() {
	for i, j := 0, len(o.Array)-1; i < j; i, j = i+1, j-1 {
		o.Array[i], o.Array[j] = o.Array[j], o.Array[i]
	}
}

func (o *Operation) SwapEvenOdd() {
	for i := 0; i < len(o.Array)-1; i += 2 {
		o.Array[i], o.Array[i+1] = o.Array[i+1], o.Array[i]
	}
}

func performOperation(op *Operation) error {
	if op.Operation == "sort" {
		op.SortArray()
	} else if op.Operation == "reverse" {
		op.ReverseArray()
	} else if op.Operation == "swap" {
		op.SwapEvenOdd()
	} else {
		return errors.New("unsupported operation")
	}
	return nil
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run main.go <input_file> <output_file> <operation>")
		return
	}

	inputPath := os.Args[1]
	outputPath := os.Args[2]
	operation := os.Args[3]

	file, err := os.Open(inputPath)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	inputArray := strings.Fields(scanner.Text())
	array := make([]int, 0, len(inputArray))
	for _, numStr := range inputArray {
		num, _ := strconv.Atoi(numStr)
		array = append(array, num)
	}

	op := &Operation{
		InputPath:  inputPath,
		OutputPath: outputPath,
		Operation:  operation,
		Array:      array,
	}

	err = performOperation(op)
	if err != nil {
		fmt.Println("Error performing operation:", err)
		return
	}

	outputFile, err := os.Create(outputPath)
	if err != nil {
		fmt.Println("Error creating output fil==e:", err)
		return
	}
	defer outputFile.Close()

}
