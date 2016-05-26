package arrays

import (
    "fmt"
)

func Run(){

    var numbers [3]int
    numbers[0] = 1
    numbers[1] = 2
    numbers[2] = 3
    fmt.Println(numbers)

    strings := [4]string{"a","b","c","d"}
    fmt.Printf("%s",strings[1])

}