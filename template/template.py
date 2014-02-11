def read_int():
    return int(raw_input())
def read_ints():
    return [int(x) for x in raw_input().split()]

def solve(N, M):
    return 0

if __name__ == "__main__":
    T = read_int()
    for t in xrange(T):
        N, M = read_ints()
        ans = solve(N, M)
        print ans
