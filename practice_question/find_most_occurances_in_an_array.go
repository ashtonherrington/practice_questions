package practice_question

import "fmt"

func main() {
	input := []int{1, 3, 5, 3, 3, 7, 1, 9}

	ih := NewIntHash(10)
	for _, value := range input {
		ih.Add(value)
	}

	var largestNumber *int
	for _, value := range input {
		if largestNumber == nil {
			largestNumber = &value
		}
		maybeNewLargest := ih.GetOccurances(value)
		if maybeNewLargest > *largestNumber {
			largestNumber = &maybeNewLargest
		}
	}

	fmt.Println(*largestNumber)
}
