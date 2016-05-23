package initDemo

import (
    "github.com/cyrilpanicker/golang-snippets/testPackages/package1"
    _ "github.com/cyrilpanicker/golang-snippets/testPackages/package2"
)

func Run(){
    package1.Package1Function()
}
