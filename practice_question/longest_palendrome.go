package practice_question

import (
	"fmt"
)

func main() {
	input := "babad"

	var longestPalindrome int
	var longestPalindromeValue string
	for i := 0; i < len(input)-1; i++ {
		for j := len(input); j > i; j-- {
			if isPalindrome(input[i:j]) {
				fmt.Println("Found a palindrome!: ", input[i:j])
				if len(input[i:j]) > longestPalindrome {
					longestPalindrome = len(input[i:j])
					longestPalindromeValue = input[i:j]
				}
			}
		}
	}
	fmt.Println(longestPalindrome, "value: ", longestPalindromeValue)
}

func isPalindrome(input string) bool {
	if input == Reverse(input) {
		return true
	}
	return false
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
