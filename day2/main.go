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

type loc struct {
	h int
	d int
	aim int
}

func (l *loc) Move1(m move) {
		switch {
		case m.dir == Forward:
			l.h += m.unit
		case m.dir == Down:
			l.d += m.unit
		case m.dir == Up:
			l.d -= m.unit
		}
}

func (l *loc) Result() int {
	return l.h * l.d
}

func (l *loc) Move2(m move) {
		switch {
		case m.dir == Forward:
			l.h += m.unit
			l.d += l.aim * m.unit
		case m.dir == Down:
			l.aim += m.unit
		case m.dir == Up:
			l.aim -= m.unit
		}
}

func main() {
	moves := processInput("day2.txt")

	l := loc{}
	for _, m := range(moves) {
		l.Move1(m)
	}
	fmt.Printf("Result 1: %v\n", l.Result())

	l = loc{}
	for _, m := range(moves) {
		l.Move2(m)
	}
	fmt.Printf("Result 2: %v\n", l.Result())
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
