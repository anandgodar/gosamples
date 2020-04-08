package main

import "fmt"

type Node struct {
	tag,text string
	children []*Node
	val int
}

func BFSTransversal(node *Node) (*Node,[]int){
	queue := make([]*Node,0)
	queue = append(queue,node)
	values := make([]int,0)
	for len(queue)>0{
		firstElQ := queue[0]
		queue = queue[1:]
		fmt.Println(firstElQ.text)
		values = append(values,firstElQ.val)
		if len(firstElQ.children)>0{
			for _,value:=range firstElQ.children{
				queue = append(queue,value)
			}
		}
	}
	return  nil,values
}


func findByTag(node *Node, tag string) *Node{
	queue := make([]*Node,0)
	queue = append(queue,node)
	for len(queue)>0{
		firstElement := queue[0]
		queue = queue[1:]
		if firstElement.tag==tag{
			return firstElement
		}
		if len(firstElement.children)>0{
			for _,value := range firstElement.children{
				queue = append(queue,value)
			}
		}
	}

	return nil
}
func main(){

	h1 := Node{tag:"h1",text:"H1",val:10}
	div := 	Node{tag:"div",text:"this is div",val:5,children:[]*Node{&h1}}
	span := 	Node{tag:"span",text:"this is span",val:4}
	para := 	Node{tag:"para",text:"this is para",val:2,children:[]*Node{&span}}
	body := Node{tag:"body",text:"this is body",val:2,children:[]*Node{&h1,&div,&para}}
	root := Node{tag:"root",text:"this is root",val:10,children:[]*Node{&body}}

	node := findByTag(&root,"div")
	fmt.Println(node)

	_,v:=BFSTransversal(&root)
	fmt.Println(v)

}
