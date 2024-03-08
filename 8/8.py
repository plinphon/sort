def total(n, m, cards):
    count = 0

    for i in range(n - 1):
        for j in range(i + 1, n):
            win = 0
            for k in range(m):
                win += abs(cards[i][k] - cards[j][k])
            count += win

    return count

t = int(input())
for _ in range(t):
    n, m = map(int, input().split())
    cards = []
    for _ in range(n):
        card = list(map(int, input().split()))
        cards.append(card)

    print(total(n, m, cards))
