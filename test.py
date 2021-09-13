



def dfs(graph, v, visited):
  visited[v] = True
  print(v, end=' ')
  for i in graph[v]:
    if not visited[i]:
      dfs(graph, i, visited)

graph = [
  []
  , [2,5,9]
  , [1,3]
  , [2,4]
  , [3]
  , [1,6,8]
  , [5,7]
  , [6]
  , [5,6]
  , [1,10]
  , [9]
]
visited = [False]*11

dfs(graph, 1, visited)