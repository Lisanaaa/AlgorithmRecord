from collections import defaultdict

n, m = [int(i) for i in input().split(' ')]
lookup = defaultdict(list)
for i in range(1, m + 1):
    a, b = [int(j) for j in input().split(' ')]
    lookup[a].append(b)


# print(lookup)

def dist(a, b):  # 求a到b的距离
    return (b + n - a) % n


def helper(s, lookup):
    res = 0
    for i in range(1, n + 1):
        if not lookup[i]:
            continue
        minn = min([dist(i, j) for j in lookup[i]])
        res = max(res, dist(s, i) + (len(lookup[i]) - 1) * n + minn)
    return res


res = []
for i in range(1, n + 1):
    res.append(str(helper(i, lookup)))
print(' '.join(res))
