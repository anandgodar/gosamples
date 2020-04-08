package main

import "fmt"

func nonRepeatingChar(s string) int{
	var counts [128]int
	for _,v := range s {
		counts[v]++
	}
	for i,v := range s{
		if counts[v]==1{
			return i
		}
	}
	return -1
	}

func main(){
fmt.Println(nonRepeatingChar("hhello"))
}