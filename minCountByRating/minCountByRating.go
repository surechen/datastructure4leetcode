package minCountByRating

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func candy(ratings []int) int {
	n := len(ratings)
	arr := make([]int, len(ratings))
	if n <= 1 {
		return 1
	}

	// left == big than left
	left := make([]int, n)
	// right == big than right
	right := make([]int, n)
	left[0] = 1
	right[n-1] = 1

	for i := 1; i < n; i++ {
		if ratings[i-1] < ratings[i] {
			// i 比左边的大
			// 所以，他的数量要比左边的多一个
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}

		if ratings[n-i-1] > ratings[n-i] {
			// n-i-1 比右边的大
			// 所以，他的数量要比右边的多一个
			right[n-i-1] = right[n-i] + 1
		} else {
			right[n-i-1] = 1
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		// i 的实际数量是max(left[i], right[i]
		arr[i] = max(left[i], right[i])
		res += arr[i]
	}
	return res
}

func MinCountByRating() {
	fmt.Println(candy([]int{1, 3, 2, 1, 3}))
}