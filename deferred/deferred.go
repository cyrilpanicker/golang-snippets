package deferred

import "fmt"

func Run(){
    i:=1
    defer fmt.Printf("deferred1 : %d\n",i)
    i++
    defer fmt.Printf("deferred2 : %d\n",i)
    i++
    fmt.Println(i)
}