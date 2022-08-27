package main

import (
	"fmt"
)

func dfs(i,j int) bool {
	
}

func main() {
	var N,M int
	fmt.Scan(&N, &M)

	graph := make([][]int, N)
	for i:=0; i<N; i++ {
		graph[i] = make([]int, M)
		for j:=0; j < M; j++ {
			fmt.Scan(&graph[i][j])
		}
	}
}