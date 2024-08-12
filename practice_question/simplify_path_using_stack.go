package practice_question

import (
	"errors"
	"fmt"
	"strings"
)

// Stack logic derived from Queue implementation here
// https://www.geeksforgeeks.org/queue-in-go-language/)

func main() {
	q := Stack{
		Size: 10,
	}

	// input := "/home/../etc"
	// input := "/home/./me"
	input := "/home//me/"

	parts := strings.Split(input, "/")
	for _, part := range parts {
		if part == "." || part == "" {
			continue
		} else if part == ".." {
			q.Pop()
		} else {
			q.Push(part)
		}
	}

	var sb strings.Builder
	for _, value := range q.Elements {
		sb.WriteString("/")
		sb.WriteString(value)
	}

	if sb.String() == "" {
		fmt.Println("/")
		return
	}

	fmt.Println(sb.String())
}

type Stack struct {
	Elements []string
	Size     int
}

func (q *Stack) Push(elem string) {
	if q.GetLength() == q.Size {
		fmt.Println("Overflow")
		return
	}
	q.Elements = append(q.Elements, elem)
}

func (q *Stack) Pop() string {
	if q.IsEmpty() {
		fmt.Println("UnderFlow")
		return ""
	}
	lastElement := q.Elements[len(q.Elements)-1]
	if q.GetLength() == 1 {
		q.Elements = nil
		return lastElement
	}
	q.Elements = q.Elements[:len(q.Elements)-2]
	return lastElement
}

func (q *Stack) GetLength() int {
	return len(q.Elements)
}

func (q *Stack) IsEmpty() bool {
	return len(q.Elements) == 0
}

func (q *Stack) Peek() (string, error) {
	if q.IsEmpty() {
		return "", errors.New("empty queue")
	}
	return q.Elements[0], nil
}
