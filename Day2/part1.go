package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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

	var ans int = 0

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		inputStr := scanner.Text()

		parts := strings.Split(inputStr, " ")
		numbers := make([]int, len(parts))
		for i, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println("Error converting to int:", err)
				return
			}
			numbers[i] = num
		}

		var max_diff, min_diff int = math.MinInt32, math.MaxInt32

		for i := 0; i < len(numbers)-1; i++ {
			var diff = numbers[i] - numbers[i+1]
			max_diff = int(math.Max(float64(max_diff), float64(diff)))
			min_diff = int(math.Min(float64(min_diff), float64(diff)))
		}

		fmt.Println(max_diff, min_diff)

		if (max_diff <= 3 && min_diff >= 1) || (max_diff <= -1 && min_diff >= -3) {
			ans++
		}
	}

	fmt.Println(ans)

}
