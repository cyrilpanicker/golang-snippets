package structType

import (
    "fmt"
)

type point struct{
    x int
    y int
}

func Run(){
    
    p1 := point{x:1,y:2}
    p1pointer := &p1
    
    p1.x = 3
    p1pointer.y=4
    
    fmt.Println(p1)
    fmt.Println(p1pointer)
    
    p2 := point{x:6}
    fmt.Println(p2)
    
    p3 := point{}
    fmt.Println(p3)
    
}