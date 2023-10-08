package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.ReadFile("./input.txt")

	fs := string(f)

	inputs := strings.Split(fs, "\n\n")

	grid := inputs[0]

	rows := strings.Split(grid, "\n")

	indexPositionMap := make(map[int]int)

	lastItem := rows[len(rows)-1]

	for i := 0; i < len(lastItem); i++ {

		if string(lastItem[i]) != " " {
			integer, _ := strconv.Atoi(string(lastItem[i]))
			indexPositionMap[i] = integer
		}
	}

	rowsWithoutNumeration := rows[:len(rows)-1]

	columns := make(map[int][]string)

	for i := 0; i < len(rowsWithoutNumeration); i++ {
		for j := 0; j < len(rowsWithoutNumeration[i]); j++ {

			val, ok := indexPositionMap[j]

			if ok {
				if string(rowsWithoutNumeration[i][j]) != " " {
					columns[val] = append(columns[val], string(rowsWithoutNumeration[i][j]))
				}
			}
		}

	}

	instructions := strings.Split(inputs[1], "\n")
	re := regexp.MustCompile("[0-9]+")

	var instructionSet [][]int

	for i := 0; i < len(instructions)-1; i++ {
		v := instructions[i]
		s := re.FindAllString(v, -1)

		var intSlice []int
		for _, st := range s {
			v, err := strconv.Atoi(st)
			if err == nil {
				intSlice = append(intSlice, v)
			}
		}
		instructionSet = append(instructionSet, intSlice)
	}

	for i := 0; i < len(instructionSet); i++ {
		itemCnt := instructionSet[i][0]
		from := instructionSet[i][1]
		to := instructionSet[i][2]

		colFrom := columns[from]
		colTo := columns[to]

		leftover := colFrom[itemCnt:]

		items := colFrom[0:itemCnt]
		newCol := append(items, colTo...)
		fmt.Printf("__start %d __", i)
		fmt.Println(colFrom, itemCnt, items, leftover)
		fmt.Printf("__end %d __", i)

		columns[from] = leftover
		columns[to] = newCol

		// fmt.Println(columns)

	}

	fmt.Println(columns)

}
