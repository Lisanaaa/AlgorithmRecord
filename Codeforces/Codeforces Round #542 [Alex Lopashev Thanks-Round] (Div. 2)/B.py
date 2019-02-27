“”“
每一轮的选择跟下一轮都没有关系

Let 𝑥𝑖 and 𝑦𝑖 be the index of the cakes with size 𝑖.

We know that if either Sasha or Dima walks from 𝑥𝑖 to 𝑥𝑖+1 
then the other one will have no choice but to walk from 𝑦𝑖 to 𝑦𝑖+1. 
Similarly, if one walks from 𝑥𝑖 to 𝑦𝑖+1 then the other one will have to walk from 𝑦𝑖 to 𝑥𝑖+1.

Therefore, 𝐶𝑖 the total distance of moving from cake of size 𝑖 to 𝑖+1 is 𝑚𝑖𝑛(|𝑥𝑖−𝑥𝑖+1|+|𝑦𝑖−𝑦𝑖+1|,|𝑥𝑖−𝑦𝑖+1|+|𝑦𝑖−𝑥𝑖+1|).

The answer is ∑𝑖=1𝑛−1𝐶𝑖.

The time complexity is (𝑛).

”“”


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
        xi, xj, yi, yj = lookup[i] + lookup[i + 1]
        res += min(abs(xi - yi) + abs(xj - yj), abs(xi - yj) + abs(xj - yi))
    return res


print(helper(a, n))
