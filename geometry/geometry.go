package main

import (
    "fmt"
    "geometry/rectangle"
    "log"
    )
// 1. package variables

var rectLen, rectWidth float64 = 0.2,-7.0

// 2. init function to check if if length and width are greater than zero

func init(){
    fmt.Println("main package initialized")
    if rectLen < 0 {
        log.Fatal("length is less than zero")
    }
    if rectWidth < 0 {
        log.Fatal("width is less than zero")
    }
}



func main(){

    fmt.Println("Geometrical shape properties")
    fmt.Printf("area of rectangle %.2f\n",rectangle.Area(rectLen, rectWidth))
    fmt.Printf("diagonal of the rectangle %.2f\n",rectangle.Diagonal(rectLen, rectWidth))
}
