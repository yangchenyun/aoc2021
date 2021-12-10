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
	numbers := make([]uint64, 0)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		s := scanner.Text()
		leng = len(s)
		n, _ := strconv.ParseUint(s, 2, leng)
		numbers = append(numbers, n)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	m := mask(leng)
	digits := make([]uint16, leng)

	for i := leng - 1; i >= 0; i-- {
		one_c := 0
		zero_c := 0
		for _, n := range(numbers) {
			if (n>>i) & 1 == 0b1 {
				one_c++
			} else {
				zero_c++
			}
		}
		// NOTE: digits saved in order
		if one_c > zero_c {
			digits = append(digits, 1)
		} else if zero_c > one_c {
			digits = append(digits, 0)
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
	gamma := int(r)
	epsilon := int(^r&m)
	fmt.Printf("%b, %b, %v", gamma, epsilon, gamma*epsilon)
}

func mask(l int) uint16 {
	var r uint16
	for i := 0; i < l; i++ {
		r = (r<<1)+1
	}
	return r
}
