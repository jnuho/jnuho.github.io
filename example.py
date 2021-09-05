change_amt = 1650
coins = [500, 100, 50, 10]

count = 0
for c in coins:
    count += change_amt // c
    change_amt %= c

print(count)