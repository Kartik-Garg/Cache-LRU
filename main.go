package main

import (
	"fmt"
)

const SIZE = 5

type Node struct{
	Val string
	Left *Node
	Right *Node
}

type Queue struct{
	Head *Node
	Tail *Node
	Length int
}

type Cache struct{
	Queue Queue
	Hash Hash
}

type Hash map[string] *Node

func NewCache() Cache{
	return Cache{Queue: NewQueue(), Hash: Hash{}}
}

func NewQueue() Queue{
	//takes empty struct of type node
	head := &Node{}
	tail := &Node{}
	//this basically denotes that the queue is empty as head and tail are always empty and if they
	//point towards each other, then it means that nothing is present in middle
	head.Right = tail
	tail.Left = head

	return Queue{Head: head, Tail: tail}
}

func (c *Cache) Check(str string){
	node := &Node{}

	//checking if the value is present in the hash declared, if it does, we remove it and add it to the
	//beginning of the Queue, else we just add
	if val, ok := c.Hash[str]; ok{
		node = c.Remove(val)
	}else{
		node = &Node{Val: str}
	}
	//adding the node at the beginning of the Q in cache and also in hash to maintain
	c.Add(node)
	c.Hash[str] = node
}

func (c *Cache) Remove(n *Node) *Node{
	fmt.Printf("Remove %s\n", n.Val)
	//these left and right of the node which we are removing
	left := n.Left
	right := n.Right

	//elements on left and right have to point towards each other, removing the deleted node 
	//from the picture

	left.Right = right
	right.Left = left
	//decrement the length of Q after removing the node
	c.Queue.Length -= 1
	//delete(built in) the value from hash and delte the value
	delete(c.Hash, n.Val)
	return n
}

func (c *Cache) Add(n *Node){
	fmt.Printf("add: %s\n", n.Val)
	tmp := c.Queue.Head.Right

	c.Queue.Head.Right = n
	n.Left = c.Queue.Head
	n.Right = tmp
	tmp.Left = n
	c.Queue.Length++

	//if the size of the cache increases more than what is already defined then we remove the last 
	//element
	if c.Queue.Length > SIZE{
		c.Remove(c.Queue.Tail.Left)
	}

}

func (c *Cache) Display(){
	c.Queue.Display()
}

func (q *Queue) Display(){
	node := q.Head.Right
	fmt.Printf("%d - [", q.Length)
	for i:=0 ; i<q.Length; i++ {
		fmt.Printf("{%s}", node.Val)
		if i<q.Length-1{
			fmt.Printf("<-->")
		}
		node = node.Right
	}
	fmt.Println("]")
}

func main(){
	fmt.Println("START CACHE")
	cache := NewCache()
	//taking slice of array of string
	for _, word := range [] string{"parrot", "avocado", "tree", "potato", "tree"}{
		cache.Check(word)
		cache.Display()
	}
}