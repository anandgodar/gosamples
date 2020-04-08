package main

import (
	"fmt"
	"strings"
)

func main() {
	f := func(c rune) bool {
		vowels := []rune{'a', 'e'}
		for _, value := range vowels {
			if c == value {
				return true
			}
		}
		return false
	}
	str := "this is a boy"
	index := strings.IndexFunc(str, f)
	fmt.Println(index)
}
