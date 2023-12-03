package utils

import (
	"bufio"
	"log"
	"os"
)

func GetScanner(fileLoc string) (*bufio.Scanner, *os.File) {
	file, err := os.Open(fileLoc)
	if err != nil {
		log.Fatal(err)
	}

	return bufio.NewScanner(file), file
}

type Set[T comparable] struct {
	data map[T]struct{}
}

func NewSet[T comparable](insertValues ...T) Set[T] {
	s := Set[T]{
		data: map[T]struct{}{},
	}

	s.Add(insertValues...)
	return s
}

func (s *Set[T]) Add(insertValues ...T) {
	for _, t := range insertValues {
		s.data[t] = struct{}{}
	}
}

func (s *Set[T]) Exists(key T) bool {
	if _, ok := s.data[key]; ok {
		return true
	}
	return false
}

func (s *Set[T]) Delete(key T) {
	delete(s.data, key)
}
