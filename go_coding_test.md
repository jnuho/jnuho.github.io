- 이것이 코딩테스트다 -> Golang 구현

# 01 코딩테스트 개요

## 복잡도 Big-O
- 시간 복잡도 : 알고리즘을 위해 필요한 연산의 횟수
- 공간 복잡도 : 알고리즘을 위해 필요한 메모리의 양

Trade-off Example: memoization 메모리를 많이 사용하여 시간을 비약적으로 줄이는 방법 Ch8


```go
package main
import "fmt"
func main() {

	array := []int{3, 5, 1, 2, 4}
	summary := 0
  // O(N)
  for _,x := range array {
    summary += x
	}
  
  // O(N^2)
	N := 0
  for i:= 0; i<len(array); i++ {
    for j:= 0; j<len(array); j++ {
			N++
		}
	}
	fmt.Println(N)
}
```

$O(N)$ 최악의 경우로 복잡도 고려

빅오 표기법 | 명칭
---|---
O(1) | 상수시간 (Constant time)
O(logN) | 로그시간 (Log Time)
O(N) | 선형 시간
O(NlogN) | 로그 선형 시간
O(N^2) |  이차 시간
O(N^3) | 삼차 시간
O(2^n) | 지수 시간

Big-O는 최악의 경우 시간복잡도 e.g.Quick Sort O(NlogN)~O(N^2)

## 시간과 메모리 측정


```go
import time

# 측정 시작
start_time = time.time()

# 프로그램 소스코드

# 측정 종료
end_time = time.time()

elapsed_time = end_time - start_time
print(f"Time elapsed : {elapsed_time}")
```

- 선택정렬의 경우 최악 $O(N^{2})$
- 파이썬 기본정렬 라이브러리 최악의 경우 $O(N\log{}N)$


```go
package main

import (
	"fmt"
	"time"
	"math/rand"
	"sort"
)

func main() {

	// 0~9정수 10개로 이루어진 int 슬라이스
	rand.Seed(time.Now().UnixNano())
	array := []int{}
	n := 10
	for i:=0; i< 10000; i++ {
		// rand.Intn(n) returns non-negative pseudo-random int [0,n)
		// It panics if n <= 0.
		array = append(array, rand.Intn(n))
	}
	// fmt.Println(array)

	// 1. 선택정렬
	var min_index int
	start_time := time.Now()
	for i:=0; i<len(array); i++ {
		min_index = i
		for j:=i+1; j<len(array); j++ {
			if array[j] < array[min_index] {
				min_index = j
			}
		}
		array[i], array[min_index] = array[min_index], array[i]
	}
	end_time := time.Now()
	elapsed_time := end_time.Sub(start_time)
	fmt.Println("Elapsed Time: ", elapsed_time)
	// fmt.Println(array)

	// Shuffle!
	rand.Shuffle(len(array), func(i, j int){
		array[i], array[j] = array[j], array[i]
	})
	// fmt.Println(array)


	// 2. 기본라이브러리 Sort정렬
	start_time = time.Now()
	sort.Ints(array)
	end_time = time.Now()
	elapsed_time = end_time.Sub(start_time)
	fmt.Println("Elapsed Time: ", elapsed_time)
	// fmt.Println(array)
}
```

# 02 코딩 테스트 유형 분석

- 알고리즘 기반, Greedy, Implmentation, DFS/BFS 탐색
- Dynamic Programming, Graph Theory (Basic Level)
- 정수론, 최단경로, Dynamic Programming (Contest Competitive Coding Test)

- 출제 비율
  - 구현 > DFS/BFS > 그리디 > 정렬 > Dynamic > 최단경로 > 이진탐색 > 그래프이론

- 카카오
  - "ACM-ICPC** 같은 어려운 알고리즘 설계 능력을 겨루는 문제가 아닌, 업무에서 있을 만한 상황을 가정하여 독창적이고 다양한 분야의 문제..."
  - 그리디, 구현, 문자열, 다양한 케이스고려
- 삼성
  - 예외상황 적절히 처리, 완전탐색, DFS/BFS, 구현유형
- 보통 2~5시간 내에 8개 이하의 문제


Upper Bound : 문제 해결 역량 & 코드포스 블루이상, ACM-ICPC 서울지역 대회 본선 수준

기업 | 날짜 | 풀이 시간 | 문제 개수 | 커트라인 | 주요 문제 유형 | 시험 유형
---|---|---| ---| ---|---|---
라인| 상반기<br>(2020-04-05)| 2시간 30분| 6문제| 4문제|구현, 문자열, 자료구조| 온라인
삼성전자| 상반기<br>(2020-06-07)| 3시간| 2문제| 2문제|완전탐색, 시뮬레이션, DFS/BFS| 오프라인


- 알고리즘 문제풀이 사이트
  - 코드시그널 https://app.codesignal.com
  - 코드포스 https://codeforces.com
  - 정올 http://www.jungol.com
  - 생활코딩 https://opentutorials.org
  - BOJ Slack https://acmicpc.slack.com


# 03 Greedy 알고리즘
- 현재 상황에서 가장 좋아 보이는 것만을 선택하는 알고리즘
- 매 순간 가장 좋아 보이는 것을 선택하며, 현재 선택이 나중에 미칠 영향은 고려 X
- 그리디 알고리즘 정당성 검토 필수 e.g. 큰단위 동전에 작은 단위 동전의 배수 일때만 가능 500, 400, 100는 적용 X
  - 배수가 아닌 무작위 동전 종류인 경우 다이나믹 프로그래밍으로 해결 가능: Ch8

```go
// 1.
// 당신은 음식점의 계산을 도와주는 직원
// 카운터에 거스름돈으로 사용할 500,100,50,10원 동전 무한개존재
// 손님에게 거슬러줘야 할 돈 N원 일때 거슬러 줘야할 최소 동전 개수는?
// N은 항상 10의 배수이다.
package main

import (
	"fmt"
)
func main() {
	// N=1260 일떄 result = 6
	var N int
	fmt.Scanln(&N)

	result := 0

	coins := []int{500, 100, 50, 10}
	for _,coin := range coins {
		result += N / coin
		N = N % coin
	}

	fmt.Println(result)
}
```


```go
// 2.
// 큰수의 법칙

// 배열크기 N
// 더해지는 횟수 M
// 반복가능 횟수 K
package main

import (
	"fmt"
)

func main() {
	// 5 8 3
	// 2 4 5 4 6 -> 6 5 4 4 2
	// 6+6+6+5+6+6+6+5
	var N,M,K int
	fmt.Scan(&N, &M, &K)

	in := make([]int, N)
	for i := range in {
		fmt.Scan(&in[i])
	}

	var first,second int
	first = in[0]
	for i:=0; i<N; i++ {
		if first < in[i] {
			second = first
			first = in[i]
		} else if second <in[i] {
			second = in[i]
		}
	}
	// a: K+1 이반복되는 횟수
	a := M / (K+1)
	// b: 마지막 사이클에서 반복적으로 더할 수 있는 횟수
	b := M % (K+1)

	result := a * (K*first + second) + b* first
	fmt.Println(result)
}
```


```go
// 3.
// 숫자카드게임
// N x M (행 x 열) 에서
// 먼저 행을 선택하고, 그 행에서 가장 작은 수를 뽑아서
// 가장 큰 수가 나오도록

// 3 3
// 3 1 2
// 4 1 4
// 2 2 2
// 2
package main

import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	var result int
	var temp int

	arr := make([][]int, N)
	for i:=0; i<N; i++ {
		temp_arr := make([]int, M)
		for j:=0; j< M; j++ {
			fmt.Scan(&temp_arr[j])
			if j==0 {
				temp = temp_arr[j]
			}
			if temp_arr[j] < temp {
				temp = temp_arr[j]
			}
		}
		arr[i] = temp_arr
		if temp > result {
			result = temp
		}
	}

	fmt.Println(result)
}
```


```go
// 4.
// 1이 될때까지
// 1. N에서 1을 뺀다
// 2. N을 K로 나눈다
// 1이 되는데 걸린 횟수 출력
package main

import (
	"fmt"
)

func main() {
	var N,K int
	fmt.Scan(&N, &K)

	count := 0
	for N != 1 {
		if N % K == 0 {
			N /= K
			count++
		} else {
			count += N - (N/K) *K
			N = N/K * K
		}
	}

	fmt.Println(count)
}
```

# 04 구현


```go
// top left (1,1) bottom right (N,N)
// 계획서 띄어쓰기 기준 L R U D 문자들이 반복적으로 적혀있음
// 움직일수 없는 곳 이동명령은 무시 됨
// 5
// R R R U D D
// 3 4
package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {
	r,c := 1,1

	var N int
	fmt.Scanln(&N)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	for _, d := range strings.Fields(line) {
		switch d {
		case "L":
			if c > 1 {
				c -=1
			}
		case "R":
			if c < N {
				c +=1
			}
		case "U":
			if r > 1 {
				r -=1
			}
		case "D":
			if r < N {
				r +=1
			}
		default:
			break
		}
	}

	fmt.Println(r, c)
}
```


```go
// 00시 00분 00초 ~ N시 59분 59초
// 3이 하나라도 포함되는 경우의 수
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var N int
	fmt.Scan(&N)

	count := 0

	// TODO
	for h:=0; h<=N; h++ {
		for m:=0; m<60; m++ {
			for s:=0; s<60; s++ {
				first := strings.Contains(strconv.Itoa(h), "3")
				second := strings.Contains(strconv.Itoa(m), "3")
				third := strings.Contains(strconv.Itoa(s), "3")
				if first || second || third {
					count++
				}
			}
		}
	}


	fmt.Println(count)
}
```


```go
// 체스판에서 (a~h, 1~8)
// 나이트 이동(2+1방식 이동)할 수 있는 개수
// a1
// 2
package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	c := scanner.Text()

	// x: 1~8, y: 1~8
	x := int(c[0]-96)
	y := int(c[1]-48)

	var count int

	steps := [][]int{
		[]int{1,2}, []int{-1,2}, []int{1,-2}, []int{-1,-2},
		[]int{2,1}, []int{-2,1}, []int{2,-1}, []int{-2,-1},
	}
	for _,move := range steps {
		if x+move[0] >=1 && x+move[0] <= 8 && y+move[1] >=1 && y+move[1] <=8 {
			count++
		}
	}
	fmt.Println(count)
}
```

```go
// 게임개발
// 1. 현재방향 왼쪽방향 부터 차례로 갈곳 정함
// 2. 왼쪽 회전 후 안 가본칸이면 전진, 방문한 칸이라면 왼쪽 회전만
// 3. 네칸 모두 가본칸이거나 바다라면, 방향 유지한채 한칸 뒤로 (뒤가 바다이면 움직임 멈춤)
// 0 북 1 동 2 남 3 서
// 0 육지 1 바다
// N x M 맵
// 캐릭터가 방문한 칸의 수 출력

// 4 4
// 1 1 0  // (1,1)에 북쪽 0을 바라보고 서 있는 캐릭터
// 1 1 1 1
// 1 0 0 1
// 1 1 0 1
// 1 1 1 1
// 3
// TODO 뒤로 돌아가는 케이스에서 일부 오류 있음
// 모든 칸을 다 방문하지 않는 예외케이스 발생 (expected: 6, actual: 5)
// 5 5
// 1 1 0
// 1 1 1 1 1
// 1 0 0 0 1
// 1 1 0 0 1
// 1 1 0 1 1
// 1 1 1 1 1
// 5
package main

import (
	"fmt"
)

func main() {
	var N,M int
	var r,c, d int

	fmt.Scanln(&N, &M)
	fmt.Scanln(&r, &c, &d)

	area := make([][]int, N)
	for r:=0; r< N; r++ {
		temp := make([]int, M)
		for c:=0; c< M; c++ {
			fmt.Scan(&temp[c])
		}
		area[r] = temp
	}

	dr := []int{-1,0,1,0}
	dc := []int{0,1,0,-1}

	count := 1
	for {
		allFour := false
		for i:=0; i<4; i++ {
			area[r][c] = 2
			dd := ((d-1) +4) %4
			d = dd
			isRange := r+ dr[dd] >=0 && r+dr[dd] <N && c+ dc[dd] >=0 && c+dc[dd] <M 
			if isRange && area[r + dr[dd]][c + dc[dd]] == 0 {
				r = r + dr[dd]
				c = c + dc[dd]
				count++
				// fmt.Println("움직인후", r,c,count, d)
				break
			}
			if i==3 {
				allFour = true
			}
		}
		if allFour {
			newr := r + dr[(d+2) % 4]
			newc := c + dc[(d+2) % 4]
			if area[newr][newc] != 1 {
				r = newr
				c = newc
			} else {
				break
			}
		}
	}
	fmt.Println(count)
}
```

# 05 DFS/BFS

### 자료구조 기초
- Search 탐색
- Data Structure
    - Stack
    - Queue


```go
// 스택구현
package main

import (
	"fmt"
	"sort"
)

func main() {
	stack := []int{}
	stack = append(stack, 5)
	stack = append(stack, 2)
	stack = append(stack, 3)
	stack = append(stack, 7)
	stack = stack[:len(stack)-1]
	stack = append(stack, 1)
	stack = append(stack, 4)
	stack = stack[:len(stack)-1]

	// 5 2 3 1
	fmt.Println(stack)

	sort.Sort(sort.IntSlice(stack))
	fmt.Println(stack)

	sort.Sort(sort.Reverse(sort.IntSlice(stack)))
	fmt.Println(stack)
}
```

```go
// 큐구현
package main

import (
	"fmt"
	"sort"
)

func main() {
	queue := []int{}
	queue = append(queue, 5)
	queue = append(queue, 2)
	queue = append(queue, 3)
	queue = append(queue, 7)
	queue = queue[1:]
	queue = append(queue, 1)
	queue = append(queue, 4)
	queue = queue[1:]

	// 3 7 1 4
	fmt.Println(queue)

	sort.Sort(sort.IntSlice(queue))
	fmt.Println(queue)

	sort.Sort(sort.Reverse(sort.IntSlice(queue)))
	fmt.Println(queue)
}
```


### 재귀 함수


```go
package main

import (
	"fmt"
)

func factorial(n int) int {
	if n <= 1{
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println(factorial(5))
}
```



### DFS

- Depth-First Search
	- 깊이 우선 탐색: 그래프에서 깊이있는 부분을 우선적으로 탐색
	- Node(또는 Vertex 정점), Edge(간선)
	- Adjacent(인접) 노드
	- 탐색 시작노드 스택삽입 & 방문처리
	- -> 스택 최상단 노드에 미방문 인접노드(일반적으로 인접노드들 중 가장 작은 숫자 먼저처) 있으면 그 인접노드를 스택에 넣고 방문처리
	- -> 방문하지 않은 인접노드가 없으면 스택에서 최상단 노드를 꺼냄

- 그래프 Graph 표현 방식
	- 인접 행렬 Adjacency Matrix: 2차원 배열 그래프 각 노드가 연결된 형태 기록
	- 인접 리스트 Adjacency List: 리스트로 그래프의 연결관계를 표현하는 방식

-| 메모리| 속도
---|---|---
인접 행렬 | 노드간 모든관계를 저장하므로, 노드많을 수록 불필요하게 메모리 증가 |
인접 리스트 | 연결된 정보만 저장하므로 효율적 메모리 사용 | 두 노드 연결여부 확인 속도는 느림

- 예를들어 노드1, 노드7 연결여부 확인시
	- 행렬의 경우 graph[1][7]만 확인 하면 됨
	- 리스트의 경우 노드1에 대한 인접 리스트를 앞에서 부터 차례로 확인 -> 특정 노드와 연결된 모든 노드를 순회할때는 메모리 공간 낭비 적음


```go
// 1-(7)-0-(5)-2
package main

import (
	"fmt"
)

type pair struct {
	node int
	distance int
}

func main() {
	// 인접 행렬
	//   0  1  2
	// 0 0  7  5
	// 1 7  0  INF
	// 2 5 INF 0
	INF := 999999999
	graph_matrix := [][]int {
		{0, 7, 5},
		{7, 0, INF},
		{5, INF, 0},
	}
	fmt.Println(graph_matrix)

	// 인접 리스트
	// 0 -> 1(7) -> 2(5)
	// 1 -> 0(7)
	// 2 -> 0(5)
	graph_list := make([][]pair, 3)
	graph_list[0] = make([]pair, 0)
	graph_list[0] = append(graph_list[0], pair{1,7})
	graph_list[0] = append(graph_list[0], pair{2,5})

	graph_list[1] = make([]pair, 0)
	graph_list[1] = append(graph_list[1], pair{0,7})

	graph_list[2] = make([]pair, 0)
	graph_list[2] = append(graph_list[2], pair{0,5})

	fmt.Println(graph_list)
}
```

- DFS 스택 자료구조 사용 시 :
	- 탐색 시작노드를 스택에 삽입하고 방문처리
	- 스택 최상단 노드에 방문하지 않은 인접노드(여러개면 번호가 관행적으로 낮은 순서부터 처리) 있으면, 스택에 넣고 방문처리. 방문하지 않은 인접노드가 없으면, 스택에서 최상단 노드를 꺼낸다
	(방문처리는 스택에 한번 삽입되어 처리된 노드가 다시 삽입되지 않게 체크)

![graph](./assets/images/graph.png)


- DFS를 재귀함수로 구현하며 느려질수 있으므로 Stack활용도 가능
	- DFS보다 BFS가 좀더 빠르게 동작

```go
// DFS 메소드 (인접리스트- 2차원리스트 활용)
// 	인접행렬 방식으로 노드 1,2,..., 8 표현
// 	graph[0]은 없으므로 empty리스트 []
// 	graph[1] ~ graph[8] 표현
package main

import (
	"fmt"
)

func dfs(graph [][]int, v int, visited []bool) {
	// 현재노드 방문처리
	visited[v] = true
	fmt.Print(v, " ")

	// 현재노드와 연결된 다른 노드를 Recursive 방문
	for _,i := range graph[v] {
		if !visited[i] {
			dfs(graph, i, visited)
		}
	}
}

func main() {
	graph := [][]int {
		{},
		{2,3,8},
		{1,7},
		{1,4,5},
		{3,5},
		{3,4},
		{7},
		{2,6,8},
		{1,7},
	}
	visited := []bool {false,false,false,false,false,false,false,false,false}

  // 1 2 7 6 8 3 4 5 
	dfs(graph, 1, visited)
	fmt.Println()
}
```


### BFS

- Breath-First Search
    - 너비 우선 탐색. 가까운 노드부터 탐색
    - 탐색 시작노드 큐 삽입 & 방문처리
    - 큐에서 노드를 꺼내 미방문 인접노드 중 방문하지 않은 노드를 모두 큐에 삽입 및 방문처리


```go
// 인접행렬 방식으로 노드 1,2,..., 8 표현
// graph[0]은 없으므로 empty리스트 []
// graph[1] ~ graph[8] 표현
package main

import (
	"fmt"
)

func bfs(graph [][]int, start int, visited []bool) {

	// 큐 구현을 위한
	queue := []int{start}

	// 현재노드 방문처리
	visited[start] = true

	for len(queue) > 0 {
		// 큐에서 원소 하나 Pop하여 출력
		v := queue[0]
		queue = queue[1:]
		fmt.Print(v, " ")

		for _, i := range graph[v] {
			if !visited[i] {
				queue = append(queue, i)
				visited[i] = true
			}
		}
	}
}

func main() {
	graph := [][]int {
		{},
		{2,3,8},
		{1,7},
		{1,4,5},
		{3,5},
		{3,4},
		{7},
		{2,6,8},
		{1,7},
	}
	visited := []bool {false,false,false,false,false,false,false,false,false}

  // 1 2 3 8 7 4 5 6 
	bfs(graph, 1, visited)
	fmt.Println()
}
```


-|DFS|BFS
---|---|---
동작원리 | 스택|큐
구현방법|재귀함수 이용 | 큐 자료구조


- 음료수 얼려 먹기

```go

```

```go

# 선택 정렬
def selection_sort(arr):
  for i in range(len(arr)):
    min_idx = i
    for j in range(i+1, len(arr)):
      if arr[j] < arr[min_idx]:
        arr[j], arr[min_idx] = arr[min_idx], arr[j]
  print(arr)

# 삽입 정렬
def insertion_sort(arr):
  for i in range(len(arr)):
    for j in range(i,0,-1):
      if arr[j] < arr[j-1]:
        arr[j], arr[j-1] = arr[j-1], arr[j]
      else:
        break 
  print(arr)

# 퀵 정렬
# def quick_sort(arr, start, end):
#   if start >= end:
#     return
#   pivot = start
#   left = start+1
#   right = end

#   while left <= right:
#     while left <= end and arr[left] <= arr[pivot]:
#       left += 1
#     while right > start and arr[right] >= arr[pivot]:
#       right -=1
#     if left > right:
#       arr[right], arr[pivot] = arr[pivot], arr[right]
#     else:
#       arr[right], arr[left] = arr[left], arr[right]

#   quick_sort(arr, start, right-1)
#   quick_sort(arr, right+1, end)
def quick_sort(arr):
  if len(arr) <= 1:
    return arr
  
  pivot = arr[0]
  tail = arr[1:]

  left_side = [x for x in tail if x <= pivot]
  right_side = [x for x in tail if x > pivot]

  return quick_sort(left_side) + [pivot] + quick_sort(right_side)

def counting_sort(arr):
  count_arr = [0] * (max(arr)+1)

  for i in arr:
    count_arr[i] += 1

  for i in range(len(count_arr)):
    for j in range(count_arr[i]):
      print(i, end=' ')




arr = [7, 5, 9, 0, 3, 1, 6, 2, 4, 8]
selection_sort(arr)

arr = [7, 5, 9, 0, 3, 1, 6, 2, 4, 8]
insertion_sort(arr)

arr = [5, 7, 9, 0, 3, 1, 6, 2, 4, 8]
arr = quick_sort(arr)
print(arr)

arr = [7, 5, 9, 0, 3, 1, 6, 2, 9, 1, 4, 8, 0, 5, 2]
counting_sort(arr)

```

    [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
    [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
    [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
    0 0 1 1 2 2 3 4 5 5 6 7 8 9 9 

- 선택 정렬 O(N^2)
- 삽입 정렬 O(N^2)
- 퀵 정렬 O(NlogN)
- 계수 정렬 O(N+K) : 모든 데이터가 양의 정수, 데이터수 N, 최대정수값 K
- 파이썬 정렬 라이브러리 O(NlogN)
  - sorted
  - sort
  - key매개변수를 이용한 정렬기준 설정 가능


```go
N = int(input())

arr = []
for i in range(N):
  arr.append(int(input()))

arr.sort(reverse=True)

for i in arr:
  print(i, end=' ')

```

    27 15 12 


```go

N = int(input())

arr = []
for i in range(N):
  arr.append(input().split())

arr = sorted(arr,key=lambda x: x[1])

for i in arr:
  print(i[0], end=' ')
```

    이순신 홍길동 


```go

N, K = map(int, input().split())

A = list(map(int, input().split()))
B = list(map(int, input().split()))

A.sort()
B.sort(reverse=True)
for i in range(K):
  if A[i] <B[i]:
    A[i], B[i] = A[i], B[i]
  else:
    break

print(sum(A))
```



```go
def binary_search(array, target, start, end):
  if start > end:
      return None
  mid = (start+end)//2
  if array[mid] == target:
    return mid
  elif array[mid] > target:
    return binary_search(array,target, start, mid-1)
  else:
    return binary_search(array, target, mid+1,end)

N, target = map(int, input().split())
array = list(map(int, input().split()))

result = binary_search(array, target, 0, N-1)
if result == None:
  print("원소가 존재하지 않습니다.")
else:
  print(result+1)

```

