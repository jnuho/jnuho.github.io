# a~h, 1~8
pos = input()
x= int(ord(pos[0]))-int(ord('a'))+1
y = int(pos[1])
moves = [(2,-1),(2,1),(-2,1),(-2,1),(1,2),(1,-2), (-1,2),(-1,-2)]

count = 0
for move in moves:
  dx = x+move[0]
  dy = y+move[1]
  if dx >=1 and dx <=8 and dy>=1 and dy <=8:
    count +=1
print(count)