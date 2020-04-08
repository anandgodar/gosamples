package channel

import (
	"fmt"
	"time"
)

func ChannelEx1(){
	ch := make(chan int)
	 go func(){
	 	ch<-1
	 }()
	val := <-ch

	fmt.Println(val)
}

func ChannelExample2(){
	for i:=0; i<24;i++{
		timerCH:= callTimer(1*time.Second)
		fmt.Println(<-timerCH)
	}
}

func callTimer(t time.Duration)<-chan int{
	timerCH := make(chan int)
	go func(){
		//time.Sleep(t)
		timerCH<-1
	}()
	return timerCH
}

func ChannelExample3(){
	var Ball int
	table := make(chan int)
	for i := 0; i < 100; i++ {
		go player(table,"p"+string(i))
	}
	table <- Ball
	time.Sleep(1* time.Second)
	<-table
}

func player(table chan int, pl string){
	for {
		ball := <-table
		ball++
		fmt.Println(ball,pl)
		time.Sleep(100 * time.Millisecond)
		table <- ball

	}
}