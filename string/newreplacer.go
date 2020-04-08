package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "men and women are equal but they are bad"

	replacer := strings.NewReplacer("men", "women", "women", "men", "bad", "good")
	fmt.Println(str)
	fmt.Println(replacer.Replace(str))
}
