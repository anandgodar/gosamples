package main

import (
	"fmt"
	"strings"
)

func repeatedSubstringPattern(s string) bool {
	sub1 := s[1:]
	sub2 := s[:len(s)-1]
	bigdick := sub1 + sub2
	if strings.Contains(bigdick,s){
		return true
	}else{
		return false
	}
}
func main(){
	fmt.Println(repeatedSubstringPattern("aa"))

	// abcabc

	// ramram  -> amramramra
}
