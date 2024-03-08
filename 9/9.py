def min(n, w):
    w.sort()  
    min_ins = float('inf')

   
    for i in range(n):
        ins = abs(w[i] - w[2*n - i - 1])
        min_ins = min(min_ins, ins)

    return min_ins

n = int(input())
w = list(map(int, input().split()))


print(min(n, w))
