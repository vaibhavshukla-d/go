package main

import "fmt"

//need a node
/*
	*		Node {
	*			data : "test"
	*			nxtPointer	: // address of next node
    *		 }
*/
//

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func FirstLinkedListProgram(){
	fmt.Println("Tes")

}

func (l *LinkedList)InsertAtBegining (data int) *Node {
	newNode := &Node{data: data,next :nil}
	l.head = newNode 
}