package main

import (
	"fmt"
	"sort"
)

type Student struct {
	name string
	age int
}

//type customSort struct {
//	student []Student
//}

type mySort []Student
func (ms mySort) Len()int {return len(ms)}
func (ms mySort) Swap(i,j int){ms[i],ms[j]=ms[j],ms[i]}
func (ms mySort) Less(i,j int)bool{return ms[i].age<ms[j].age}

func main(){
	students:= []Student{{"Annad",12},{"Amrita",24},{"Kenzel",2}}
	sort.Sort(mySort(students))
	fmt.Println(students)
}
