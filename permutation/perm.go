package main

import (
	"fmt"
	"strings"
)

func main(){
	str:= "ABc"
 findPermutation(str,0,len(str))
 //fmt.Println("swap",swap("ABC",1,2))
}

func findPermutation(str string,start,end int){
	if start == end {
		fmt.Println(str)
		//strs = append(strs,str)
	}else{
		for i:=start;i<end;i++{
			swappedStr := swap(str,start,i)
			findPermutation(swappedStr,start+1,end)
		}
	}
}


func swap(str string,start,end int)string{
	strArray := strings.Split(str,"")
	strArray[start],strArray[end]=strArray[end],strArray[start]
	return strings.Join(strArray,"")
}