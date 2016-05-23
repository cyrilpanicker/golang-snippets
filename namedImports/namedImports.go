package namedImports

import (
    "fmt"
    custom_fmt "github.com/cyrilpanicker/golang-snippets/testPackages/fmt"
)

func Run(){
    fmt.Println("from fmt standard library")
    custom_fmt.Print()
}
