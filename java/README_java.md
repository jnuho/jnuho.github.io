# udemy lecture note

# Data Structure & Algorithms
- Data Structure is a way to organize data in a way that enables it to be processed in an efficient time
  - Physical : array, linked list
  - Logical : stack, queue, tree, graph
  
- Algorithm is a set of rules to be followed to solve a problem

## Recursion
- Same operation is performed multiple times with different input
- In every step we try to make the problem smaller
- Mandatorily need to have a base condition, which tells system when to stop the recursion
- Makes code easy to write, whenever a given problem can be broken into similar sub-problem
- Heavily used in data structures like Tree, Graphs, etc used in 'divide and conquer', 'greedy', 'dynamic programming'

Recursive case : case where the function recur
Base case : case where the function does not recur
```
sampleRecur(param){
  if(base case is satisfied)
    return some base case value
  else
    sampleRecur(modified param)
}
```

```
// Example: Binary Search
search(root, valueToSearch){
  if(root == null) return null
  else if (root == valueToSearch) return root
  else if (valueToSearch < root) search(root.left, valueToSearch)
  else search(root.right, valueToSearch)
}
```

Stack memory managed by system :
```
  Main(){
    bar()
  }
  bar(){
    doWork();
  }
  doWork(){
    doMore();
  }
  doMore(){
    syso();
  }
  [Stack]  pop/push LIFO
    doMore()
    doWork()
    bar()
    Main()
```

```
// Example : Factorial
fact(n){
  if(n < 0) return 1
  else n*fact(n-1)
}
[Stack]
  foo(1)
  foo(2)
  foo(3)
  Main()
```

```
// Example : Fibonacci series
//   series of numbers in which each number is the sum of the two preceding numbers
//   e.g. 0 1 1 2 3 5 8 13 ...

fib(n) {
  if(n < 1)
    return error
  else if n ==1 or n==2
    return n-1
  else
    return fib(n-1) + fib(n-2)
}
```

### Recursion vs Iteration
- Space efficient? Recur X Iter O
  - recur: push recursive calls to stack memory, iter: no stack memory required
- Time efficient? Recur X Iter O
- Ease of code to solve sub-problem? Recur O Iter X

- When to use/avoid Recursion?
  if we can breakdown a problem into similar subproblems
  if ok with extra overhead(both time and space)
    - avoid : embedded system
    - avoid : mission critical project-  airbag software - fraction of seconds counts
    - use: quick working solution instead of efficient one

### Practical use of Recursion
- Stack
- Tree - Traversal/Searching/Insertion/Deletion
- Sorting - quick sort, merge sort
- Divide and Conquer
- Dynamic Programming - breakdown problems. save the answer and reuse for bigger problems
- ETC

## Algorithm Run Time Analysis
Study of a given algorithm's running time, by identifying its behavior, as the input size for the algorithm increases.
'How much time will the given algorithm will take to run' to measure efficiency of a given algorithm.

### Notations for 'Algorithm Run Time Analysis'
Car run 'km' per 1L -> city traffic? 10km highway? 20km mixed environment? 15km -> Notation for best/worst/avg case

- Omega : lower bound of a given algorithm
  layman's language : for any given input, running time of a given algorithm will not be 'less than' given time
  minimum required time
- Big-o (O)
  tighter upper bound of a given algorithm
  layman's language : for any given input, running time of a given algorithm will not be 'more than' given time
  e.g. airbag software How much seconds does the system takes to deploy the airbag in case of an accidents
        does not take more than 2 seconds
- Theta
  decides whether upper bound and lower bound of a given algorithm are same or not. in a very huge data set
  'on an average' be equal to given time

Example of Notations
```
  find #10 in an array of size 'n'
  Omega(1) - O(n) - Theta(n)
```

### Examples of 'Algorithm run time complexities' 
O(1) linear
O(logn) logarithmic
O(n) linear
O(nlogn) linear logarithmic
O(n^2) quadratic
O(n^3) cubic
O(2^n) exponential
O(n!) factorial

### How to calculate 'Algorithm Time Complexity'?
- Iterative Algorithm
```
int findBiggest(int arr[]){
  biggest = arr[0]; // O(1)
  for(i=0; i<arr.length; i++){ // O(n)
    if(arr[i] > arr[0]) //O(1)
      biggest = arr[i]
  }
  return biggest; //O(1)
}

// Time complexity
// O(1)+O(n) + O(1) + O(1) = O(n)
```

- Recursive Algorithm
```
// Example: find biggest num in an array of size n

int findBiggest(int arr[], n){ // T(n)
  static highest = Integer.Min // 0? O(1)
  if(n == -1) // O(1)
    return highest;
  else{
    if(A[n] > highest) //O(1)
      update highest
    return findBiggest(arr, n-1) // T(n-1)
  }
}

// Time complexity
  T(-1) = O(1)

  T(n) = O(1) + T(n-1)
  T(n-1) =  O(1) + T(n-2)
  T(n-2) =  O(1) + T(n-3)
  ...
  T(1) = O(1) + T(0)
  T(0) = O(1) + T(-1)
+ --------------
T(n) = (n+1)*O(1) + T(-1) = (n+2) * O(1) = O(n)
```

```
// Example : 10 20 30 40 50 [60] 70 80 90 100 110
// binary search - middel val = 60
// Find number 110. for a even number , mid point is n/2 or n/2+1

int binSearch(Number, int arr[], start, end){
  if(start equals end){ // base cond
    if(arr[start] == Number)
      return start;
    else
      return Exception
  }
  mid = Findmid(arr[], start, end)

  if(mid > Number)
    binSearch(Number, arr, start, mid)
  else if (mid < Number)
    binSearch(Number, arr, mid, end)
  else if (mid == Number)
    return mid
}

// Time complexity :
  T(1) = O(1)

  T(n) = O(1) + T(n/2)
  T(n/2) = O(1) + T(n/4)
  T(n/4) = O(1) + T(n/8)
  ...
  T(n) = 2*O(1) + T(n/4) = ... = k*O(1) + T(n/2^k) = logn *O(1) + T(1) = T(1) + logn = O(1) + logn
  [ n/2^k=1 <=> n= 2^k <=> logn = klog2 <=> k=logn ]
  : T(n) = O(logn)
```

## Array DS
- Array is a data structure consisting of a collection of elements, each identified by array index.
- An array is stored such that the position of each element can be computed from its index cell by a mathematical formula.
  - can store specified data type
  - contiguous memory locations
  - every cell has an unique index
  - idx starts with 0
  - size of array needs to be specified mandatorily and cannot be modified
```
1D array
Muli-dim-array (dim >= 2)
  arr[row][col]
    arr[3][7] has 3 rows and 7 columns and 21 elements
  arr[depth][row][col]
```

# 자료구조와 함께 배우는 알고리즘 입문

# 1장 기본 알고리즘

# 2장 기본 자료구조

## 배열
```java
배열.clone()
// 복제된 배열은 기존 배열을 참조하지 않음.
int[] a = {1,2,3,4,5};
int[] b = a.clone();
b[3] = 0;
for(int i : a) {
  System.out.println(i);
}
for(int i : b) {
  System.out.println(i);
}
```


# 3장 검색
# 4장 스택과 큐 - 1
# 4장 스택과 큐 - 2
# 5장 재귀 알고리즘 1
# 5장 재귀 알고리즘 2
# 6장 정렬 1
# 6장 정렬 2
# 7장 집합
# 8장 문자열 검색
# 9장 리스트 1
# 9장 리스트 2
# 10장 트리 1
# 10장 트리 2
# 11장 해시