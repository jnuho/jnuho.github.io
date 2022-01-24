

















from typing import List

def insertion_sort(arr: List[int]) -> List[int]:
  for i in range(len(arr)):
    for j in range(i, 0, -1):
      if arr[j] < arr[j-1]:
        arr[j], arr[j-1] = arr[j-1], arr[j]
      else:
        break

  return arr

arr = [7, 5, 9, 0, 3, 1, 6, 2, 4, 8]
print(arr)
print(insertion_sort(arr))

