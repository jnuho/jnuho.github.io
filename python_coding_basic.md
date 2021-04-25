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
b = round(a, 4)
print(a)
print(b)

```

    1000000000.0
    752.5
    3.954
    0.8999999999999999
    0.9
    

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
# 리스트 인덱싱, 슬라이싱
a = []
print(a)

n = 10
a = [0]*n
print(a)

a = [1,2,3,4,5,6,7,8,9]
print(a)

print(a[-1])
print(a[-3])
print(a[1:4]) # [2,3,4]
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
    


```python
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
```


```python
# 리스트 메소드
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
  - 코딩테스트
  - '원소 in 사전' (리스트,튜플에도 있는 메소드)


```python
data = dict()
data['사과'] = 'Apple'
data['바나나'] = 'Banana'
data['코코넛'] = 'Coconut'

# '사과'를 키로 가진 데이터 존재
if '사과' in data:
  print(data)

key_list = data.keys() # 키 데이터만 담은 리스트
value_list = data.values() # 값 데이터만 담은 리스트
print(key_list)
print(value_list)

for key in key_list:
  print(key, ":", data[key])
```

    {'사과': 'Apple', '바나나': 'Banana', '코코넛': 'Coconut'}
    dict_keys(['사과', '바나나', '코코넛'])
    dict_values(['Apple', 'Banana', 'Coconut'])
    사과: Apple
    바나나: Banana
    코코넛: Coconut
    

- 집합 자료형
  - 중복을 허용하지 않는다.
  - 순서가 없다 (인덱싱으로 value 얻을 수 없음. 키존재 X)
  - 특정 원소 존재 확인 O(1)



```python
# 초기화 방법
data = set([1, 1, 2, 3, 4, 4, 5])
data = {1, 1, 2, 3, 4, 4, 5}
print(data)

# 집합 자료형의 연산
a = {1, 2, 3, 4, 5}
b = {3, 4, 5, 6, 7}

print(a | b) # 합집합
print(a & b) # 교집합
print(a - b) # 차집합

# 집합 메소드
# add(), remove() O(1)
# update() 여러개의 값을 한번에 추가
data = {1, 2, 3}
print(data)

data.add(4)
print(data)

# 새로운 원소 여러개 추가
data.update([5, 6])
print(data)

# 특정한 값을 갖는 원소 삭제
data.remove(3)
print(data)

```

    {1, 2, 3, 4, 5}
    {1, 2, 3, 4, 5, 6, 7}
    {3, 4, 5}
    {1, 2}
    {1, 2, 3}
    {1, 2, 3, 4}
    {1, 2, 3, 4, 5, 6}
    {1, 2, 4, 5, 6}
    

- 조건문
  - if cond1: do1
  - elif cond2: do2
  - elif cond3: do3
  - else: do4

- 비교 연산자
  - x == y
  - x != y
  - x > y
  - x < y
  - x >= y
  - x <= y

- 논리 연산자
  - x and y
  - x or y
  - not x

- 기타 연산자
  - x in 리스트
  - x not in 문자열
  - pass 조건문 형태만 만들어 놓음 ```if cond: pass```
  - 줄바꿈 없이 작성 ```result = "Success" if score >= 80 else "Fail"```
  - 조건문에 수학 부등식 가능 ```if 0 < x < 20: print('True')```


```python
# 한줄로 conditional expression
score = 85
result = "Success" if score >= 80 else "Fail"
print(result)


a = [1,2,3,4, 5,5,5]
remove_set = {3, 5}
result = [i for i in a if i not in remove_set]
print(result)
```

    Success
    [1, 2, 4]
    

- 반복문
  - for
  - while


```python
# 1~9 홀수의 합
i = 1
result = 0
while i<= 9:
  if i %2 == 1:
    result += i
  i += 1
print(result)

# for 문
scores = [90, 85, 77, 65, 97]
black_list = {2, 4}
for i in range(5):
  if i+1 in black_list:
    continue
  if scores[i] >= 80:
    print(i+1, "번 학생은 합격입니다.")

# 구구단
for i in range(2,4):
  for j in range(1, 10):
    print(i, "X", j, "=", i*j)
  print()
```

    25
    1 번 학생은 합격입니다.
    5 번 학생은 합격입니다.
    2 X 1 = 2
    2 X 2 = 4
    2 X 3 = 6
    2 X 4 = 8
    2 X 5 = 10
    2 X 6 = 12
    2 X 7 = 14
    2 X 8 = 16
    2 X 9 = 18
    
    3 X 1 = 3
    3 X 2 = 6
    3 X 3 = 9
    3 X 4 = 12
    3 X 5 = 15
    3 X 6 = 18
    3 X 7 = 21
    3 X 8 = 24
    3 X 9 = 27
    
    

- 함수


```python
# 함수 call
def divide(a,b):
  return a/b

print(divide(3,7))
print(divide(b=3,a=7))
```

    0.42857142857142855
    2.3333333333333335
    10
    10
    10
    


```python
# global 변수
# 함수 안에서 함수 밖의 변수 데이터를 변경해야 할 때 global 선언 하면
# 해당 함수에서 지역 변수를 만들지 않고, 바깥에 선언된 변수를 바로 참조
a = 0
def func():
  global a
  a += 1

for i in range(10):
  func()

print(a)
```

    10
    


```python
# 일반 표현식
def add(a, b):
  return a+b

print(add(3,7))

# 람다 표현식
# 정렬 라이브러리 사용시 정렬 기준 key 설정할 때 자주 사용 : Chapter6. sort
print((lambda a,b: a+b)(3,7))
```

- 입출력
  - 입력 예시 :
  ```
  5
  65 90 75 34 99
  ```
  - ```list(map(int, input().split()))``` 사용
    - 입력값을 공백 구분 리스트로 바꾸고, 모든 원소 int 전환하여 다시 리스트 형 변환


```python
# 데이터 개수 5
n = int(input("데이터 수"))

# 정수 입력(공백구분) 65 90 75 34 99
data = list(map(int, input(f"정수 {n}개").split()))

data.sort(reverse=True)
print(data)

```

    [2]
    


```python
# n, m, k 공백 구분하여 입력
n, m, k = map(int, input().split())
print(n, m, k)

```

    32 23 1
    


```python
# 입력 개수가 많을 경우 느려짐
# 파이썬의 경우 동작 드린 input() 보다
# sys 라이브러리 sys.stdin.readline() 사용
# 한 줄 씩 입력 받음
import sys

# 5 6
# 1 3 5 6 9
# 5개 수 중에 6보다 작은 수들 출력
# Answer: 1,3 5
N,x = map(int, sys.stdin.readline().strip().split())

scores = [i for i in sys.stdin.readline().strip().split() if int(i) < x]
print(' '.join(scores))
```


    ---------------------------------------------------------------------------

    ValueError                                Traceback (most recent call last)

    <ipython-input-1-17fe48d9bb7d> in <module>
          5 import sys
          6 
    ----> 7 N,x = map(int, sys.stdin.readline().strip().split())
          8 
          9 scores = [i for i in sys.stdin.readline().strip().split() if int(i) < x]
    

    ValueError: not enough values to unpack (expected 2, got 0)


- 출력 예시


```python
a = 1
b = 2
print(a,b)
print("a는 " + str(a)+ ", b는 " + str(b))
print("a는", a, ", b는", b)
print(f"a는 {a}, b는 {b}")
```

    1 2
    a는 1, b는 2
    a는 1 , b는 2
    a는 1, b는 2
    

- 주요 라이브러리 문법
  - https://docs.python.org/ko/3/library/index.html

- 내장함수: print() input() sorted() 같은 기본 내장 라이브러리
- itertools: 반복되는 형태 데이터 처리 기능제공. 순열과 조합 라이브러리 제공
- heapq: 힙 기능을 제공. 우선순위 큐 기능을 구현
- bisect: 이진 탐색(Binary Search) 기능 제공
- collections: deque, Counter 등 유용한 자료구조를 포함
- math: 필수적인 수학적 기능 제공. 팩토리얼, 제곱근, 최대공약수GCD, 삼각함수메소드, PI 등


```python
# 내장함수
result = sum([1, 2, 3, 4, 5])
print(result)

result = max(7, 3, 5, 2)
print(result)

result = eval("(3+5) * 7")
print(result)

result = sorted([9, 1, 8, 5, 4])
print(result)

result = sorted([9, 1, 8, 5, 4], reverse=True)
print(result)

# 두번째 원소 x[1] 기준 내림차순
result = sorted([('홍길동',35), ('이순신',75), ('아무개',50)], key = lambda x: x[1], reverse=True)
print(result)

# sort()는 내부 데이터도 바꿈
data = [9, 1, 8, 5, 4]
data.sort()
print(data)
```

    15
    7
    56
    [1, 4, 5, 8, 9]
    [9, 8, 5, 4, 1]
    [('이순신', 75), ('아무개', 50), ('홍길동', 35)]
    [1, 4, 5, 8, 9]
    


```python
# itertools
from itertools import permutations
from itertools import combinations
from itertools import product
from itertools import combinations_with_replacement

# 2개 뽑아 순서O 나열
data = ['A', 'B', 'C']
result = list(permutations(data, 2))
print(f"3P3: {result}")

# 2개 뽑아 순서X 나열
data = ['A', 'B', 'C']
result = list(combinations(data, 2))
print(f"3C2: {result}")

# 2개 뽑아 순서O 나열 (원소 중복)
data = ['A', 'B', 'C']
result = list(product(data, repeat=2))
print(f"3P2 (repeat=2): {result}")

# 2개 뽑아 순서X 나열 (원소 중복)
data = ['A', 'B', 'C']
result = list(combinations_with_replacement(data, 2))
print(f"3C2 (repeat=2): {result}")

```

    3P3: [('A', 'B'), ('A', 'C'), ('B', 'A'), ('B', 'C'), ('C', 'A'), ('C', 'B')]
    3C2: [('A', 'B'), ('A', 'C'), ('B', 'C')]
    3P2 (repeat=2): [('A', 'A'), ('A', 'B'), ('A', 'C'), ('B', 'A'), ('B', 'B'), ('B', 'C'), ('C', 'A'), ('C', 'B'), ('C', 'C')]
    3C2 (repeat=2): [('A', 'A'), ('A', 'B'), ('A', 'C'), ('B', 'B'), ('B', 'C'), ('C', 'C')]
    

- heapq
  - heap 기능을 위해 사용
  - 다익스트라 최단경로 알고리즘 포함, 다양한 알고리즘에서
  - 우선순위 큐 기능을 구현 시 사용
  - PriorityQueue도 있지만, 코딩테스트 환경에서는 heapq가 더 빠르게 동작
  - 파이썬 힙은 Min Heap으로 구성되어 있으므로, 원소를 힙에 전부 넣고 뺄때 시간복잡도 O(NlogN)이 걸리며 오름차순 정렬 됨
  - 힙 자료구조의 최상단 원소는 항상 '가장 작은' 원소
  - 원소 삽입 heapq.heappush()
  - 원소 꺼냄 heapq.heappop()

- Heap Sort를 heapq로 구현 하는 예제


```python
import heapq

# min heap
def heapsort(iterable):
  h = []
  result = []
  # 모든 원소를 차례대로 힙에 삽입
  for value in iterable:
    heapq.heappush(h, value)
  # 힙에 삽입된 모든 원소를 차례대로 꺼내어 담기
  for i in range(len(h)):
    result.append(heapq.heappop(h))

  return result

result = heapsort([1, 3, 5, 7, 9, 2, 4, 6, 8, 0])
print(result)

# max heap
# 파이썬은 최대힙 max heap을 제공 하지 않음
# heapq 라이브러리를 사용해서 구하려면, 임의로 원소 부호를 바꿔서 함
# 힙에 원소 삽입 전 잠시 부호 바꿨다가, 힙에서 원소 꺼낸 뒤 다시 원소 부호 바꾸면 됨
def heapsort_max(iterable):
  h = []
  result = []
  # 모든 원소를 차례대로 힙에 삽입
  for value in iterable:
    heapq.heappush(h, -value)
  # 힙에 삽입된 모든 원소를 차례대로 꺼내어 담기
  for i in range(len(h)):
    result.append(-heapq.heappop(h))

  return result

result = heapsort_max([1, 3, 5, 7, 9, 2, 4, 6, 8, 0])
print(result)

```

    [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
    [9, 8, 7, 6, 5, 4, 3, 2, 1, 0]
    [0, 1, 2, 6, 3, 5, 4, 7, 8, 9]
    

- bisect
  - 이진탐색 구현한 라이브러리
  - '정렬된 배열'에서 특정한 원소를 찾아야 할 때 매우 효과적
  - 주요메소드 bisect_left() bisect_right() 시간 복잡도 O(logN)
    - bisect_left(a,x) : 정렬순서 유지하면서, 리스트 a에 데이터 x를 삽입할 가장 왼쪽 인덱스 찾음
    - bisect_right(a,x) : 정렬순서 유지하면서, 리스트 a에 데이터 x를 삽입할 가장 오른쪽 인덱스 찾음
      - a= [1,2,4,4,8] 일때, bisect_left(a,4)==2, bisect_right(a,4)==4
    - '정렬된 리스트'에서 '값이 특정범위에 속하는 원소의 개수' 구할 때 효과적: 시간 O(logN)


```python
from bisect import bisect_left, bisect_right

a = [1,2,4,4,8]
x = 4
print(bisect_left(a,x))
print(bisect_right(a,x))
```

    2
    2
    


```python
from bisect import bisect_left, bisect_right

# '정렬된 리스트'에서 '값이 특정범위에 속하는 원소의 개수' 구할 때 효과적: O(logN)
# 값이 [left_value, right_value]인 데이터 개수를 반환
def count_by_range(a, left_value, right_value):
  right_index = bisect_right(a, right_value)
  left_index = bisect_left(a, left_value)
  return right_index - left_index

# 리스트 선언
a = [1,2, 3,3,3,3, 4,4, 8,9]

# 값이 4인 데이터 개수 출력
print(count_by_range(a, 4,4))

# 값이 [-1,3] 범위에 있는 데이터 개수 출력
print(count_by_range(a, -1, 3))
```

    2
    6
    

- collections: 유용한 자료구조 제공 라이브러리
  - list: append(), pop()은 가장 뒤쪽 원소 기준. 앞쪽 원소 제거시 O(N) 오래걸림
  - deque: 큐 구현 (Queue는 자료구조 라이브러리 아님)
    - 리스트 같은 인덱싱, 슬라이싱은 사용할 수 없지만, 데이터 시작, 끝부분 삽입,삭제 효과적
    - 스택이나 큐의 기능 포함하기 때문에 이들 자료구조를 대체하여 쓸 수 있음
    - 처음제거 popleft(), 마지막제거 pop(), 첫삽입 appendleft(x), 마지막삽입 append(x)
    - 큐 구조는 append()와 popleft()를 사용하여 구현
  - Counter: 리스트 같은 iterable 객체 내부에 해당원소가 몇번 등장하는지 계산
  - math

설명 |리스트|deque
---|---|---
가장 앞쪽 추가 | O(N) | O(1)
가장 뒤쪽 추가 | O(1) | O(1)
가장 앞쪽 제거 | O(N) | O(1)
가장 뒤쪽 제거 | O(1) | O(1)


```python
from collections import deque

data = deque([2,3,4])
# 처음과 마지막 삽입
data.appendleft(1)
data.append(5)

print(data)
print(list(data)) # 리스트형 변환
```

    deque([1, 2, 3, 4, 5])
    [1, 2, 3, 4, 5]
    


```python
from collections import Counter

counter = Counter(['red', 'blue', 'red', 'green', 'blue', 'blue'])

print(counter['blue']) # blue 등장 횟수
print(counter['green']) # green 등장 횟수
print(dict(counter)) # 사전 자료형으로 변환
```

    3
    1
    {'red': 2, 'blue': 3, 'green': 1}
    


```python
import math

print(math.factorial(5)) # 5 팩토리얼
print(math.sqrt(7)) # 7의 제곱근
print(math.gcd(21, 14)) # 최대 공약수 GCD
print(math.pi) # pi
print(math.e) # 자연상수 e
```

    120
    2.6457513110645907
    7
    3.141592653589793
    2.718281828459045
    


```python
# 깃허브 알고리즘 노트
# 2020년 카카오 기출문제 '자물쇠와 열쇠'
''' 2차원 리스트(행렬)를 90도 회전 '''

def rotate_a_matrix_by_90_degree(a):
  row = len(a)
  col = len(a[0])
  res = [[0]*row for _ in range(col)]

  for r in range(row):
    for c in range(col):
      res[c][row-1-r] = a[r][c]

  return res

a = [
  [1,2,3,4]
  , [5,6,7,8]
  , [9,10,11,12]
]
print(rotate_a_matrix_by_90_degree(a))
```

    [[9, 5, 1], [10, 6, 2], [11, 7, 3], [12, 8, 4]]
    
