
array = [3, 5, 1, 2, 4]

summary = 0

# O(N)
for x in array:
    summary += x

# O(N^2)
for i in array:
    for j in array:
        temp = i*j

