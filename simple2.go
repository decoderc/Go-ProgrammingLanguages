package main

import "fmt"

func main() {
    slice1 := []int{1,3,5,7,9}
    slice2 := make([]int,3)
    copy(slice2,slice1)
    fmt.Println("slice1:",slice1, "\nslice2:",slice2)
    
}
