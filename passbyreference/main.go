package main

import "fmt"

type Student struct {
	name string
	age int
}

type Students struct {
	sts []Student
}
func (s *Student) update(name string){
	s.name = name
}

func (s Student) updatet(name string){
	s.name = name
}

func main(){
	st1 := Student{"Ram",7}
	st1.update("sgytsco")
	fmt.Println(st1)

	//sts := Students{}


	st := &Student{name:"Rabin"}
	st.update("arbing")
	st.updatet("df")
	fmt.Println(st)


}