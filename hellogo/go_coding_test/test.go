// NxM 미로
// (1,1) -> NxM 최단경로로 괴물을 피해서 갈때 움직인 칸의 수
// 0: 괴물 1:괴물없는 부분
package main

import (
	"fmt"
	"strconv"
)

type coord struct {
	x int
	y int
}

func bfs(graph [][]int) int {
	N := len(graph)
	M := len(graph[0])

	r,c := 0,0

	dr := []int{1, 0, -1, 0}
	dc := []int{0, 1, 0, -1}

	queue := make([]coord, 0)
	queue = append(queue, coord{k, j})

	var nx,ny int

	for queue {
		for i:=0; i<4; i++ {
			nr = r + dr[i]
			nc = c + dc[i]
			if nr < 0 || nr >= N || nc < 0 || nc >= M {
				continue
			}
			if graph[nr][nc] == 0 {
				continue
			}
			if graph[nr][nc] == 1 {
				graph[nr][nc] = graph[r][c] + 1

			}

		}
	}
}

func main() {
	var N,M int
	fmt.Scan(&N, &M)

	graph := make([][]int, N)
	var line string
	for i:=0; i<N; i++ {
		graph[i] = make([]int, M)
		fmt.Scanln(&line)
		for j:=0; j < M; j++ {
			n,_ := strconv.Atoi(string(line[j]))
			graph[i][j] = n
		}
	}

	result := 0
	// for i:=0; i<N; i++ {
	// 	for j:=0; j < M; j++ {
	// 		}
	// 	}
	// }

	fmt.Println(result)
}