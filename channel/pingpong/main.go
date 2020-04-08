package main

import (
	"fmt"
	"time"
)

func main(){
	var Ball int
	table := make(chan int)
	for i:=1;i<100;i++ {
		go player(table, i)
		//go player(table, "Amrita")
	}
	table<-Ball
	time.Sleep(5*time.Second)
	<-table
}

func player(table chan int , player int){
	for {
		value:=<-table
		value++
		fmt.Println("player" ,player," hit number : ",value)
		time.Sleep(100*time.Millisecond)
		table<-value
	}
}
