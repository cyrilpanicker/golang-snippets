package namedImports

import (
    "fmt"
    custom_fmt "cyril/golang-snippets/testPackages/fmt"
)

func Run(){
    fmt.Println("from fmt standard library")
    custom_fmt.Print()
}