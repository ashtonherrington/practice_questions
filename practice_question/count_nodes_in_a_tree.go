package practice_question

import (
	"fmt"
	"sort"
)

func main() {

	slice1 := []int{7, 8, 3, 10, 9, 1, 11, 2, 12, 4, 5, 6}
	slice2 := []int{2, 7, 10, 11, 12, 6, 9, 1, 3, 4, 5}

	sort.Ints(slice1)
	sort.Ints(slice2)

	fmt.Println(slice1)
	fmt.Println(slice2)

	for i := range slice1 {
		if slice1[i] == slice2[i] {
			continue
		}
		fmt.Println(slice1[i], "is missing")
		break
	}
}
