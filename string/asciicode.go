package main

import "fmt"

func main() {
	str := "aA Bboy"
	//Print the ASCII char
	for _, s := range str {
		fmt.Println(s)
		fmt.Printf("%T",s)
	}
}
