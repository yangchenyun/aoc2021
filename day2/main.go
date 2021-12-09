package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Dir string

const (
	Forward Dir = "forward"
	Down        = "down"
	Up          = "up"
)

type move struct {
	dir  Dir
	unit int
}

func main() {
	h := 0
	d := 0
	moves := processInput("day2.txt")

	for _, m := range(moves) {
		switch {
		case m.dir == Forward:
			h += m.unit
		case m.dir == Down:
			d += m.unit
		case m.dir == Up:
			d -= m.unit
		}
	}
	fmt.Printf("Result 1: %v\n", h * d)
}


func processInput(f string) []move {
    file, err := os.Open(f)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
	moves := make([]move, 0)
    for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		v, _ := strconv.Atoi(parts[1])
		m := move{}
		m.dir = Dir(parts[0])
		m.unit = v
		moves = append(moves, m)
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
	return moves
}
