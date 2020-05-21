package sync4l

import (
	"fmt"
	"sync"
)

func MutexFuncErr() {
	var m sync.Mutex
	m.Lock()
	fmt.Println("a")
	fmt.Println("b")
	defer m.Unlock()
	m.Lock()
	fmt.Println("c")
	fmt.Println("d")
	m.Unlock()
}