package main
import (
    "fmt"
    "strings"
    "io"
    "bufio"
    "os"
    "strconv"
    "sort"
)

var _phi []int

func coprimes(x, y int) (count int) {
    // set x < y
    if y > x {
        x, y = y, x
    }
    if x == y {
        for i:=1; i<=x; i++ {
            count += _phi[i]
        }
        count = 2 * count - 1
        return
    }

    for i:=1; i<=x; i++ {
        m := y/x
        count += m * _phi[i]
        rem := y - m * i
        for j:=1; j<=rem; j++ {
            if i % j == 0 {
                count++
            }
        }
    }
    return
}

func solve(n []int) (ans int) {
    sort.Ints(n)
    a := n[0]
    c := make([]int, a+1)
    for i:=1; i<=a; i++ {
        c[i] = 1
    }

    for _,b := range n[1:] {
        // b >= a
        for i:=1; i<=a; i++ {
            c[i] = coprimes(a/i, b/i)
        }
    }

    for i,v := range c {
        ans += i * v
    }

    return
}

func main() {
    _phi = phi(100001)

    input := NewInputParser(os.Stdin)
    var T int
    input.Scan(&T)

    for t:=0; t<T; t++ {
        K := input.ParseInt()
        n := input.ParseIntArray(K)
        ans := solve(n)
        fmt.Println(ans)
    }

}
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

func phi(N int) (phi []int) {
    // Mark composite numbers as true
    phi = make([]int, N)
    for i := range phi {
        phi[i] = i
    }
    c := make([]bool, N)
    for i:=2; i<N; i++ {
        if c[i] == true {
            // i is composite
            continue
        }
        // i is prime
        p := i
        // mark multiples of p
        for k:=p; k<N; k+=p {
            phi[k] /= p
            phi[k] *= (p-1)
            c[k] = true
        }
    }
    return
}

func factorizeFactorial(N int) (F []int) {
    primes := sieve(N+1)
    for _,p := range primes {
        if p > N {
            break
        }
        f := 0
        for d:=p; d<=N; d*=p {
            f += N/d
        }
        F = append(F, f)
    }
    return
}
/* Prints an array without the brakets */
func printArray(a []int) {
    fmt.Println(strings.Trim(fmt.Sprint(a),"[]"))
}

type InputParser struct {
    *bufio.Reader
}

func NewInputParser(rd io.Reader) *InputParser {
    return &InputParser{ bufio.NewReader(rd) }
}

/* Returns a single line of input (without the ending newline) */
func (in *InputParser) ParseString() string {
    s,err := in.ReadString('\n')
    if err != nil && err != io.EOF { panic(err) }
    s = strings.TrimSpace(s)
    return s
}

/* Parses a single integer */
func (in *InputParser) ParseInt() int {
    s := in.ParseString()
    i,err := strconv.Atoi(s)
    if err != nil { panic(err) }
    return i
}

/* Scan for variables (pass in pointers)
    var a int
    var b float
    var c string
    in.Scan(&a, &b, &c)
*/
func (in *InputParser) Scan(a ...interface{}) int {
    s := in.ParseString()
    i,err := fmt.Sscan(s, a...)
    if err != nil { panic(err) }
    return i
}

/* Parses an array of integers separated by space */
func (in *InputParser) ParseIntArray(capacity int) []int {
    a := make([]int, 0, capacity)
    s := in.ParseString()
    for _,w := range strings.Split(s, " ") {
        n,err := strconv.Atoi(w)
        if err != nil { panic(err) }
        a = append(a, n)
    }
    return a
}
