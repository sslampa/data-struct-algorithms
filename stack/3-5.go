package main

import (
	"errors"
	"fmt"
)

type Stack []string

func main() {
	fmt.Println(revString("Hello"))
}

func (s *Stack) Push(v string) {
	*s = append(*s, v)
}

func (s *Stack) Pop() (string, error) {
	len := len(*s)
	if len == 0 {
		return "", errors.New("Empty Stack")
	}

	last := (*s)[len-1]
	*s = (*s)[:len-1]
	return last, nil
}

func (s *Stack) Peek() (string, error) {
	len := len(*s)
	if len == 0 {
		return "", errors.New("Empty Stack")
	}

	last := (*s)[len-1]
	return last, nil
}

func (s *Stack) isEmpty() bool {
	len := len(*s)
	if len == 0 {
		return true
	}
	return false
}

func (s *Stack) Size() int {
	return len(*s)
}

func revString(mystr string) string {
	var s Stack

	for _, str := range mystr {
		s.Push(string(str))
	}

	var newStr string
	for !s.isEmpty() {
		insertChar, _ := s.Pop()
		newStr += insertChar
	}

	return newStr
}
