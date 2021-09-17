
# 선택 정렬
def selection_sort(arr):
  for i in range(len(arr)):
    min_idx = i
    for j in range(i+1, len(arr)):
      if arr[j] < arr[min_idx]:
        arr[j], arr[min_idx] = arr[min_idx], arr[j]
  print(arr)

# 삽입 정렬
def insertion_sort(arr):
  for i in range(len(arr)):
    for j in range(i,0,-1):
      if arr[j] < arr[j-1]:
        arr[j], arr[j-1] = arr[j-1], arr[j]
      else:
        break 
  print(arr)

# 퀵 정렬
# def quick_sort(arr, start, end):
#   if start >= end:
#     return
#   pivot = start
#   left = start+1
#   right = end

#   while left <= right:
#     while left <= end and arr[left] <= arr[pivot]:
#       left += 1
#     while right > start and arr[right] >= arr[pivot]:
#       right -=1
#     if left > right:
#       arr[right], arr[pivot] = arr[pivot], arr[right]
#     else:
#       arr[right], arr[left] = arr[left], arr[right]

#   quick_sort(arr, start, right-1)
#   quick_sort(arr, right+1, end)
def quick_sort(arr):
  if len(arr) <= 1:
    return arr
  
  pivot = arr[0]
  tail = arr[1:]

  left_side = [x for x in tail if x <= pivot]
  right_side = [x for x in tail if x > pivot]

  return quick_sort(left_side) + [pivot] + quick_sort(right_side)

def counting_sort(arr):
  count_arr = [0] * (max(arr)+1)

  for i in arr:
    count_arr[i] += 1

  for i in range(len(count_arr)):
    for j in range(count_arr[i]):
      print(i, end=' ')
  print()



arr = [7, 5, 9, 0, 3, 1, 6, 2, 4, 8]
selection_sort(arr)

arr = [7, 5, 9, 0, 3, 1, 6, 2, 4, 8]
insertion_sort(arr)

arr = [5, 7, 9, 0, 3, 1, 6, 2, 4, 8]
arr = quick_sort(arr)
print(arr)

arr = [7, 5, 9, 0, 3, 1, 6, 2, 9, 1, 4, 8, 0, 5, 2]
counting_sort(arr)

# 파이썬 정렬 라이브러리!
arr = [7, 5, 9, 0, 3, 1, 6, 2, 9, 1, 4, 8, 0, 5, 2]
arr.sort()
print(arr)
arr = [7, 5, 9, 0, 3, 1, 6, 2, 9, 1, 4, 8, 0, 5, 2]
print(sorted(arr))

# 정렬 기준 설정
arr = [('바나나', 2), ('사과', 5), ('당근', 3)]

# 각데이터의 두번째 원소 기준 - 2,3,5
def setting(data):
  return data[1]

result = sorted(arr, key=setting)