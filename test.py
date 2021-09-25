import sys

N = int(sys.stdin.readline().rstrip())
arr1 = list(map(int, sys.stdin.readline().rstrip().split()))
arr1.sort()

M = int(sys.stdin.readline().rstrip())
arr2 = list(map(int, sys.stdin.readline().rstrip().split()))

def binary_search(array, target, start, end):
  while start <= end:
    mid = (start + end) // 2

    if target == arr1[mid]:
      return mid
    elif target > arr1[mid]:
      start=mid+1
    else:
      end = mid-1
  return None

for i in arr2:
  result = binary_search(arr1, i, 0, N-1)
  if result != None:
    print("yes", end= ' ')
  else:
    print("no", end= ' ')