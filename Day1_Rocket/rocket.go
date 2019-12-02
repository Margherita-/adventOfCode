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

		fuel := int(math.Round(float64(distance/3))) - 2
		totalFuel = fuel + totalFuel
	}

	fmt.Printf("Santa needs %d liters of fuel \n", totalFuel)

}
