def read_int():
    return int(raw_input())
def read_ints():
    return [int(x) for x in raw_input().split()]

def solve(N, C, M):
    chocolates = N // C
    wrappers = chocolates
    while wrappers >= M:
        bonus = wrappers // M
        chocolates += bonus
        wrappers -= bonus * (M-1)
    return chocolates

if __name__ == "__main__":
    T = read_int()
    for t in xrange(T):
        N, C, M = read_ints()
        ans = solve(N, C, M)
        print ans
