
package passReference

import (
    "fmt"
)

func Run(){

    var mainNumber = 1
    
    func(func1Number int){
        func1Number = 2
    }(mainNumber)
    fmt.Println(mainNumber)
    
    func(func2NumberPointer *int){
        *func2NumberPointer = 2
    }(&mainNumber)
    fmt.Println(mainNumber)

}