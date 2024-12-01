package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func Task1() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()
	outputFile, err := os.Create("output1.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()
	os.Stdin = inputFile
	os.Stdout = outputFile

	scanner := bufio.NewScanner(inputFile)

	var arr1, arr2 []float64
	var ans float64

	for scanner.Scan() {
		var num1, num2 float64
		_, err := fmt.Sscanf(scanner.Text(), "%f %f", &num1, &num2)
		if err != nil {
			fmt.Println("Error scanning line:", err)
			continue
		}
		arr1 = append(arr1, num1)
		arr2 = append(arr2, num2)
	}

	sort.Float64s(arr1)
	sort.Float64s(arr2)

	for i := 0; i < len(arr1); i++ {
		ans += math.Abs(arr1[i] - arr2[i])
	}

	fmt.Println(int(ans))

}
