/*
	chenshuo test
*/

package sort4l

import (
	"fmt"
	"sort"
)

type s struct {
	key   string
	medal int
}

type arrs []s

func (arr arrs) Len() int {
	return len(arr)
}

func (arr arrs) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (arr arrs) Less(i, j int) bool {
	if arr[i].medal < arr[j].medal {
		return true
	} else if arr[i].medal == arr[j].medal {
		return arr[i].key < arr[j].key
	}
	return false
}

func reverse(arr []int) {
	if len(arr) == 0 {
		return
	}
	for begin, end := 0, len(arr)-1; begin != end; begin, end = begin+1, end-1 {
		arr[begin], arr[end] = arr[end], arr[begin]
	}
	return
}

func slicedelete(arr []int, i int) []int {
	arr = append(arr[:i], arr[i+1:]...)
	return arr
}

func Sort4lFunc() {
	intList := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
	sort.Sort(sort.Reverse(sort.IntSlice(intList)))
	fmt.Println("intList", intList)

	float8List := []float64{4.2, 5.9, 12.3, 10.0, 50.4, 99.9, 31.4, 27.81828, 3.14}
	sort.Sort(sort.Reverse(sort.Float64Slice(float8List)))
	fmt.Println("float8List", float8List)

	stringList := []string{"a", "c", "b", "d", "f", "i", "z", "x", "w", "y"}
	sort.Sort(sort.Reverse(sort.StringSlice(stringList)))
	fmt.Println("stringList", stringList)

	a := []int{1, 2, 4}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println("a", a)

	b := []int{6, 7, 8}
	fmt.Println("b", append(a, b[0:3]...))

	ar := arrs{s{key: "zhongguo", medal: 100}, s{key: "meiguo", medal: 100}, s{key: "russia", medal: 77}, s{key: "germany", medal: 101}}
	sort.Sort(ar)
	fmt.Println("ar", ar)

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(arr)
	reverse(arr)
	fmt.Println(arr)
	arr = slicedelete(arr, 3)
	fmt.Println(arr)
	arr = arr[4:5:5]
	fmt.Println(arr, cap(arr), len(arr))
}
