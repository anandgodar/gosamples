package main

import "fmt"

func reverse(bytes []byte, low int, high int)  {
	for low<high {
		bytes[low], bytes[high] = bytes[high], bytes[low]
		low++
		high--
	}
}

func reverseWords(s string) string {
	bytes := []byte(s)
	n := len(s)
	low := 0
	high := 0

	for high<n {
		for high<n && s[high]!= ' ' {
			high++
		}
		reverse(bytes, low, high-1)
		low = high+1
		high++
	}

	return string(bytes)
}

func main(){
	fmt.Println(reverseWords("this is a nepal"))
}
