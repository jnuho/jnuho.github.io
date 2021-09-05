
N,M,K = map(int,input().split())
arr= list(map(int, input().split()))
sum = 0
arr.sort(reverse=True)

if len(arr) <= 1:
  sum = M*arr[0]
else:
  n1= arr[0]
  n2= arr[1]
  count = (M // (K+1))*K + M % (K+1)
  sum += count * n1
  sum += (M-count) * n2

print(sum)