/*
	chenshuo test
*/

package list4l

import (
	"container/list"
	"fmt"
)

type node struct {
	key string
	val int
}

func ListFunc() {
	l := list.New()
	for i := 0; i < 5; i++ {
		l.PushBack(i)
	}
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
	fmt.Println("")
	fmt.Println(l.Front().Value)
	fmt.Println(l.Back().Value)
	l.InsertAfter(6, l.Front())
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
	fmt.Println("")
	l.MoveBefore(l.Front().Next(), l.Front())
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
	fmt.Println("")
	l.MoveToFront(l.Back()) //将尾部元素移动到首部
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
	fmt.Println("")
	l2 := list.New()
	l2.PushBack(1)
	l2.PushBack(node{key: "ss9", val: 2})
	l2.PushBackList(l)
	l2.PushBack(node{key: "ss", val: 1})
	for e := l2.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
	fmt.Println("")
	l.PushBack(node{key: "ss", val: 1})
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
	fmt.Println("")
	l.Remove(l.Front())
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
	fmt.Println("")
}
