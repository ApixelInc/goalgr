package main

import (
    "github.com/ApixelInc/goalgr/LLIST"
    "fmt"
)

func linkedlist_ex() {
    queue := LLIST.New()

    for i:=0;  i< 5; i++ {
        queue.PushItem( fmt.Sprintf( "item:%d",i )  )
    }

    for !queue.IsEmpty() {
        s := queue.Pop()
        fmt.Printf("%+v ", s )
    }
}
