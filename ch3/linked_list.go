package main

import "fmt"

// Node структура узла
type Node struct {
	property int
	nextNode *Node
}

// LinkedList структура связного списка
type LinkedList struct {
	headNode *Node
}

// AddToHead добавляет узел со значением в начало списка
func (linkedList *LinkedList) AddToHead(property int) {
	var node = Node{}
	node.property = property
	if linkedList.headNode != nil {
		node.nextNode = linkedList.headNode
	}
	linkedList.headNode = &node
}

// IterateList итерируется по списку и выводит его содержимое на экран
func (linkedList *LinkedList) IterateList() {
	var node *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)
	}
}

// LastNode получает последний узел в списке
func (linkedList *LinkedList) LastNode() *Node {
	var node *Node
	var lastNode *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

// AddToEnd добавляет узел со значением в конец списка
func (linkedList *LinkedList) AddToEnd(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	var lastNode = linkedList.LastNode()
	if lastNode != nil {
		lastNode.nextNode = node
	}
}

// NodeWithValue получает узел по указанному значению
func (linkedList *LinkedList) NodeWithValue(property int) *Node {
	var node *Node
	var nodeWith *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			nodeWith = node
			break
		}
	}
	return nodeWith
}

// AddAfter добавляет узел со значением, следующий за узлом с переданным параметром
func (linkedList *LinkedList) AddAfter(nodeProperty, property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	var nodeWith = linkedList.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		nodeWith.nextNode = node
	}
}

func main() {
	linkedList := &LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	linkedList.AddToEnd(5)
	linkedList.AddAfter(1, 7)
	linkedList.IterateList()
}

// 3
// 1
// 7
// 5
