from collections import deque

# BFS 메소드
def bfs(graph, start, visited):
  queue = deque([start])


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

bfs(graph,1,visited)