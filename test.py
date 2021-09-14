from collections import deque

def bfs(graph, start, visited):
  # 큐Queue 구현을 위해 deque 라이브러리 사용사용
  queue = deque([start])
  # 현재 노드 방문 처리
  visited[start] = True
  # 큐가 빌 때까지 반복
  while queue:
    # 큐에서 하나의 원소를 뽑아 출력
    v = queue.popleft()
    print(v, end=' ')
    for i in graph[v]:
      if not visited[i]:
        queue.append(i)
        visited[i] = True


graph = [
  []
  , [2,3,8]
  , [1,7]
  , [1,4,5]
  , [3,5]
  , [3,4]
  , [7]
  , [2,6,8]
  , [1,7]
]
visited = [False]*9

bfs(graph, 1, visited)