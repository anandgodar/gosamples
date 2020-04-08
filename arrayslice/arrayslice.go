package main
import (
	"fmt"
	"sync"
)

func increment(wg *sync.WaitGroup) {
	var x=0
	x = x + 1
	fmt.Println(x)
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	for i := 0; i < 100; i++ {
		w.Add(1)
		go increment(&w)
	}
	w.Wait()
	//fmt.Println("final value of x", x)
}

//package arrayslice
//import "fmt"

func ArraySlice(){
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := a[2:4]

	fmt.Println("address of a[2]", &a[2])
	fmt.Println("address of s[0]", &s[0])
	fmt.Println("&a[2] == &s[0] is", &a[2] == &s[0])
}