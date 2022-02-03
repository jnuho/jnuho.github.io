








def binary_search(arr, target, start, end):
  if len(arr) < 1:
    return -1

  mid = (start + end) // 2

  if arr[mid] == target:
    return mid
  elif arr[mid] < target :
    return binary_search(arr, target, mid+1, end)
  else:
    return binary_search(arr, target, start, mid-1)

arr = [1,2,3,4,5,6,7,8,9,999, 100000]
target = 100000
idx = binary_search(arr, target, 0, len(arr)-1)
print(idx)
