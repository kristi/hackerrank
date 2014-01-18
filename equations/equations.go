package main
import "fmt"

// return a list of primes less than N
func sieve(N int) (primes []int) {
    // Mark composite numbers as true
    c := make([]bool, N)
    for i:=2; i<N; i++ {
        if c[i] == true {
            // i is composite
            continue
        }
        // i is prime
        primes = append(primes, i)
        // remove multiples of i
        for k:=2*i; k<N; k+=i {
            c[k] = true
        }
    }
    return
}

func readInt() (x int) {
    fmt.Scan(&x)
    return
}

func factorizeFactorial(N int) (F []int) {
    primes := sieve(N)
    for _,p := range primes {
        f := 0
        for d:=p; d<=N; d*=p {
            f += N/d
        }
        F = append(F, f)
    }
    return
}

func solve(N int) (ans int) {
    M := 1000007
    ans = 1
    for _,p := range sieve(N+1) {
        v := 0
        for d:=p; d<=N; d*=p {
            v += N/d
        }
        x := (2 * v + 1) % M
        ans = (ans * x) % M
    }
    return
}

func main() {
    for _,n := range []int{1,2,3,4,5,32327,40921} {
        fmt.Println(n, solve(n))
    }
}
