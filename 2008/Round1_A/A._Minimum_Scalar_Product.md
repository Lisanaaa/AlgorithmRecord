v1 升序排列，v2逆序排列


```python
for i in range(int(input())):
    n = int(input())
    v1 = [int(i) for i in input().split(' ')]
    v2 = [int(i) for i in input().split(' ')]
    v1.sort()
    v2.sort(reverse=True)
    res = sum(a * b for a, b in zip(*(v1, v2)))
    print("Case #{}:".format(i+1), res)
```
