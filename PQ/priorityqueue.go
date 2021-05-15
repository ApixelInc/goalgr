package PQ

import (
	"container/heap"
	// "fmt"
)

// An InternalItem is something we manage in a priority queue.
type InternalItem struct {
	Value    interface{} // The value of the item; arbitrary.
	Priority int    // The priority of the item in the queue.
	// The index is needed by update(heap.Fix) and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap
    count int // for make sort stable
}

// A PQ_ARRAY implements heap.Interface and holds Items.
type PQ_ARRAY []*InternalItem

func (pq PQ_ARRAY) Len() int { return len(pq) }

func (pq PQ_ARRAY) Less(i, j int) bool {
    // tie break
    if pq[i].Priority == pq[j].Priority {
        return pq[i].count < pq[j].count
    }
    // use `less than` for minimum heap, or use `greater than` for maximum heap
	return pq[i].Priority < pq[j].Priority
}

func (pq PQ_ARRAY) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ_ARRAY) Push(x interface{}) {
	n := len(*pq)
	item := x.(*InternalItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PQ_ARRAY) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an InternalItem in the queue.
func (pq *PQ_ARRAY) update(item *InternalItem, value interface{}, priority int) {
	item.Value = value
	item.Priority = priority
    // call heap method
	heap.Fix(pq, item.index)
}

// ===============================================

type PriorityQueue struct {
    queue PQ_ARRAY
    global_count int
}

func New() *PriorityQueue {
    queue := PQ_ARRAY {}
    // call heap method
    heap.Init(&queue)
    return &PriorityQueue{ queue: queue , global_count: 0 }
}

// return internal _ITEM
func (self *PriorityQueue) Push( v interface{}, priority int ) interface{} {
    x := &InternalItem{ Value: v, Priority: priority, count:self.global_count  }

    self.global_count ++

    // call heap method
    heap.Push( &self.queue, x )

    return x
}

func (self *PriorityQueue) Pop() interface{} {
    // call heap method
    return heap.Pop(&self.queue).(*InternalItem).Value
}

func (self *PriorityQueue) Update(item *InternalItem, value interface{}, priority int)  {
    self.queue.update(item , value,  priority )
}

func (self *PriorityQueue) IsEmpty() bool {
    return self.queue.Len() == 0
}




