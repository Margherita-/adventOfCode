package main

import (
	"encoding/csv"
	"fmt" // file readeer
	"log"
	"os"
	"strconv"
)

// error check function
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func calculus(code [1000]int) int {
	oP := 0
	aP := 1
	bP := 2
	dP := 3

	opcode := code[oP]
	positionA := code[aP]
	positionB := code[bP]
	destination := code[dP]

Going:
	for {
		switch opcode {
		case 99:
			break Going
		case 1:
			code[destination] = code[positionA] + code[positionB]
		case 2:
			code[destination] = code[positionA] * code[positionB]
		default:
			fmt.Printf("This is wrong, we are all going to die! \n")
		}

		oP = oP + 4
		aP = aP + 4
		bP = bP + 4
		dP = dP + 4

		opcode = code[oP]
		positionA = code[aP]
		positionB = code[bP]
		destination = code[dP]
	}

	return code[0]
}

func main() {

	Intcode, err := os.Open("input.csv") //read the file
	check(err)

	p := csv.NewReader(Intcode) //manage the file as csv

	codeS, err := p.Read()
	var original [1000]int
	for i := 0; i < len(codeS); i++ {
		var _ string
		original[i], _ = strconv.Atoi(codeS[i])
	}

	var code [1000]int
	noun, verb := 0, 0

	for i := 0; i < 10000; i++ {

		code = original
		code[1] = noun
		code[2] = verb

		result := calculus(code)

		if result == 19690720 {
			fmt.Printf("Gravity equation is solved! %d%d \n", noun, verb)
			break
		}

		// adding 1 to each index, one at a time

		if i%100 == 0 {
			noun = noun + 1
			verb = 0
		} else {
			verb = verb + 1
		}

	}

}
