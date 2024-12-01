package main

import (
	"bufio"
	"fmt"
	"os"
)

func Task2() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()
	outputFile, err := os.Create("output2.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()
	os.Stdin = inputFile
	os.Stdout = outputFile

	scanner := bufio.NewScanner(inputFile)

	var arr1, arr2 []int64

	for scanner.Scan() {
		var num1, num2 int64
		_, err := fmt.Sscanf(scanner.Text(), "%d %d", &num1, &num2)
		if err != nil {
			fmt.Println("Error scanning line:", err)
			continue
		}
		arr1 = append(arr1, num1)
		arr2 = append(arr2, num2)
	}

	var similarity int64
	frequencies := make(map[int64]int64)

	for _, num := range arr2 {
		frequencies[num]++
	}

	for _, num := range arr1 {
		similarity += num * frequencies[num]
	}

	fmt.Println(similarity)

}
