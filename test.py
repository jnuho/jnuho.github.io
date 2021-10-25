
# 5, [1,2,3,4,5]
# 3, [1,7,5]
# N = int(input())
# n_list = set(map(int, input().split()))

# M = int(input())
# m_list = map(int, input().split())

def binary_search(arr:list, target:int, start:int, end:int) -> int:

  while start <= end:
    mid = (start + end) // 2

    if arr[mid] == target: 
      return mid
    elif arr[mid] > target:
      end = mid - 1
    else:
      start = mid + 1

  return None


# N 떡개수
# M 타겟길이
# dd 리스트
N, M = map(int, input().split())
dd = list(map(int, input().split()))
dd.sort()

start = 0
end = len(dd) -1
while start <= end:
  mid = (start + end) // 2
  for i in dd:

  

# ans = binary_search(dd, M, start, end)
# print(ans)

