






# top left (1,1) bottom right (N,N)
# 계획서 띄어쓰기 기준 L R U D 문자들이 반복적으로 적혀있음
# 움직일수 없는 곳 이동명령은 무시 됨
# 5
# R R R U D D
# 3 4
move_type = ['L', 'R', 'U', 'D']

# LRUD
r_moves = [0, 0, 1, -1]

# 행열
c_moves = [-1, 1, 0, 0]
r, c = 1,1
n = int(input())
moves = list(input.split()) # LRUD 방향 리스트

