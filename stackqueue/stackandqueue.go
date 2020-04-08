package main

import "fmt"

type Stack struct{
	nodes []*Node
	count int
}

type Node struct {
	val int
}

func (s *Stack) push(node *Node ){
	s.count = s.count+1
	s.nodes = append(s.nodes,node)
}

func (s *Stack) pop() *Node{
	if s.count==0{
		return nil
	}
	s.count--
	return s.nodes[s.count]

}


func newStack() *Stack{
	return &Stack{}
}
func enqueue(){

}

func deque()int {
	return 1
}

func main(){
	myStack := newStack()
	myStack.push(&Node{1})
	myStack.push(&Node{2})
	myStack.push(&Node{3})
	fmt.Println(myStack,myStack.pop(),myStack.pop(),myStack.pop())

}