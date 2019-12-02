package main

import (
	"encoding/csv"
	"fmt"
	"io" // file readeer
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

	stars, err := os.Open("input.csv") //read the file
	check(err)

	r := csv.NewReader(stars) //manage the file as csv

	totalFuel := 0

	for {
		star, err := r.Read()
		if err == io.EOF {
			break
		}

		check(err)
		distance, _ := strconv.Atoi(star[0])

		for {
			if distance < 9 {
				break
			}

			distance = fuelToStar(distance)
			totalFuel = totalFuel + distance
		}
	}

	fmt.Printf("Santa needs %d liters of fuel \n", totalFuel)
}
