/*
	chenshuo test
*/

package heap4l

import (
	"container/heap"
	"fmt"
)

type nodelist []int

func (n nodelist) Less(i, j int) bool {
	return n[i] > n[j]
}

func (n nodelist) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n nodelist) Len() int {
	return len(n)
}

func (n *nodelist) Push(x interface{}) {
	*n = append(*n, x.(int))
}

func (n *nodelist) Pop() interface{} {
	l := len(*n) - 1
	x := (*n)[l]
	*n = (*n)[0:l]
	return x
}

type node struct {
	key   string
	value int
	index int
}

type nodeHeap []*node

func (n nodeHeap) Less(i, j int) bool {
	return n[i].value < n[j].value
}

func (n nodeHeap) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
	n[i].index = i
	n[j].index = j
}

func (n nodeHeap) Len() int {
	return len(n)
}

func (n *nodeHeap) Push(x interface{}) {
	l := len(*n)
	item := x.(*node)
	item.index = l
	*n = append(*n, item)
}

func (n *nodeHeap) Pop() interface{} {
	l := len(*n) - 1
	x := (*n)[l]
	x.index = -1
	*n = (*n)[0:l]
	return x
}

func (n *nodeHeap) Update(item *node, key string, value int) {
	item.value = value
	item.key = key
	heap.Fix(n, item.index)
}

func (n *nodeHeap) Delete(item *node) {
	heap.Remove(n, item.index)
}

func HeapFunc() {
	var heapInfo nodelist
	heap.Init(&heapInfo)
	heap.Push(&heapInfo, 94)
	heap.Push(&heapInfo, 91)
	heap.Push(&heapInfo, 92)
	heap.Push(&heapInfo, 88)
	heap.Push(&heapInfo, 44)
	heap.Push(&heapInfo, 199)
	for len(heapInfo) > 0 {
		fmt.Println(heap.Pop(&heapInfo).(int))
	}

	var heap1 nodeHeap
	heap.Init(&heap1)
	heap.Push(&heap1, &node{key: "yao5", value: 94})
	heap.Push(&heap1, &node{key: "yao3", value: 91})
	heap.Push(&heap1, &node{key: "yao4", value: 92})
	heap.Push(&heap1, &node{key: "yao2", value: 88})
	heap.Push(&heap1, &node{key: "yao6", value: 198})
	x := &node{key: "yao1", value: 44}
	heap.Push(&heap1, x)
	heap1.Update(x, "yao1", 89)
	for len(heap1) > 0 {
		fmt.Println("after ", heap.Pop(&heap1).(*node))
	}
}
