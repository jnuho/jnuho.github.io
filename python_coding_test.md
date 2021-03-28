# 00 파이썬 기초

- 자료형
  - 정수형 ```a= 1000```
  - 실수형 ```a= 3.1415 b= -.7 c=5. d=1e9 ```
  : 컴퓨터 수처리 방식은 2진수 이기 때문에 실수형을 저장할때, 4또는 8바이트라는
  고정된 크기의 메모리를 할당함. 예를들어 컴퓨터가 .3+.6=.9를 표현할때 미세한 오차가 존재
  (.89999...)


```python
a = 1e9
print(a)

b = 75.25e1
print(b)

c = 3954e-3
print(c)

a = .3 + .6
print(a)

b = round(a, 4)
if round(a, 4) == .9:
    print(True)
else:
    print(False)

```

    1000000000.0
    752.5
    3.954
    0.8999999999999999
    True
    

- 연산
  - 나누기 /
  - 나머지 %
  - 몫 //
  - 거듭제곱 **


```python
a = 7
b = 3

print(a / b)
print(a % b)
print(a // b)
print(a ** b)


```

    2.3333333333333335
    1
    2
    343
    

- 리스트

함수명 | 사용법 | 설명 | 시간복잡도
---|---|---|---
append() | 변수.append() | 리스트에 원소하나 삽입 | O(1)
sort() | 변수.sort() | 기본정렬기능 오름차순 | O(NlogN)
sort() | 변수.sort(reverse=True) | 내림차순정렬 | O(NlogN)
reverse() | 변수.reverse() | 리스트 원소순서 뒤집음 | O(N)
insert() | 변수.insert(index, value) | 특정 인덱스위치 value 삽입 | O(N)
count() | 변수.count(value) | 리스트에서 value 해당 데이터 수 | O(N)
remove() | 변수.remove(value) | 리스트에서 value 해당 '첫번째' 데이터만 제거. | O(N)


```python
a = []
print(a)

n = 10
a = [0]*n
print(a)

a = [1,2,3,4,5,6,7,8,9]
print(a)

# 리스트 인덱싱, 슬라이싱
print(a[-1])
print(a[-3])
print(a[1:4]) # [2,3,4]

# 리스트 컴프리 헨션
# 0~19중 홀수만
array = [i for i in range(20) if i%2 == 1]
print(array)
# 같은 결과
array = []
for i in range(20):
  if i%2 == 1:
    array.append(i)
print(array)

# 1~9 제곱값
array = [i*i for i in range(1,10)]
print(array)
# 2차원 리스트 N by M
n = 3
m = 3
array = [[0]*m for _ in range(n)]
print(array)

a = [1,4,3]
print("기본: ", a)

a.append(2)
print("a.append(2): ", a)

a.sort()
print("a.sort(): ", a)

a.sort(reverse=True)
print("a.sort(reverse=True): ", a)

a.reverse()
print("a.reverse(): ", a)

a.insert(1, 3)
print("a.insert(1, 3): ", a)

print("a.count(3): ", a.count(3))

a.remove(3)
print("a.remove(3) ", a)

# 특정 값 '모두' 제거
a = [1,2,3,4,5,5,5]
remove_set = {3, 5}
print(a)
a = [i for i in a if i not in remove_set]
print(a)
```

    []
    [0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
    [1, 2, 3, 4, 5, 6, 7, 8, 9]
    9
    7
    [2, 3, 4]
    [1, 3, 5, 7, 9, 11, 13, 15, 17, 19]
    [1, 3, 5, 7, 9, 11, 13, 15, 17, 19]
    [1, 4, 9, 16, 25, 36, 49, 64, 81]
    [[0, 0, 0], [0, 0, 0], [0, 0, 0]]
    기본:  [1, 4, 3]
    a.append(2):  [1, 4, 3, 2]
    a.sort():  [1, 2, 3, 4]
    a.sort(reverse=True):  [4, 3, 2, 1]
    a.reverse():  [1, 2, 3, 4]
    a.insert(1, 3):  [1, 3, 2, 3, 4]
    a.count(3):  2
    a.remove(3)  [1, 2, 3, 4]
    [1, 2, 3, 4, 5, 5, 5]
    [1, 2, 4]
    

- 문자열


```python
data = 'Hello World'
print(data)
data = "Dont't you know \"Python\"?"
print(data)

a = "Hello"
b = "World"
print(a+" "+b)
print(a*3)

a = "ABCDEF"
print(a[2:4]) # CD

```

    Hello World
    Dont't you know "Python"?
    Hello World
    HelloHelloHello
    CD
    

- 튜플
  - 한번 선언된 값 변경 불가능
  - 튜플은 소괄호 ```()```
  - 그래프 알고리즘 구현 시 자주 사용 (다익스트라 최단경로 알고리즘에서, 우선순위 queue에 한번 들어간 값은 변경 되지 않기 때문에, 우선순위 큐에 들어가는 데이터는 튜플로 선언)
  - 다익스트라 최단경로 알고리즘에서 (비용, 노드번호) 형태로 묶어 관리
  - 리스트에 비해 공간 효율적
  - 각 원소 성질이 서로 다를때 주로 사용


```python
a = (1,2,3,4)
# a[2] = 7 # ERROR!

```

- 사전
  - key,value 쌍
  - 내부적으로 Hash Table을 이용하므로 데이터 검색, 수정은 O(1) 리스트보다 훨씬 빠름


```python
data = dict()
data['사과'] = 'Apple'
data['바나나'] = 'Banana'
data['코코넛'] = 'Coconut'

print(data)
```

    {'사과': 'Apple', '바나나': 'Banana', '코코넛': 'Coconut'}
    

# 01 코딩테스트 개요

## 복잡도 Big-O
- 시간 복잡도 : 알고리즘을 위해 필요한 연산의 횟수
- 공간 복잡도 : 알고리즘을 위해 필요한 메모리의 양

Trade-off Example: memoization 메모리를 많이 사용하여 시간을 비약적으로 줄이는 방법 Ch8


```python
array = [3, 5, 1, 2, 4]
summary = 0

# O(N)
for x in array:
    summary += x

# O(N^2)
for i in array:
    for j in array:
        temp = i*j

```

$O(N)$ 최악의 경우로 복잡도 고려

|빅오 표기법 | 명칭 |
|:---|:-------|
|$O(1)$ | 상수시간 (Constant time)|
|$O(logN)$ |  로그시간 (Log Time) |
|$O(N)$ | 선형 시간  |
|O($NlogN$)| 로그 선형 시간 |
|$O(N^{2})$ |  이차 시간 |
|$O(N^{3})$ | 삼차 시간  |
|$O(2^{n})$ | 지수 시간 |

Big-O는 최악의 경우 시간복잡도 e.g.Quick Sort O(NlogN)~O(N^2)

## 시간과 메모리 측정


```python
import time

# 측정 시작
start_time = time.time()

'''
프로그램 소스코드
'''

# 측정 종료
end_time = time.time()

elapsed_time = end_time - start_time
print(f"Time elapsed : {elapsed_time}")
```

- 선택정렬의 경우 최악 $O(N^{2})$
- 파이썬 기본정렬 라이브러리 최악의 경우 $O(N\log{}N)$


```python
from random import randint
import time

# 배열에 10,000개의 정수 (a<= 랜덤 <=b) 삽입
# randint(a,b) 또는 randrange(a, b+1)
array = []
for _ in range(10000):
    array.append(randint(1, 100)) # 1<= i <=100 랜덤한 정수

# 선택 정렬 프로그램 성능 측정
start_time = time.time()

# 선택 정렬 프로그램 소스코드
for i in range(len(array)):
    # 가장 작은 원소의 인덱스
    min_index = i
    for j in range(i+1, len(array)):
        if array[min_index ] > array[j]:
            min_index = j
    # 스와프
    array[i], array[min_index] = array[min_index], array[i]

# 선택 정렬 측정 종료
end_time = time.time()
elapsed_time = end_time - start_time
print(f"선택 정렬 걸린 시간 : {elapsed_time}")


# 기본 정렬 라이브러리 성능 측정
start_time = time.time()

# 기본 정렬 라이브러리 사용
array.sort()

end_time = time.time()
elapsed_time = end_time - start_time
print(f"기본 정렬 걸린 시간 : {elapsed_time}")
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

| | 날짜 | 풀이 시간 | 문제 개수 | 커트라인 | 주요 문제 유형 | 시험 유형|
|:---|:---|:---|:---|:---|:---|:---|
|라인| 상반기<br>(2020-04-05)| 2시간 30분| 6문제| 4문제|구현, 문자열, 자료구조| 온라인|
|삼성전자| 상반기<br>(2020-06-07)| 3시간| 2문제| 2문제|완전탐색, 시뮬레이션, DFS/BFS| 오프라인|


# 03 Greedy 알고리즘


```python
# 가장 적은 수의 거스름돈
n = 1260
coins = [500, 100, 50, 10]

count = 0
for coin in coins:
    count += n // coin
    n %= coin

print(count)
```


```python
# 큰 수의 법칙
n, m, k = map(int, input().split())

data = list(map(int, input().split()))
data.sort(reverse=True)


```
