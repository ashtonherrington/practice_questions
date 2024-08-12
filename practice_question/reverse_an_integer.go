package practice_question

import (
	"fmt"
	"strconv"
)

func main() {
	t := strconv.Itoa(123456789)

	byteString := []byte(t)

	var newString []byte
	for i := len(byteString) - 1; i >= 0; i-- {
		newString = append(newString, byteString[i])
	}

	reversedInt, err := strconv.Atoi(string(newString))
	if err != nil {
		fmt.Println(fmt.Errorf("converting string to integer: %w", err))
	}

	fmt.Println(reversedInt)
}
