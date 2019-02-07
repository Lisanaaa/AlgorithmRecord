一次确定一行，根据每行最后一个1出现的位置，所以最后一个1出现在越前面的行就应该被换到越前面去

```python
def helper(grid):
    n = len(grid)
    res = 0
    a = [0] * n
    for i in range(n):
        a[i] = -1
        for j in range(n):
            if grid[i][j] == 1:
                a[i] = j
    for i in range(n):
        pos = -1
        for j in range(i, n):
            if a[j] <= i:
                pos = j
                break
        a = a[:i] + a[pos:pos+1] + a[i:pos] + a[pos+1:] # 调换第pos行和第i行
        res += pos - i # 共需要换 pos - i 次
    return res




for i in range(int(input())):
    n = int(input())
    grid = []
    for j in range(n):
        grid.append([int(_) for _ in input()])
    print('Case #{}:'.format(i+1), helper(grid))


```
