package main

import "fmt"

func main() {
	// Raw string literal
	str := `this is raw \n string literal`
	fmt.Println(str)
	//Interpolated String literal

	iStr := "this is raw \n string literal"
	fmt.Println(iStr)
}
