# goalgr

## PriorityQueue

```go
imporrt "github.com/ApixelInc/goalgr/PQ"

    // create and init with existing items
    pq := PQ.New()

    items := []string{
        "banana", "apple", "pear",
    }
    for i, value := range items {
        priority := i
        pq.Push( value, priority )
    }

    item := pq.Push( "orange", 1 )

    // but we want change the priority...
    pq.Update(item.(*PQ.InternalItem), "orange2"  , 0)

    // Take the items out; they arrive in decreasing priority order.
    for !pq.IsEmpty() {
        fruit := pq.Pop()
        fmt.Printf("%+v ", fruit )
    }
```
