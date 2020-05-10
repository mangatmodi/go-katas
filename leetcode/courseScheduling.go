package main

import "fmt"

//https://leetcode.com/problems/course-schedule/
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := makeGraph(prerequisites)
	done := map[int]bool{}
	for u := range graph {
		if len(done) == numCourses {
			//got all nodes
			return true
		}
		if done[u] {
			continue
		}
		//start new dfs from u as root
		if hasCycle(graph, u, map[int]bool{}, done) {
			return false
		}
	}
	return true
}

func hasCycle(graph map[int][]int, u int, current, done map[int]bool) bool {
	current[u] = true
	a := graph[u]
	for _, v := range a {
		if current[v] {
			//cycle
			return true
		}
		if hasCycle(graph, v, current, done) {
			return true
		}
	}
	done[u] = true
	delete(current, u)
	return false
}

func makeGraph(prerequisites [][]int) map[int][]int {
	ans := make(map[int][]int, len(prerequisites)/2)
	for _, a := range prerequisites {
		u := a[0]
		v := a[1]
		neighbours, ok := ans[u]
		if !ok {
			neighbours = []int{}
		}
		neighbours = append(neighbours, v)
		ans[u] = neighbours
	}
	return ans
}

func main() {
	fmt.Println(canFinish(3, [][]int{{1, 0},{0,2},{2,1}}))
}
