# 00:00:00 ~ N:59:59
# 3포함 개수
N = int(input())
count = 0
for h in range(0, N+1):
  for m in range(0, 60):
    for s in range(0, 60):
      if '3' in str(h) + str(m) + str(s):
        count+=1

print(count)
