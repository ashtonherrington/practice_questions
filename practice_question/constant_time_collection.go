package practice_question

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	collection := NewO1Collection()
	fmt.Println(collection.Add("1"))
	fmt.Println(collection.Add("1"))
	fmt.Println(collection.Add("2"))
	fmt.Println(collection.Add("3"))
	fmt.Println(collection.GetRandom())

	fmt.Println(collection.Delete("1"))
	fmt.Println(collection.Delete("1"))
}

type O1Collection struct {
	valueStore   []string
	indexTracker map[string]int
}

func NewO1Collection() O1Collection {
	return O1Collection{
		indexTracker: make(map[string]int),
	}
}

func (o *O1Collection) Add(value string) bool {
	if containsString(o.valueStore, value) {
		return true
	}

	o.valueStore = append(o.valueStore, value)
	o.indexTracker[value] = len(o.valueStore) - 1 // inserted into final index
	return false
}

func (o *O1Collection) GetRandom() string {
	rand.Seed(time.Now().UnixNano())

	// Generate a random Int type number between 1 and 10
	randNum := rand.Intn(len(o.valueStore))
	return o.valueStore[randNum]
}

func (o *O1Collection) Delete(value string) bool {
	finalIndex := len(o.valueStore) - 1
	valueIndex, ok := o.indexTracker[value]
	if !ok || valueIndex == -1 {
		return false
	}
	//handle case where value index is 0, last index, or in between
	if valueIndex == finalIndex {
		o.valueStore = o.valueStore[:len(o.valueStore)-1]
	} else if valueIndex == 0 {
		o.valueStore = o.valueStore[1:]
	} else {
		o.valueStore = append(o.valueStore[:valueIndex], o.valueStore[valueIndex+1:]...)
	}
	o.indexTracker[value] = -1
	return true
}

func containsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
