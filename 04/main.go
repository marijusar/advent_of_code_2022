package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isWithin(a string, b string) bool {

	aPair := strings.Split(a, "-")
	bPair := strings.Split(b, "-")

	ab, _ := strconv.Atoi(aPair[0])
	ac, _ := strconv.Atoi(aPair[1])
	bc, _ := strconv.Atoi(bPair[0])
	bd, _ := strconv.Atoi(bPair[1])

	if ab <= bc && ac >= bd {
		fmt.Printf("%d <= % d , %d >= %d \n", ab, bc, ac, bd)
		return true
	}

	return false
}

func isOverlapping(a string, b string) bool {

	acc := make(map[int]bool)

	aPair := strings.Split(a, "-")
	bPair := strings.Split(b, "-")

	ab, _ := strconv.Atoi(aPair[0])
	ac, _ := strconv.Atoi(aPair[1])
	bc, _ := strconv.Atoi(bPair[0])
	bd, _ := strconv.Atoi(bPair[1])

	for i := ab; i <= ac; i++ {
		acc[i] = true
	}

	for i := bc; i <= bd; i++ {
		if _, ok := acc[i]; ok {
			return true
		}
	}

	return false
}

func main() {
	input, _ := os.ReadFile("./input.txt")

	pairs := strings.Split(string(input), "\n")

	var count int

	for i := 0; i < len(pairs)-1; i++ {
		pairArr := strings.Split(pairs[i], ",")
		a := pairArr[0]
		b := pairArr[1]

		if isOverlapping(a, b) {
			fmt.Printf("a : %s, b : %s\n", a, b)
			count++
		}
	}

	fmt.Println(count)
}
