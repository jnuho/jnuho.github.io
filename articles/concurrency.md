
Concurrency vs. Parallelism

- https://www.codeproject.com/Articles/1267757/Concurrency-vs-Parallelism
- https://stackoverflow.com/questions/1897993/what-is-the-difference-between-concurrent-programming-and-parallel-programming
- https://go.dev/tour/concurrency/1

If your program is using threads (concurrent programming), it's not necessarily going to be executed as such (parallel execution), since it depends on whether the machine can handle several threads.

1. Threads on a non-threaded machine:

```
        --  --  --
     /              \
>---- --  --  --  -- ---->>
```

2. Threads on a threaded machine:

```
     ------
    /      \
>-------------->>
```

The dashes represent executed code. As you can see, they both split up and execute separately, but the threaded machine can execute several separate pieces at once.


- Python concurrency
  - https://brownbears.tistory.com/292
  - GIL(Gloval Internal Lock)
  - Concurrent.futures

```python
"""
Concurrency in Python can be achieved using the threading library, which does create actual threads.
However, due to Python’s Global Interpreter Lock (GIL), these threads don’t run in true parallel on multiple cores.
Instead, they `context switch`, meaning only one thread executes at a time, and they take turns using the CPU.

Under the hood, when you create a thread in Python, it creates a separate flow of control within the program.
Each thread has its own program counter, stack, and local variables.
However, all threads within a process share global variables, allowing for communication between threads.

Here’s a simple example of creating threads in Python:
"""
import threading

def print_numbers():
    for i in range(10):
        print(i)

def print_letters():
    for letter in 'abcdefghij':
        print(letter)

# Create threads
t1 = threading.Thread(target=print_numbers)
t2 = threading.Thread(target=print_letters)

# Start threads
t1.start()
t2.start()

# Wait for both threads to finish
t1.join()
t2.join()
```

```
In this example, print_numbers and print_letters are running concurrently. However, due to the GIL, they are not truly running in parallel if you’re using standard CPython. If true parallelism is required, you might want to look into using processes with the multiprocessing module, or using a different Python interpreter like Jython or IronPython that doesn’t have a GIL. Alternatively, you can use native extensions like NumPy or TensorFlow that release the GIL when doing heavy computations.

Remember, concurrency is not the same as parallelism. Concurrency is about dealing with a lot of things at once (like handling a lot of I/O-bound tasks) while parallelism is about doing a lot of things at once (like CPU-bound tasks that benefit from multiple cores). Python’s threading is good for I/O-bound tasks, but for CPU-bound tasks, you might want to look into using multiprocessing or a native extension that can release the GIL.
```




- Context

While both Go and Python use the concept of a context, they apply it in different ways.
Go’s context is more about managing cancellation signals and carrying request-scoped data in concurrent programming,
while Python’s context is more about managing resources and ensuring proper cleanup in a runtime environment.

