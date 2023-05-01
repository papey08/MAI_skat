def solve_system_tridiagonal(self, b):
    to_use = True
    n = len(self)
    eq = 0
    leq = 0
    for i in range(1, n-1):
        if abs(self[i][1]) < abs(self[i][0]) + abs(self[i][2]) or \
                abs(self[i][1]) < abs(self[i-1][-1]) + abs(self[i+1][0]):
            eq += 1
            to_use = False
        else:
            leq += 1
    to_use *= leq < eq
    p, q = [0] * n, [0] * n
    ans = [0] * n
    p[0] = self[0][1] / -self[0][0]
    q[0] = b[0] / self[0][0]
    for i in range(1, n-1):
        p[i] = -self[i][2] / (self[i][1] + self[i][0]*p[i-1])
        q[i] = (b[i] - self[i][0]*q[i-1]) / (self[i][1] + self[i][0]*p[i-1])
    p[-1] = 0
    q[-1] = (b[-1] - self[-1][0]*q[-2]) / (self[-1][1] + self[-1][0]*p[-2])
    ans[-1] = q[-1]
    for i in range(n-1, 0, -1):
        ans[i-1] = p[i-1] * ans[i] + q[i-1]
    return ans
