package main

import "fmt"

func main(){
	//var ar1 [3]*string
	ar2 := [3]*string{new(string),new(string),new(string)}
	*ar2[0]= "red"
	*ar2[1] = "green"
	*ar2[2] = "white"
	fmt.Println(*ar2[1])

}
