package LLIST

import (
    "container/list"
)


type LinkedList struct {
    queue *list.List
}

func New() *LinkedList {
    return &LinkedList { queue: list.New() }
}

func (self *LinkedList) PushItem(x interface{}) interface{} {
    self.queue.PushBack( x )
    return x
}

// in order keep the interface in Graph Search
// for simple pusth, use PushItem instead
func (self *LinkedList) Push(x interface{}, noused int) interface{} {
    self.queue.PushBack( x )
    return x
}

func (self *LinkedList) Pop() interface{} {
    e := self.queue.Front() // First element
    self.queue.Remove(e)
    return e.Value
}

func (self *LinkedList) IsEmpty() bool {
    return self.queue.Len() == 0
}
