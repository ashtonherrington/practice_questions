package practice_question

import "fmt"

func main() {
	var highestCount int
	input := "abcdeffg"
	for i := 0; i < len(input)-1; i++ {
		for j := len(input); j > i; j-- {
			if distinctChars(input[i:j]) <= 3 && len(input[i:j]) > highestCount {
				highestCount = len(input[i:j])
				fmt.Println(input[i:j], " has highest count of ", highestCount)
			}
		}
	}
	fmt.Println(highestCount)
}

func distinctChars(input string) int {
	var uniqueChars int
	var readChars []byte
	for _, character := range []byte(input) {
		if !containsByte(readChars, character) {
			readChars = append(readChars, character)
			uniqueChars++
		}
	}
	return uniqueChars
}

func containsByte(s []byte, e byte) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
