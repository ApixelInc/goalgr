package main

import (
    "github.com/ApixelInc/goalgr/PQ"
    "container/heap"
    "fmt"
)

func priorityqueue_test() {
    // Some items and their priorities.
    items := map[string]int{
        "banana": 3, "apple": 2, "pear": 4,
    }

    // create and init with existing items
    pq := make(PQ.PriorityQueue, len(items))
    i := 0
    for value, priority := range items {
        pq[i] = &PQ.Item{
            Value:    value,
            Priority: priority,
            Index:    i,
        }
        i++
    }
    heap.Init(&pq)

    // Insert a new item 
    item := &PQ.Item{
        Value:    "orange",
        Priority: 1,
    }
    heap.Push(&pq, item)

    // but we want change the priority...
    pq.Update(item, item.Value, 5)

    // Take the items out; they arrive in decreasing priority order.
    for pq.Len() > 0 {
        item := heap.Pop(&pq).(*PQ.Item)
        fmt.Printf("%.2d:%s ", item.Priority, item.Value)
    }
}
