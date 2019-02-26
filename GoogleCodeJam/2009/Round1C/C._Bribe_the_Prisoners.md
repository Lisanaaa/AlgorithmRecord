dp[i][j]表示的是将从A[i]号囚犯到A[j]号囚犯（not inclusive）的连续部分里的所有囚犯全部释放所需的最少金币数

```python
import sys


def helper(p, q, a):
    dp = [[0] * (q + 2) for _ in range(q + 1)]
    a = [0] + a[:] + [p + 1]
    for w in range(2, q + 2):
        for i in range(q + 2 - w):
            j = i + w
            t = sys.maxsize
            for k in range(i + 1, j):
                t = min(t, dp[i][k] + dp[k][j])
            dp[i][j] = t + a[j] - a[i] - 2
    # print(dp)
    return dp[0][-1]


for i in range(int(input())):
    p, q = [int(_) for _ in input().split(' ')]
    a = [int(_) for _ in input().split(' ')]

    print('Case #{}:'.format(i + 1), helper(p, q, a))

```
