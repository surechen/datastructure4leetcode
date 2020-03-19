/*
leetcode 547
*/
package unioncheckingset

import "fmt"

func getroot(arr []int, i int) int {
	for i != arr[i] {
		i = arr[i]
	}
	return i
}

func settoor(arr []int, i int, j int) {
	if j != arr[j] {
		j = getroot(arr, j)
	}
	i = getroot(arr, i)
	arr[j] = i
}

func findCircleNum(M [][]int) int {
	arr := make([]int, len(M))
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
	for i := 0; i < len(M); i++ {
		for j := 0; j < len(M[0]); j++ {
			if i >= j || M[i][j] == 0 {
				continue
			}
			settoor(arr, i, j)
		}
	}
	mapinfo := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		mapinfo[getroot(arr, i)] = 1
	}
	return len(mapinfo)
}

func UnionCheckingSet() {
	fmt.Println(findCircleNum([][]int{[]int{1,1,0}, []int{1,1,0}, []int{0,0,1}}))
	fmt.Println(findCircleNum([][]int{[]int{1,1,1}, []int{1,1,1}, []int{1,1,1}}))
	//[[1,0,0,1],[0,1,1,0],[0,1,1,1],[1,0,1,1]]
	fmt.Println(findCircleNum([][]int{[]int{1,0,0,1}, []int{0,1,1,0}, []int{0,1,1,1}, []int{1,0,1,1}}))
}
