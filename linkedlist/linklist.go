package main

import "fmt"

//Definition for singly-linked list.
 type ListNode struct {
    Val int
  Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	current := &ListNode{}
	//currentVal := l1.Val
	c1 := l1
	c2 := l2
	carry := 0
	//current := listNode

	for c1!=nil || c2!=nil{
		x := c1.Val
		y := c2.Val
		sum := carry +x +y
		carry = sum/10

		fmt.Println(sum,carry)
		//listNode.Val = sum
		//listNode.Next
		tempNode := &ListNode{sum,nil}
		if current.Val==0{
			current = tempNode
		}else{
			//next := current.Next
			nextNode := current
			for nextNode.Next!=nil{
				nextNode = nextNode.Next
			}
			nextNode.Next =tempNode
			//nextNode :=  current.Next
			//if nextNode.Next==nil{
			//	nextNode.Next = tempNode
			//}
			//if nextNode.Next!=nil {
			//	//latNode = nextNode.Next
			//	nextNode = nextNode.Next
			//}
			//nextNode.Next = tempNode

		}

		c1 =c1.Next
		c2 = c2.Next

		//addTwoNumbers(l1,l2)
	}



	return current
}
func PrintNodes(l *ListNode){
	for l.Next!=nil{
		fmt.Println(l.Val)
		l = l.Next
	}
	fmt.Println(l.Val)
}

func main(){


	l1:=&ListNode{2,&ListNode{3,&ListNode{4,nil}}}
	l4:=&ListNode{2,&ListNode{3,&ListNode{4,nil}}}

	//PrintNodes(l1)
	PrintNodes(addTwoNumbers(l1,l4))

}