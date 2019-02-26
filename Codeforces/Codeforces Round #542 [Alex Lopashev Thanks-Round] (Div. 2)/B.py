from collections import defaultdict
n = int(input())
a = [int(i) for i in input().split(' ')]

def helper(a, n):
    lookup = defaultdict(list)
    for i, num in enumerate(a):
        lookup[num].append(i)
    lookup[0] = [0, 0]
    res = 0
    for i in range(n):
        xi, xj, yi, yj = lookup[i] + lookup[i+1]
        res += min(abs(xi-yi) + abs(xj-yj), abs(xi-yj) + abs(xj-yi))
    return res

print(helper(a, n))
