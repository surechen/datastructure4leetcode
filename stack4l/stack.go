/*
copy from https://www.cnblogs.com/chinabrle/p/3541993.html
edit 2020/03/17
*/

package stack4l

import (
	"container/list"
	"fmt"
)

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	list := list.New()
	return &Stack{list}
}

func (stack *Stack) Push(value interface{}) {
	stack.list.PushBack(value)
}

func (stack *Stack) Pop() interface{} {
	e := stack.list.Back()
	if e != nil {
		stack.list.Remove(e)
		return e.Value
	}
	return nil
}

func (stack *Stack) Top() interface{} {
	e := stack.list.Back()
	if e != nil {
		return e.Value
	}

	return nil
}

func (stack *Stack) Len() int {
	return stack.list.Len()
}

func (stack *Stack) Empty() bool {
	return stack.list.Len() == 0
}

func (stack *Stack) Scan() []interface{} {
	var re []interface{}
	end := stack.list.Back()
	for end != nil {
		re = append(re, end.Value)
		end = end.Prev()
	}
	return re
}

type node struct {
	key string
	val int
}

func StackFunc() {
	stack := NewStack()
	stack.Push(&node{key: "yao", val: 96})
	stack.Push(&node{key: "james", val: 91})
	stack.Push(&node{key: "o'neal", val: 99})
	stack.Push(&node{key: "durant", val: 99})
	stack.Push(&node{key: "bryant", val: 99})
	sre := stack.Scan()
	for _, v := range sre {
		fmt.Print(v.(*node))
	}
	fmt.Println()
	fmt.Println(stack.Scan())
}
