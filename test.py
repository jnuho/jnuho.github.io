
from typing import List

def quick_sort(arr: List[int]) -> List[int]:
  if len(arr) <= 1:
    return arr

  idx = len(arr) // 2
  pivot = arr[idx]

  left = [x for x in arr if x < pivot]
  right = [x for x in arr if x > pivot]
  return quick_sort(left) + [pivot] + quick_sort(right)

arr = [7, 5, 9, 0, 3, 1, 6, 2, 4, 8]
print(arr)
print(quick_sort(arr))
