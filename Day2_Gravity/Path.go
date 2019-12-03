package main

import (
	"encoding/csv"
	"fmt" // file readeer
	"log"
	"math"
	"os"
	"strconv"
)

// error check function
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func fuelToStar(distance int) int {
	fuel := int(math.Round(float64(distance/3))) - 2
	return fuel
}

func main() {

	Intcode, err := os.Open("input.csv") //read the file
	check(err)

	p := csv.NewReader(Intcode) //manage the file as csv

	codeS, err := p.Read()
	var code [1000]int
	for i := 0; i < len(codeS); i++ {
		var _ string
		code[i], _ = strconv.Atoi(codeS[i])
	}

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

	fmt.Print("Gravity has been solved ", code[0], "\n")
}
