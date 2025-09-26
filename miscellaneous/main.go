// Write your answer here, and then test your code.
// Your job is to implement the Add(index int, v T) method.

package main

import (
	"fmt"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = true
const showHints = true

type Node[T any] struct {
	value      T
	next, prev *Node[T]
}

type DoublyLinkedList[T any] struct {
	size int
	head *Node[T]
	tail *Node[T]
}

// Add appends the given value at the given index.
// Returns an error in the case that the index exceeds the list size.
func (l *DoublyLinkedList[T]) Add(index int, v T) error {
	if l.head == nil {
		nodeToAdd := &Node[T]{value: v}
		l.head = nodeToAdd
		l.size++
		return nil
	}

	targetNode := l.head
	for i := 0; i < l.size; i++ {
		if index == i {
			break
		}
		if targetNode == nil {
			break
		}
		targetNode = targetNode.next
	}

	nodeToAdd := &Node[T]{value: v}
	if targetNode == l.head {
		nodeToAdd.next = targetNode
		targetNode.prev = nodeToAdd
		l.head = nodeToAdd
		l.size++
		return nil
	}

	if targetNode.next == nil {
		nodeToAdd.prev = targetNode.prev
		nodeToAdd.next = targetNode
		targetNode.prev = nodeToAdd
		targetNode.prev.next = nodeToAdd
		l.tail = targetNode
		l.size++
		return nil
	}

	//if nodeToAdd.next == targetNode.next

	//l.size++

	return nil
}

func (l *DoublyLinkedList[T]) AddElements(elements []struct {
	index int
	value T
}) error {
	for _, e := range elements {
		if err := l.Add(e.index, e.value); err != nil {
			return err
		}
	}

	return nil
}

func (l *DoublyLinkedList[T]) PrintForward() string {
	if l.size == 0 {
		return ""
	}
	current := l.head
	output := "HEAD"
	for current != nil {
		output = fmt.Sprintf("%s -> %v", output, current.value)
		current = current.next
	}

	return fmt.Sprintf("%s -> NULL", output)
}

func (l *DoublyLinkedList[T]) PrintReverse() string {
	if l.size == 0 {
		return ""
	}
	current := l.tail
	output := "NULL"
	for current != nil {
		output = fmt.Sprintf("%s <- %v", output, current.value)
		current = current.prev
	}
	return fmt.Sprintf("%s <- HEAD", output)
}

func main() {
	testCases := []struct {
		index int
		value string
	}{
		{index: 0, value: "C"},
		{index: 0, value: "A"},
		{index: 1, value: "B"},
		{index: 3, value: "D"},
	}
	dl := &DoublyLinkedList[string]{}
	err := dl.AddElements(testCases)
	if err != nil {
		fmt.Printf("Error adding elements: %v\n", err)
		return
	}
	forwardPrint := dl.PrintForward()
	reversePrint := dl.PrintReverse()
	fmt.Println(forwardPrint)
	fmt.Println(reversePrint)

}
