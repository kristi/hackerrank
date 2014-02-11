def read_int():
    return int(raw_input())
def read_ints():
    return [int(x) for x in raw_input().split()]

def solve(A, K):
    A.sort()
    unfairness = min(y - x for x,y in zip(A, A[K-1:]))
    return unfairness

if __name__ == "__main__":
    N = read_int()
    K = read_int()
    A = [read_int() for _ in xrange(N)]
    ans = solve(a, K)
    print ans
