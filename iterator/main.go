package main

import (
	"fmt"
	"iter"
)

func main() {
	// for i := range iter1 {
	// 	fmt.Printf("%d hello", i)
	// }

	var s1 Slice
	s1 = []int{1, 2, 4, 5, 10}

	for i2 := range s1.All() {
		fmt.Println("hello", i2)
	}

	for i3 := range s1.Filter(func(i int) bool {
		return i != 10
	}) {
		fmt.Println("hello2 ", i3)
	}

	var s2 SliceStr
	s2 = []string{"ab", "cd", "ef"}

	for i := range s2.All() {
		fmt.Println("string ", i)
	}

	for i := range s2.Filter(func(i string) bool {
		return i != "cd"
	}) {
		fmt.Println("string2 ", i)
	}

	next, stop := iter.Pull(s2.All())
	defer stop()

	for {
		result, valid := next()
		if !valid {
			break
		}

		fmt.Println("iterator ", result)
	}
}

func iter1(yield func(i int) bool) {
	for i := range 3 {
		if !yield(i) {
			return
		}
	}
}

type Slice []int

func (s Slice) All() func(yield func(i int) bool) {
	return func(yield func(i int) bool) {
		for i := range s {
			if !yield(s[i]) {
				return
			}
		}
	}
}

func (s Slice) Filter(filerFunc func(i int) bool) func(yield func(i int) bool) {
	return func(yield func(i int) bool) {
		for i := range s {
			if filerFunc(s[i]) {
				if !yield(s[i]) {
					return
				}
			}
		}
	}
}

type SliceStr []string

func (s SliceStr) All() func(yield func(i string) bool) {
	return func(yield func(i string) bool) {
		for i := range s {
			if !yield(s[i]) {
				return
			}
		}
	}
}

func (s SliceStr) Filter(filerFunc func(i string) bool) func(yield func(i string) bool) {
	return func(yield func(i string) bool) {
		for i := range s {
			if filerFunc(s[i]) {
				if !yield(s[i]) {
					return
				}
			}
		}
	}
}
