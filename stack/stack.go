package main

import "fmt"

func main(){
	node:= NewStack()
	node.Push(1)
	node.Push(2)
	fmt.Println(node.Pop(),node.Pop())
}

type Stack interface {
	Push(v int)
	Pop() int
	Empty() bool
}

type Node struct {
	counter int
	value int
	nodes []*Node
}

func (n *Node) Push(v int){
	newNode := &Node{value:v}
	n.nodes = append(n.nodes,newNode)
	n.counter++
}

func (n *Node) Pop()int{
	if n.counter==0{
		return 0
	}
	topValue := n.nodes[len(n.nodes)-1].value
	n.nodes = n.nodes[0:len(n.nodes)-1]
	n.counter--
	return topValue
}

func (n *Node) Empty() bool{
	if n.counter==0 {
		return true
	}
	return false
}

func NewStack()*Node{
	return &Node{}
}
