


def binary_search(arr, target, start, end):
  if start >= end:
    mid = (start + end) // 2

    if arr[mid] == target:
      return mid
    elif arr[mid] < target:
      binary_search(arr, target, mid+1, end)
    else:
      binary_search(arr, target, start, mid-1)

arr = [1,2,3,4,5,6,7,8,9,10]
target = 9

idx = binary_search(arr, 9, 0, len(arr)-1)
print(idx)
