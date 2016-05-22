package initDemo

import (
    "cyril/golang-snippets/testPackages/package1"
    _ "cyril/golang-snippets/testPackages/package2"
)

func Run(){
    package1.Package1Function()
}