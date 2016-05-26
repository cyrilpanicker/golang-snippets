package slices

import (
    "fmt"
)

type number int

func Run(){
    
    numbers := [6]int{0,1,2,3,4,5}
    
    var slice1 []int
    slice1 = numbers[1:4]
    
    slice2 := numbers[0:4]
    
    fmt.Println(numbers)
    fmt.Println(slice1)
    fmt.Println(slice2)
    
    slice1[2] = 8
    
    fmt.Println("------")
    
    fmt.Println(numbers)
    fmt.Println(slice1)
    fmt.Println(slice2)
    
}