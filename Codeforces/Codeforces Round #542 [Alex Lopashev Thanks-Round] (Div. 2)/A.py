n = int(input())
a = [int(i) for i in input().split(' ')]


def helper(a):
    pos, neg = 0, 0
    for i in a:
        if i > 0:
            pos += 1
        elif i < 0:
            neg += 1
    if pos >= n / 2:
        return 1
    elif neg >= n / 2:
        return -1
    else:
        return 0


print(helper(a))
