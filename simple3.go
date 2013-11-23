package main

import "fmt"

func main() {
    var x [4]float64
    x[0] = 23
    x[1] = 34
    x[2] = 45
    x[3] = 56
    
    var avarage float64 = 0
    for _, value := range x {
    avarage += value
    }
    fmt.Println(avarage / float64(len(x)))
    
}
