package linkedlist

type Node struct {
	data int
	next *Node
}
type LinkedList struct {
	head *Node
}

func (l *LinkedList) InsertAtBack(data int) {
	newNode := &Node{data: data, next: nil}
	if l.head == nil {
		l.head = newNode
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func CreateList() {}
