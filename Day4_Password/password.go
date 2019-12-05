package main

import (
	"fmt"
	"strconv"
	"strings"
)

// check if the numbers are in crescent order
func isCrescent(code []int) bool {
	valid := 0
	if code[1] >= code[0] {
		valid++
	}
	if code[2] >= code[1] {
		valid++
	}
	if code[3] >= code[2] {
		valid++
	}
	if code[4] >= code[3] {
		valid++
	}
	if code[5] >= code[4] {
		valid++
	}
	if valid == 5 {
		return true
	}
	return false
}

// check if there are two adiacent numbers
func hasAdiacent(code []int) bool {
	valid := 0
	if code[1] == code[0] {
		valid++
	}
	if code[2] == code[1] {
		valid++
	}
	if code[3] == code[2] {
		valid++
	}
	if code[4] == code[3] {
		valid++
	}
	if code[5] == code[4] {
		valid++
	}
	if valid >= 1 {
		return true
	}
	return false

}

// convert an int into an array of int
func numToArray(i int) []int {
	numToString := strconv.Itoa(i)
	stringToArray := strings.Split(numToString, "")
	arrayOfInts := []int{}

	for _, j := range stringToArray {
		n, err := strconv.Atoi(j)
		if err != nil {
			panic(err)
		}
		arrayOfInts = append(arrayOfInts, n)
	}
	return arrayOfInts
}

func main() {

	start := 152085
	end := 670283

	var code []int
	numOfComb := 0

	for i := start; i < end; i++ {
		code = numToArray(i)
		if isCrescent(code) && hasAdiacent(code) {
			numOfComb++
		}
	}
	fmt.Printf("combination %d \n", numOfComb)
}
