# goalgr

## PriorityQueue

```go
imporrt "github.com/ApixelInc/goalgr/PQ"

    // push
    item := &PQ.Item{
        Value:    "orange",
        Priority: 1,
    }
    pq.Push(item)

    // pop
    if !pq.IsEmpty() {
        item := pq.Pop()
    }

```
