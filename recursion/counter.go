package main

import "fmt"

func counter(s []int) (r int) {
	if len(s)==0{
		return 0
	}
	_,s=s[len(s)-1],s[:len(s)-1]
	r++
	fmt.Println(r)
	return r+counter(s)
}

func main(){
	x:= []int{3,4,5,6,7}
	fmt.Println(counter(x))
}