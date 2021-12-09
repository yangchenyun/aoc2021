package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("day1.txt")
	depths := strings.Split(string(input), "\n")
	numbers := make([]int, 0)
	for _, d := range(depths) {
		v, _ := strconv.Atoi(d)
		if v > 0 {
			numbers = append(numbers, v)
		}
	}
	inc_count := 0
	for i := 0; i + 1 < (len(numbers)); i++ {
		if (i + 1) == len(input) - 1 {
			break
		}
		// fmt.Printf("Compare: %v > %v?\n", input[i+1], input[i])
		if numbers[i + 1] > numbers[i]  {
			inc_count += 1
		}
	}
	fmt.Printf("Result: %v", inc_count)
}
