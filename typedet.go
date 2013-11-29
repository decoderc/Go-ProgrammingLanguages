package main 

import "fmt"

type Matris struct{
    a, b, c, d float64
}

func (x *Matris) det() float64{
    y := x.a * x.d - x.b * x.c
    return y
}

func main(){
    x := Matris{1, 2, 3, 4}
    fmt.Println("kare matris determinantı:", x.det())
}
