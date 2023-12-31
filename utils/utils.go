package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
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

func NewSet[T comparable](insertValues ...T) *Set[T] {
	s := &Set[T]{
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

func (s *Set[T]) Intersection(s2 *Set[T]) *Set[T] {
	res := NewSet[T]()
	for k := range s.data {
		if ok := s2.Exists(k); ok {
			res.Add(k)
		}
	}

	return res
}

func (s *Set[T]) Union(s2 *Set[T]) *Set[T] {
	res := NewSet[T]()
	for key := range s2.data {
		res.Add(key)
	}
	for key := range s.data {
		res.Add(key)
	}

	return res
}

func (s *Set[T]) Len() int {
	return len(s.data)
}

func (s *Set[T]) Slice() []T {
	var res []T
	for k := range s.data {
		res = append(res, k)
	}

	return res
}

func IsNumber(s string) (int, bool) {
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, false
	}
	return v, true
}
