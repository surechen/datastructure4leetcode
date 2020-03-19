package topologicalsort

func TopologicalSortBfs(numCourses int, prerequisites [][]int) bool {
	arr := make([][]int, numCourses)
	for i := 0; i < len(arr); i++ {
		arr[i] = make([]int, numCourses)
	}
	indgree := make([]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		arr[prerequisites[i][1]][prerequisites[i][0]] = 1
		indgree[prerequisites[i][0]]++
	}
	list := make([]int, 0)
	for i := 0; i < len(indgree); i++ {
		if indgree[i] == 0 {
			list = append(list, i)
		}
	}
	count := 0
	for len(list) > 0 {
		v := list[0]
		list = list[1:]
		count++
		for j := 0; j < len(arr[0]); j++ {
			if arr[v][j] == 1 {
				indgree[j]--
				if indgree[j] == 0 {
					list = append(list, j)
				}
			}
		}
	}

	return count == numCourses
}



