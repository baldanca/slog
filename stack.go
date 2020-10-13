package slog

import (
	"fmt"
	"io"

	jsoniter "github.com/json-iterator/go"
)

type (
	// Stack model
	Stack struct {
		data [][]byte
	}
)

// GetAll stack function
func (s *Stack) GetAll() [][]byte {
	return s.data
}

// GetAllString stack function
func (s *Stack) GetAllString() []string {
	stackString := []string{}
	for _, value := range s.data {
		stackString = append(stackString, string(value))
	}
	return stackString
}

// PrintAll stack function
func (s *Stack) PrintAll() {
	for _, value := range s.data {
		fmt.Printf(string(value))
	}
}

// Save stack
func (s *Stack) Save(w io.Writer) {
	data, err := jsoniter.Marshal(s)
	if err != nil {
		panic(err)
	}
	w.Write(data)
}

// Write a log to stack
func (s Stack) Write(p []byte) (int, error) {
	s.data = append(s.data, p)
	return len(p), nil
}
