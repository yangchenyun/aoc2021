package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// read numbers from input
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
	for i := 0; i + 1 < len(numbers); i++ {
		if numbers[i + 1] > numbers[i]  {
			inc_count += 1
		}
	}
	fmt.Printf("Part 1 Result: %v\n", inc_count)

	windows := make([]int, 0)
	for i := 0; i + 2 < len(numbers); i++ {
		windows = append(windows, (numbers[i] + numbers[i + 1] + numbers[i + 2]))
	}

	inc_count = 0
	for i := 0; i + 1 < len(windows); i++ {
		if windows[i + 1] > windows[i]  {
			inc_count += 1
		}
	}
	fmt.Printf("Part 2 Result: %v\n", inc_count)
}
