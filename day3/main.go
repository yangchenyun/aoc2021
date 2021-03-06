package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
    file, err := os.Open("day3.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	var leng int
	numbers := make([]uint16, 0)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		s := scanner.Text()
		leng = len(s)
		n, _ := strconv.ParseUint(s, 2, leng)
		numbers = append(numbers, uint16(n))
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	m := mask(leng)
	digits := make([]uint16, leng)
	gamma := ""
	epsilon := ""

	for i := leng - 1; i >= 0; i-- {
		one_c := 0
		zero_c := 0
		for _, n := range(numbers) {
			if (n>>i)&1 == 0b1 {
				one_c++
			} else {
				zero_c++
			}
		}
		// NOTE: digits saved in order
		if one_c > zero_c {
			digits = append(digits, 1)
			gamma += "1"
			epsilon += "0"
		} else if zero_c > one_c {
			digits = append(digits, 0)
			gamma += "0"
			epsilon += "1"
		} else {
			log.Printf("Unexpecte results %v, %v", zero_c, one_c)
		}
	}
	fmt.Printf("mask: %b\n", m)
	fmt.Printf("digits: %v\n", digits)

	// convert the digits to bits
	var r uint16
	for _, d := range(digits) {
		r = (r<<1)+d
	}

	g, _ := strconv.ParseInt(gamma, 2, 64)
	e, _ := strconv.ParseInt(epsilon, 2, 64)
	fmt.Printf("Result 1: %b, %b, %v\n", g, e, g*e)

	g2 := int(r)
	e2 := int(^r&m)
	fmt.Printf("Result 1: %b, %b, %v\n", g2, e2, g2*e2)

	oxy := filter(numbers, leng - 1, mostCommonFn)
	co2 := filter(numbers, leng - 1, leastCommonFn)
	fmt.Printf("Result 2: %v\n", int(oxy) * int(co2))

}

type keepFn func(int, int) int

var mostCommonFn keepFn = func(one_c, zero_c int) int {
	if one_c == zero_c {
		return 1
	} else if one_c > zero_c {
		return 1
	} else {
		return 0
	}
}
var leastCommonFn keepFn = func(one_c, zero_c int) int {
	if one_c == zero_c {
		return 0
	} else if one_c > zero_c {
		return 0
	} else {
		return 1
	}
}

func filter(numbers []uint16, i int, fn keepFn) uint16{
	if len(numbers) == 1 {
		return numbers[0]
	}

	one_c := 0
	zero_c := 0
	for _, n := range(numbers) {
		if (n>>i)&1 == 0b1 {
			one_c++
		} else {
			zero_c++
		}
	}

	new_numbers := make([]uint16, 0)
	for _, n := range(numbers) {
		if (n>>i)&1 == uint16(fn(one_c, zero_c)) {
			new_numbers = append(new_numbers, n)
		}
	}
	return filter(new_numbers, i - 1, fn)
}

func mask(l int) uint16 {
	var r uint16
	for i := 0; i < l; i++ {
		r = (r<<1)+1
	}
	return r
}
