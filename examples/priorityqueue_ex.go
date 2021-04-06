package main

import (
    "github.com/ApixelInc/goalgr/PQ"
    "fmt"
)

func priorityqueue_test() {
    // Some items and their priorities.
    items := []string{
        "banana", "apple", "pear",
    }

    // create and init with existing items
    pq := PQ.New()

    for i, value := range items {
        item := &PQ.Item{
            Value:    value,
            Priority: i,
        }
        pq.Push( item )
    }

    item := &PQ.Item{
        Value:    "orange",
        Priority: 1,
    }
    pq.Push(item)

    // but we want change the priority...
    pq.Update(item, item.Value, 5)

    // Take the items out; they arrive in decreasing priority order.
    // fmt.Printf("%+v", pq)
    for !pq.IsEmpty() {
        item := pq.Pop()
        fmt.Printf("%.2d:%s ", item.Priority, item.Value)
    }
    fmt.Println()
}


