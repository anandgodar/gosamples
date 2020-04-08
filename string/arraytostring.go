package main

import (
	"fmt"
	"strings"
)

func main() {
	arrStr := []string{"one", "two", "three"}
	str := strings.Join(arrStr, ",")
	fmt.Println(str)
}
