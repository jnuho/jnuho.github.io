


















# N = 25, K = 3
# 24    25 % 3 = 1:  count += mod(1)
# 8     24 // 3:  count += 1
# 7     8 % 3 = 2:  count += mod(2)
# 6     6 // 3: count+= 1
# 2     2 % 3 = 2 count+=mod(2)
# 1     
# result = 6

# 24
# 25 % 3 = 1
# 25 //3 = 8

# 8
# 8 % 3 = 2
# 8 // 3 = 2

# 6
# 2 % 3 = 2
# 2 // 3 = 0

N, K = map(int, input().split())
count = 0
mod = 0
while N > 1:
  mod = (N % K)
  N -= mod
  count += (mod + (N//K >0))
  N //= K

if mod > 0:
  count -= 1
print(count)

# N이 1이 될때까지
# (1번) N = N-1
# (2번) N을 K로 나눈다
# 2<=N<=100000, 2<=K<=100000,  N >= K
# 1이 될때까지 1또는 2번 수행해야하는 최소 횟수
# 17 4
# 3
# 25 5
# 2
# 25 3
# 6