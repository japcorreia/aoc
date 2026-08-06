package main

import (
	"fmt"
)

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(value T) {
	s.elements = append(s.elements, value)
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("stack is empty")
	}
	index := len(s.elements) - 1
	value := s.elements[index]

	var zero T
	s.elements[index] = zero
	s.elements = s.elements[:index]

	return value, nil
}

func (s *Stack[T]) UpdateTop(newValue T) error {
	if s.IsEmpty() {
		return fmt.Errorf("stack is empty")
	}
	s.elements[len(s.elements)-1] = newValue
	return nil
}

func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("stack is empty")
	}
	return s.elements[len(s.elements)-1], nil
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}
