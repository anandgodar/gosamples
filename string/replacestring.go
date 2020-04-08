package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main(){
	fmt.Println(replaceString("force is Force"))
	fmt.Println(replacer("force is Force"))
}
func replaceString(str string) string{
	str1 := strings.Split(str,"")
	buf := bytes.Buffer{}
	//replaceStr := ""
	for _,v:=range str1{
		if v=="f" || v== "F"{
			buf.WriteString("K"+v)
		}else {
			buf.WriteString(v)
		}

	}

	return buf.String()
}

func replacer(value string)string{
	// Create replacer with pairs as arguments.
	r := strings.NewReplacer("f", "Kf",
		"F", "KF")

	// Replace all pairs.
	result := r.Replace(value)
	return result
}