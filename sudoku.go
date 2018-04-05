// Copyright 2014 Zanicar. All rights reserved.

// Utilizes a BSD-3 license. Refer to the included LICENSE file for details.

// Package main implements the dlx package to solve Sudoku problems. It is
// terminal based and allows Sudoku sizes of 2, 4, 9, 16 and 25.
package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/zanicar/dlx"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Invalid usage parameter:")
		fmt.Println("\tCorrect usage is: sudoku {size} {string}")
		fmt.Println("\t{size} denotes the board size, e.g. 9")
		fmt.Println("\t{string} denotes the initial solution and must be of length (size X size)")
		return
	}
	s, err := strconv.Atoi(os.Args[1])
	if err != nil || (s != 2 && s != 4 && s != 9 && s != 16 && s != 25) {
		fmt.Printf("Invalid usage parameter: %s\n", os.Args[1])
		fmt.Println("\tCorrect usage is: sudoku {size} {string}")
		fmt.Println("\t{size} must be a valid integer of value {2,4,9,16,25}")
		return
	}
	init := os.Args[2]
	if len(init) != s*s {
		fmt.Println("Invalid usage parameter: {string}")
		fmt.Println("\tCorrect usage is: sudoku {size} {string}")
		fmt.Println("\t{string} must be of length (size X size)")
		return
	}
	m := createSudokuMatrix(s, init)
	t0 := time.Now()
	solutions := m.Solve()
	t1 := time.Now()
	t := t1.Sub(t0)
	solstrs := convertDLXSolutions(solutions)

	fmt.Println()
	fmt.Println("Initial Solution:")
	printSolution(init, s)
	fmt.Println()
	printSolutions(solstrs, s)
	fmt.Printf("\nFound %d solutions in %f seconds.\n", len(solstrs), t.Seconds())
}

// CreateSudokuMatrix creates a matrix of size^2 (size x size) and initializes
// it with preset values given in string is (initial solution); Integers denote
// filled values, whereas "." represents nil values. It returns a pointer to
// the newly created matrix.
func createSudokuMatrix(size int, is string) *dlx.Matrix {
	sizeSqr := size * size
	sizeSqrt := int(math.Sqrt(float64(size)))
	m := dlx.New()
	createMatrixHeader(size, m)

	// Traverse each cell position
	// p - cell position
	// r - row
	// c - column
	// b - block

	boardSize := sizeSqr
	var re *dlx.Element
	var hpv, hrv, hcv, hbv, hev string
	for p := 0; p < boardSize; p++ {
		r := p / size
		c := p % size
		b := p/(size*sizeSqrt)*sizeSqrt + p/sizeSqrt%sizeSqrt
		hpv = fmt.Sprintf("p%v", p)
		// Check initial solution
		isv := string(is[p])
		if isv != "." {
			// A single row in the dlx matrix indicating the initial solution
			hrv = fmt.Sprintf("r%vv%v", r, isv)
			hcv = fmt.Sprintf("c%vv%v", c, isv)
			hbv = fmt.Sprintf("b%vv%v", b, isv)
			// Set row element nil, indicating a new row
			re = nil
			for ce := m.Head(); ce != nil; ce = ce.Right() {
				hev = ce.Value.(dlx.Head).Name()
				if hev == hpv || hev == hrv || hev == hcv || hev == hbv {
					re = m.PushItem(re, ce)
				}
			}
		} else {
			// Traverse each option value
			for v := 0; v < size; v++ {
				// Every iteration of (p,v) represents a row in the dlx matrix
				vs := getValueString(v)
				hrv = fmt.Sprintf("r%vv%v", r, vs)
				hcv = fmt.Sprintf("c%vv%v", c, vs)
				hbv = fmt.Sprintf("b%vv%v", b, vs)
				// Set row element nil, indicating a new row
				re = nil
				for ce := m.Head(); ce != nil; ce = ce.Right() {
					hev = ce.Value.(dlx.Head).Name()
					if hev == hpv || hev == hrv || hev == hcv || hev == hbv {
						re = m.PushItem(re, ce)
					}
				}
			}
		}
	}
	return m
}

// CreateMatrixHeader initializes the constraints, namely the cell position
// constraint, the row constraint, the column constraint and the block
// constraint.
func createMatrixHeader(size int, m *dlx.Matrix) {
	cSize := size * size * 4
	for i := 0; i < cSize; i++ {
		switch {
		// cell position constraint
		case i < size*size:
			{
				p := fmt.Sprintf("p%v", i)
				m.PushHead(p)
				break
			}
		// row constraint
		case i < (size * size * 2):
			{
				r := (i - size*size) / size
				v := (i - size*size) % size
				vs := getValueString(v)
				p := fmt.Sprintf("r%vv%v", r, vs)
				m.PushHead(p)
				break
			}
		// column constraint
		case i < (size * size * 3):
			{
				c := (i - size*size*2) / size
				v := (i - size*size*2) % size
				vs := getValueString(v)
				p := fmt.Sprintf("c%vv%v", c, vs)
				m.PushHead(p)
				break
			}
		// block constraint
		case i < (size * size * 4):
			{
				b := (i - size*size*3) / size
				v := (i - size*size*3) % size
				vs := getValueString(v)
				p := fmt.Sprintf("b%vv%v", b, vs)
				m.PushHead(p)
				break
			}
		}
	}
}

// Returns an appropriate value.
func getValueString(v int) string {
	arr := [25]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O"}
	return arr[v]
}

// Converts the exact cover problem solutions, which are in the form of a slice
// of strings (each string representing a valid option according to the problem
// constraints; and the number of strings exactly matching the space of the
// problem scope), to a single string denoting the cell position and its
// corresponding zero indexed value. The function thus returns a slice of
// stings each representing a valid solution according to the given input.
func convertDLXSolutions(s [][]string) []string {
	str := make([]string, len(s))
	for sc := range s {
		sol := make([]rune, len(s[sc]))
		for rc := range s[sc] {
			preg := regexp.MustCompile(`^.*p(\d+).*$`)
			ps := preg.ReplaceAllString(s[sc][rc], "$1")
			p, _ := strconv.Atoi(ps)
			preg = regexp.MustCompile(`^.*v(.).*$`)
			v := preg.ReplaceAllString(s[sc][rc], "$1")
			sol[p] = rune(v[0])
		}
		str[sc] = string(sol)
	}
	return str
}

// Iterates through all discovered solutions and prints them to console.
func printSolutions(s []string, size int) {
	for i := range s {
		fmt.Printf("\nSolution %d:\n", i+1)
		printSolution(s[i], size)
		fmt.Println()
	}
}

// Prints the solution in a neatly human readable format.
func printSolution(s string, size int) {
	sizeSqr := size * size
	sizeSqrt := int(math.Sqrt(float64(size)))
	for p := 0; p < sizeSqr; p++ {
		vs := s[p : p+1]
		if size < 10 {
			v, err := strconv.Atoi(vs)
			v++
			if err != nil {
				vs = " "
			} else {
				vs = strconv.Itoa(v)
			}
		}
		switch {
		case p%(size*sizeSqrt) == 0 && p/(size*sizeSqrt) != 0:
			{
				fmt.Println()
				for i := 0; i < sizeSqrt; i++ {
					for j := 0; j < sizeSqrt; j++ {
						fmt.Print(" -")
					}
					if (i+1)%sizeSqrt != 0 {
						fmt.Print(" +")
					}
				}
				fmt.Println()
				break
			}
		case p%size == 0:
			{
				fmt.Println()
			}
		case p%sizeSqrt == 0:
			{
				fmt.Print(" |")
				break
			}
		}
		if vs == "." {
			fmt.Print("  ")
		} else {
			fmt.Printf(" %s", vs)
		}
	}
}
