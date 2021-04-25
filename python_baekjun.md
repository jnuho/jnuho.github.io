## BAEKJUN 단계별


```python
# BAEKJOON 11022 A+B n번 호출
n = int(input())

for _ in range(n):
    A,B = map(int, input().split())
    print(A+B)
print()
```

    5
    1 2
    3
    3 5
    8
    2 3
    5
    1 2
    3
    5 6
    11
    
    


```python
# BAEKJOON 1110  N (0<=N<=99)의 사이클 길이를 출력
# (N=26) 2+6 = 8, 6 +8 = 14, 8+4 = 12, 4+2 = 6 : 사이클=4
n = int(input())
x = n
count = 0

while True:
  x = (x%10)*10 + (x//10 + x%10)%10
  count += 1
  if x == n:
    break

print(count)

```

    26
    4
    
