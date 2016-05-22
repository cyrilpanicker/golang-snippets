package waitGroup

import (
	"fmt"
	"sync"
)

//WaitGroup : demo on WaitGroup
func Run() {
	fmt.Println("before func")
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	go func ()  {
		fmt.Println("func1")
		waitGroup.Done();
	}()
	go func ()  {
		fmt.Println("func2")
		waitGroup.Done();
	}()
	fmt.Println("after func")
	waitGroup.Wait()
}