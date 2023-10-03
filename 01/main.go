package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	path := os.Args[1]

	input, err := os.ReadFile(path)

	if err != nil {
		fmt.Printf("Error while reading file %e", err)
		return
	}

	inputStr := string(input)

	elveCalories := strings.Split(inputStr, "\n\n")

	var arr []int

	for i := 0; i < len(elveCalories); i++ {
		e := elveCalories[i]

		cArr := strings.Split(string(e), "\n")

		acc := 0

		for j := 0; j < len(cArr); j++ {
			strI, _ := strconv.Atoi(cArr[j])
			fmt.Println(cArr[j])
			acc = acc + strI
		}

		arr = append(arr, acc)
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})

	sum := arr[0] + arr[1] + arr[2]

	fmt.Println(sum)
}
