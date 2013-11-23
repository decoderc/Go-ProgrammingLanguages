package main

import "fmt"

func main(){
    add := func(x,y int) int{
        return x + y
    }
    
    sub := func(x,y int) int{
        return x-y
    }
    
    mul := func(x,y int) int{
        return x*y
    }
    
    div := func(x,y float64) float64{
        return x/y
    }
    
        
    
    fmt.Println(add(2,3))
    fmt.Println(sub(96,45))
    fmt.Println(mul(43,46))
    fmt.Println(div(67,32))
}




    
