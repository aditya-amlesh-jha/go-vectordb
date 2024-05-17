package internal

type ListNode struct {
	next *ListNode
	prev *ListNode
	data string
}

type List interface {
	AddNode(data string) *ListNode
	DeleteNode(node *ListNode) *ListNode
	GetHead() *ListNode
	GetTail() *ListNode
}

type DoublyLinkedList struct {
	head *ListNode
	tail *ListNode
}

func (list *DoublyLinkedList) AddNode(data string) *ListNode {
	newNode := &ListNode{data: data}
	if list.tail == nil {
		list.head = newNode
		list.tail = newNode
	} else if list.tail != nil {
		list.tail.next = newNode
		newNode.prev = list.tail
		list.tail = newNode
	}
	return newNode
}

func (list *DoublyLinkedList) DeleteNode(node *ListNode) *ListNode {

	// check if node is nil
	if node == nil {
		return node
	}

	if node.prev != nil {
		node.prev.next = node.next
	} else {
		list.head = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		list.tail = node.prev
	}

	node.next = nil
	node.prev = nil

	return node
}

func (list *DoublyLinkedList) GetHead() *ListNode {
	return list.head
}

func (list *DoublyLinkedList) GetTail() *ListNode {
	return list.tail
}
