package practice_question

import "fmt"

func main() {
	results := FindSubSliceWithSum([]int32{3, 2, 1}, 3)
	fmt.Println(results, "distinct sets:", len(results))

	results = FindSubSliceWithSum([]int32{2, 2, -4, 1, 1, 2}, -3)
	fmt.Println(results, "distinct sets:", len(results))

	results = FindSubSliceWithSum([]int32{2, 2, 2}, 4)
	fmt.Println(results, "distinct sets:", len(results))
}

func FindSubSliceWithSum(completeSlice []int32, expectedValue int32) [][]int32 {

	var foundSubArrays [][]int32
	for index, value := range completeSlice {
		// if this is the final index, you cannot +1 to it, so compare directly
		if value == expectedValue {
			foundSubArrays = append(foundSubArrays, []int32{value})
			continue
		}
		if index == len(completeSlice)-1 {
			continue
		}
		runningSum := value
		for index2 := index + 1; index2 < len(completeSlice); index2++ {
			runningSum = runningSum + completeSlice[index2]
			if runningSum == expectedValue {
				foundSubArrays = append(foundSubArrays, completeSlice[index:index2+1])
			}
		}
	}
	return foundSubArrays
}
