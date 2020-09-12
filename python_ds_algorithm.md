[01. 알고리즘 기초](#01-알고리즘-기초)<br>
[02. 기본 자료구조와 배열](#02-기본-자료구조와-배열)<br>

# 01. 알고리즘 기초


```python
print('세정수 최대값')
a = int(input('a = '))
b = int(input('b = '))
c = int(input('c = '))

maximum = a
if b > maximum:
    maximum = b
if c > maximum:
    maximum = c

print(f'최대값 = {maximum}')
    
```

    세정수 최대값
    a = 2
    b = 3
    c = 1
    최대값 = 3



```python
# 10진수 문자열 -> 10진수 정수형 변환
int('17')
```




    17




```python
# 2진수 문자열 -> 10진수 정수형 변환
int('0b110', 2)
```




    6




```python
# 8진수 문자열 -> 10진수 정수형 변환
int('0o75', 8)
```




    61




```python
# 10진수 문자열 -> 10진수 정수형 변환
int('13', 10)
```




    13




```python
# 16진수 문자열 -> 10진수 정수형 변환
int('0x3F', 16)
```




    63




```python
# 문자열을 실수형 형변환
float('3.14')
```




    3.14




```python
def max3(a,b,c):
    """a,b,c의 최대값 반환"""
    max_val = a
    if b > max_val:
        max_val = b
    if c > max_val:
        max_val = c
        
    return max_val

max3(1,222,1255)
```




    1255



- [PEP 8 Guide](https://www.python.org/dev/peps/pep-0008)


```python
def med3(a, b, c):
    """a, b, c의 중간값 반환"""
    if a>= b and a<=c:
        return a
    if b >= a and b <= c:
        return b
    else:
        return c

print(med3(1, 8, 5))


def med3_better(a, b, c):
    """a, b, c의 중간값 반환"""
    if a>= b:
        if a <= c:
            return a
        elif b<=c:
            return c
        else:
            return b
    elif b <= c:
        return b
    elif a <= c:
        return c
    else:
        return a

print(med3_better(111,21, 222))
```

    5
    111



```python
# 삼항 연산자
x = 10
y = 5
a = x if x>y else y
print(a)

c = 0
print('c는 0입니다.'if c==0 else 'c는 0이 아닙니다.')
```

    10
    c는 0입니다.


- 반복하는 알고리즘


```python
def input_two_integer():
    a= int(input("정수 a : "))
    b= int(input("정수 b : "))
    return a, b


def sum_while_loop(a, b):
    if a > b:
        a, b = b, a    
        
    sum = 0
    i = a
    while i <= b:
        sum += a
        i += 1
        
    print(f'{a}~{b}의 합 (while)= {sum}')
    
    
def sum_for_loop(a, b):
    # for 반복문
    sum = 0
    if a > b:
        a, b = b, a
    for i in range(a, b):
        sum += i
        print(f'{i} + ', end = '')

    print(f'{b} = ', end='')
    sum += b
    print(f'{sum}')


a, b = input_two_integer() 
sum_while_loop(a, b)
        
a, b = input_two_integer() 
sum_for_loop(a, b)
```

    정수 a : 22
    정수 b : 11
    11~22의 합 (while)= 132
    정수 a : 22
    정수 b : 22
    22 = 22



```python
# +와 -를 번갈아 출력
n = int(input('+와 -를 번갈아 총 몇개 출력? '))

for _ in range(n // 2):
    print('+-', end='')
    
if n%2 == 1:
    print('+')
    

```

    +와 -를 번갈아 총 몇개 출력? 4
    +-+-


```python
# *를 n개마다 출력하되 w개 마다 줄바꿈하기
print('*를 출력합니다')
n = int(input('몇개를 출력 할까요? '))
w = int(input('몇 개마다 줄바꿈 할까요? '))

for _ in range(n//w):
    print('*'*w)

rest = n%w
if rest != 0:
    print('*'* (n%w))
```

    *를 출력합니다
    몇개를 출력 할까요? 2
    몇 개마다 줄바꿈 할까요? 1
    *
    *



```python
# 1~n 정수의 합
while True:
    n = int(input('정수 n? '))
    if n>0:
        break

sum = 0
for i in range(1, n+1):
    sum += i
    
print(f'합 = {sum}')
    
    
```

    정수 n? 33
    합 = 561



```python
# 넓이값이 주어졌을때, 정수 가로, 세로 길이 나열
area = int(input('직사각형의 넓이? '))

for i in range(1, area //2 +1):
    if i > area//i:
        break
    if area % i ==0:
        print(f'{i} x {area // i} = {area}')
        
```

    직사각형의 넓이? 33
    1 x 33 = 33
    3 x 11 = 33



```python
# 10~99사이 난수 n개 생성 (13나오면 중단)
# random.randint(a,b)  a이상 b이하 난수
import random
n = int(input('난수 n? '))
for _ in range(1, n+1):
    r = random.randint(10, 99)
    print(r, end=' ')
    if r== 13:
        break

```

    난수 n? 3
    15 15 83 


```python
# 드모르간
# if no >= 10 and no <= 99:
# if 10 <= no <= 99:
# if not(no < 10 or no > 99):

# 다중루프: 구구단
# :>n  fixed width n align right
# :^n  fixed width n align center
# :<n  fixed width n align left
print('-'*27)
for i in range(1, 10):
    for j in range(1, 10):
        print(f'{i*j :2}', end=' ')
    print()

print('-'*27)
```

    ---------------------------
     1  2  3  4  5  6  7  8  9 
     2  4  6  8 10 12 14 16 18 
     3  6  9 12 15 18 21 24 27 
     4  8 12 16 20 24 28 32 36 
     5 10 15 20 25 30 35 40 45 
     6 12 18 24 30 36 42 48 54 
     7 14 21 28 35 42 49 56 63 
     8 16 24 32 40 48 56 64 72 
     9 18 27 36 45 54 63 72 81 
    ---------------------------



```python
n = int(input('직삼각형 짧은 변의 길이? '))
# LB 직삼각형
# for i in range(1,n+1):
#     print('*'*i)
for i in range(n):
    for _ in range(i+1):
        print('*', end='')
    print()

# RB
# for i in range(n):
#     print(' '*(n-i-1), end='')
#     print('*'*(i+1))
for i in range(n):
    for _ in range(n-i-1):
        print(' ', end='')
    for _ in range(i+1):
        print('*', end='')
    print()
```

    직삼각형 짧은 변의 길이? 3
    *
    **
    ***
      *
     **
    ***



```python
# 파이썬 변수
# 정수 리터럴 17과 n의 식별변호가 같음 (n is binded with 17)
# 객체에 이름을 부여
n = 17
print(id(n))
print(id(17))

# 함수 내부,외부에서 정의한 변수와 객체의 식별 번호를 출력하기
n = 1 # 전역변수(함수 내부,외부에서 사용)
def put_id():
    x = 1 # 지역변수(함수 내부에서만 사용)
    print(f'id(x) = {id(x)}')
    
# 전역, 지역, 리터럴 =1 모두 id 같음
# n,x는 int형 객체 1을 참조하는 이름에 불과함
print(f'id(1) = {id(1)}')
print(f'id(n) = {id(n)}')
put_id();


# int 객체 10개 생성
for i in range(1, 11):
    print(f'i = {i:3}   id(i) = {id(i)}')
```

    4449181872
    4449181872
    id(1) = 4449181360
    id(n) = 4449181360
    id(x) = 4449181360
    i =   1   id(i) = 4449181360
    i =   2   id(i) = 4449181392
    i =   3   id(i) = 4449181424
    i =   4   id(i) = 4449181456
    i =   5   id(i) = 4449181488
    i =   6   id(i) = 4449181520
    i =   7   id(i) = 4449181552
    i =   8   id(i) = 4449181584
    i =   9   id(i) = 4449181616
    i =  10   id(i) = 4449181648


# 02. 기본 자료구조와 배열

학생들의 점수 var 1-n까지 저장하기 보다, 다음 요구사항을 해결하기 위해 배열 사용
- 학생수 n 변경가능
- 학생 i의 데이터 변경
- 최저, 최저 데이터 조회 및 정렬 필요

파이썬에서 배열은 1. 리스트 2.튜플 로 구현


```python
# 리스트
l = [] 
l = [1, 2, 3]
l = [1, 2, 3,]
l = list()
l = list([1,2,3])
l = list(range(1,3+1))

l[1] = 999

print(f'l[0] = {l[0]}');
print(f'l[len(l)-1] = {l[len(l)-1]}');
print(f'l[-len(1)//2] = {l[-len(l)//2]}');

l = [None]*5
print(l)

# 리스트 슬라이스
# s[i:j] s[i]~s[j-1]
# s[i:j:k] s[i]~s[j-1] k씩 건너뛰며
# s[:] 전체 리스트
# s[:n] 맨앞부터 n개
# s[i:] s[i]부터 맨끝
# [-n:] 맨끝에서 n개까지
# s[::k] 맨앞부터 k개씩 건너뛰며 출력
# s[::-1] 맨끝에서부터 전부 출력
l = list(range(20))
print(l[0:20:2])
print(l[:10])


# 튜플
t = ()
t = 1, # int 1이 아닌 튜플 (1)
t = (1,)
t = 1, 2, 3
t = (1, 2, 3,)
print(t)

# 값 받기
x = [1,2,3]
a,b,c = x # x를 언팩하여 a,b,c에 각각 assign
print(a)
print(b)
print(c)

# 뮤터블 자료형 : 리스트, 딕셔너리, 집합
# 이뮤터블 자료형 : 수, 문자열, 튜플
```

    l[0] = 1
    l[len(l)-1] = 3
    l[-len(1)//2] = 999
    [None, None, None, None, None]
    [0, 2, 4, 6, 8, 10, 12, 14, 16, 18]
    [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
    (1, 2, 3)
    1
    2
    3


- 자료구조의 개념

```데이터 단위와 데이터 자체 사이의 물리적 또는 논리적인 관계```


```python
w = [15, 64, 7]
x = [15, 64, 7, 3.14]
y = [15, 64, 7, 3.14, 2]
z = [15, 64, 7, 3.14]
print(f'len(x) = {len(x)}')
print(f'min(x) = {min(x)}')
print(f'max(x) = {max(x)}')

if x:
    print('비어있지 않은 리스트')
else:
    print('비어있는 리스트')
    
print(x == z)
print(w < x < y)

```

    len(x) = 4
    min(x) = 3.14
    max(x) = 64
    비어있지 않은 리스트
    True
    True



```python
# chap02/max.py
from typing import Any, Sequence

# Any : 제약이 없는 임의의 자료형
# Sequence : 시퀀스형 (list, bytearry, str, tuple, bytes)

# function주석 a는 Sequence타입, 리턴 타입 Any
def max_of (a: Sequence) -> Any:
    """시퀀스형 a원소의 최대값을 반환"""
    maximum = a[0]
    for i in range(1, len(a)):
        if a[i] > maximum:
            maximum = a[i]
    
    return maximum


# 스크립트 프로그램이 직접 실행될 때 변수 __name__은 '__main__' 입니다
# 스크립트 프로그램이 임포트될 때 변수 __name__은 원래의 모듈 이름 입니다
if __name__ == '__main__':
    n = int(input('배열 크기 n? '))
    l = [None] * n
    for i in range(n):
        l[i] = int(input(f'x[{i}] 값을 입력 : '))
        
    print(f'max = {max_of(l)}')
    
```

    배열 크기 n? 5
    x[0] 값을 입력 : 1
    x[1] 값을 입력 : 22
    x[2] 값을 입력 : -2
    x[3] 값을 입력 : 1
    x[4] 값을 입력 : 23
    max = 23



```python
from chap02.max import max_of

print('배열의 최대값을 구합니다 (q 입력 시 프로그램 종료)')

i = 0
arr = []
while True:
    v = input(f'x[{i}] = ')
    if v == 'q':
        break
    arr.append(int(v))
    i += 1

print(f'{i}개 입력완료')
print(f'최대값은 {max_of(arr)}')
```

    배열의 최대값을 구합니다 (q 입력 시 프로그램 종료)
    x[0] = 4
    x[1] = 2
    x[2] = -1
    x[3] = 33
    x[4] = 2
    x[5] = 0
    x[6] = q
    6개 입력완료
    최대값은 33



```python

```
