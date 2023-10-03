package main

import (
	"fmt"
	"os"
	"strings"
)

func getBadge(group []string) string {
	i := group[0]

	for _, character := range i {
		if characterExists(group[1], string(character)) && characterExists(group[2], string(character)) {
			return string(character)
		}
	}

	return ""
}

func characterExists(word string, char string) bool {
	for _, wordCharacter := range word {
		if string(wordCharacter) == char {
			return true
		}

	}

	return false
}

func createAlphabet() string {

	alphabet := make([]byte, 0, 26)

	// Loop over 'A' to 'Z' and keep on appending
	// to alphabet slice
	var ch byte
	for ch = 'A'; ch <= 'Z'; ch++ {
		alphabet = append(alphabet, ch)
	}

	// Print alphabet
	return string(alphabet)
}

func main() {
	file, _ := os.ReadFile("./input.txt")

	inputs := strings.Split(string(file), "\n")

	var badges []string

	for i := 0; i < len(inputs)-1; i = i + 3 {
		group := inputs[i : i+3]

		badge := getBadge(group)

		fmt.Println(badge)

		badges = append(badges, badge)
	}

	alphabet := strings.ToLower(createAlphabet()) + createAlphabet()
	var score int

	for _, v := range badges {
		for i, r := range alphabet {
			if v == string(r) {
				score = score + i + 1
			}
		}

	}

	fmt.Println(score)

}
