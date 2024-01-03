package main

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

type Hash map[string]*Node

type Cache struct {
	Queue Queue
	Hash  Hash
}

func NewCache() Cache {
	return Cache{Queue: NewQueue(), Hash: Hash{}}
}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}

	head.Right = tail
	tail.Left = head
	return Queue{Head: head, Tail: tail, Length: 0}
}

func (c *Cache) Check(str string) {
	node := &Node{}
}

func main() {
	cache := NewCache()
	for _, word := range []string{"egg", "meat", "fish"} {
		cache.Check(word)
		// cache.Display()
	}
}
