package main

import (
	"container/list"
	"fmt"
)

type node struct{
	Id string
	friends map[string]*node
}

func traverseBFS(n *node)[]*node{
	queue := list.New()
	queue.PushBack(n)

	visited := make(map[string]*node)
	visited[n.Id] = n

	for queue.Len()>0{
		topNode := queue.Front()
		for id,node := range topNode.Value.(*node).friends{
			if _, ok := visited[id]; !ok {
				visited[id] = node
				queue.PushBack(node)
			}
		}
		queue.Remove(topNode)
	}
	nodes := make([]*node,0)
	//nodes := make([]*node,5)
	for _, node := range visited {
		nodes = append(nodes, node)
	}

	return nodes


}
 func main (){
 	node1 := &node{"node1",make(map[string]*node)}
	 node2 := &node{"node2",make(map[string]*node)}
	 node3 := &node{"node3",make(map[string]*node)}
	 node4 := &node{"node4",make(map[string]*node)}


	// node2.friends[node1.Id] = node1
	// node3.friends[node2.Id] = node2
	// node1.friends[node3.Id] = node3
	 node3.friends[node4.Id] = node4

	 root := &node{"root",make(map[string]*node)}

	 root.friends[node1.Id] = node1
	 root.friends[node2.Id] = node2
	 root.friends[node3.Id] = node3


	nodes := traverseBFS(root)

	 for _, node := range nodes {
		 fmt.Printf("%+v\n", node.Id)
	 }
 }