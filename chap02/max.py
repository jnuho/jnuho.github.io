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
    