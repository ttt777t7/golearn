package main

import (
    "fmt"
    "sync"
)

var x = 0

func increment(wg *sync.WaitGroup){
    fmt.Println("hello,world")
    wg.Done()
}

func main(){
    var w sync.WaitGroup
    for i := 0;i<5;i++{
        w.Add(1)
        go increment(&w)
    }
    w.Wait()
    fmt.Println("Final value of x",x)

}
