package topologicalsort

func dfs(arr [][]int, flags []int, cur int) bool {
	if flags[cur] == 1 {
		return false
	}
	if flags[cur] == -1 {
		return true
	}
	flags[cur] = 1
	for j := 0; j < len(arr[cur]); j++ {
		if (arr[cur][j] == 1 && !dfs(arr, flags, j)) {
			return false
		}
	}
	flags[cur] = -1
	return true
}

func TopologicalSortDfs(numCourses int, prerequisites [][]int) bool {
	flags := make([]int, numCourses)
	arr := make([][]int, numCourses)
	for i := 0; i < len(arr); i++ {
		arr[i] = make([]int, numCourses)
	}
	for i := 0; i < len(prerequisites); i++ {
		arr[prerequisites[i][1]][prerequisites[i][0]] = 1
	}
	for i := 0; i < numCourses; i++ {
		if (!dfs(arr, flags, i)) {
			return false
		}
	}
	return true
}
