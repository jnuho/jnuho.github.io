N, K = map(int, input().split())

A = list(map(int, input().split()))
B = list(map(int, input().split()))

A.sort()
B.sort(reverse=True)
for i in range(N):
  tempA = A[i]
  tempB = B[i]
  A[i] = tempB
  B[i] = tempA

print(sum(A))