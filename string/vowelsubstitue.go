package main

import (
	"fmt"
	"strings"
)

func IsVowel(c string)bool{
	//fmt.Println(c)
	vowels := []string{"a","e","i","o","u"}
	for _,v:=range vowels{
		if v== c{
			return true
		}
	}

	return false

}

func isVowels(c byte) bool {
	vowels := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}

	for _,v := range vowels{
		if c == v {
			return true
		}
	}

	return false
}
func replaceVowelInString(str string) string{
	s := strings.Split(str,"")
	high := len(str)-1
	low := 0
	ret:= make([]string,len(str))
	for low<=high{
		if IsVowel(s[low]) && IsVowel(s[high]){
			ret[low],ret[high]=s[high],s[low]
			low++
			high--
		}else if(!IsVowel(s[low])){
			ret[low] = s[low]
			low++
		}else if(!IsVowel(s[high])){
				ret[high]=s[high]
				high--
		}

	}
	return strings.Join(ret,"")
}

func main(){
	//str := strings.Split("hello","")
	//IsVowel(str[0])

	isVowels("dfd"[0])
	fmt.Println(replaceVowelInString("heisaboy"))
}
