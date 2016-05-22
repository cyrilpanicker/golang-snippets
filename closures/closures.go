package closures;

import (
    "fmt"
    "sync"
)

func Run(){
    
    var value1 int
    var value2 int
    
    for i:=1;i<=5;i++ {
        value1 = i
        value2 = i
        func(value2 int){
            fmt.Printf("%v, %v\n",value1,value2)
        }(value2)
    }
    
    fmt.Println("----");
    
    var waitGroup sync.WaitGroup
    waitGroup.Add(5)
    
    for i:=1;i<=5;i++ {
        value1 = i
        value2 = i
        go func(value2 int){
            fmt.Printf("%v, %v\n",value1,value2)
            waitGroup.Done()
        }(value2)
    }
    
    waitGroup.Wait()
}