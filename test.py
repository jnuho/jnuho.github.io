
N, M = map(int, input().split())
dduks = list(map(int, input().split()))
dduks.sort()

start = 0
end = max(dduks)

while start <= end:
  mid = (start+end) //2
  res = 0
  for i in dduks:
    res += max(0, i-mid)

  if res == M:
    print(mid)
    break
  elif res < M:
    end = mid-1
  else:
    start = mid+1

  
