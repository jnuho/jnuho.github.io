


















# 체스판에서 (a~h, 1~8)
# 나이트 위치가 주어졌을때, 이동 가능한 경우의 수

user_input = input()
x, y  = int(ord(user_input[0])) - 96, int(user_input[1])
moves = [(2,1), (2,-1), (1,2), (-1,2), (-2,1), (-2,-1), (1,-2), (-1,-2)]

count = 0
for m in moves:
  new_x = m[0] + x
  new_y = m[1] + y

  if new_x >= 1 and new_x <= 8 and new_y >= 1 and new_y <=8:
    count += 1

print(count)