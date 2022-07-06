package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

func main() {

    var num int64

    var wg sync.WaitGroup

    for i := 0; i < 10; i++ {
        wg.Add(1)

        go func() {
            for c := 0; c < 100; c++ {

                atomic.AddInt64(&num, 1)
            }
            wg.Done()
        }()
    }

    wg.Wait()

    fmt.Println("No:", num)
}