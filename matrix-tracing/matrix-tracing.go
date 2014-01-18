package main
import (
    "fmt"
    "strings"
    "io"
    "bufio"
    "os"
    "strconv"
    "math/big"
)

var Primes []int

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

func factorizeFactorial(N int) (F []int) {
    for _,p := range Primes {
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

func binomialMod(n int, k int) (ans int) {
    nFact := factorizeFactorial(n)
    kFact := factorizeFactorial(k)
    jFact := factorizeFactorial(n-k)
    for i,v := range kFact {
        nFact[i] -= v
    }
    for i,v := range jFact {
        nFact[i] -= v
    }
    M := int64(1000000007)
    a := int64(1)
    for i,v := range nFact {
        if v == 0 {
            continue
        }
        a *= pow(int64(Primes[i]), int64(v))
        if a > M {
            a = a % M
        }
    }
    ans = int(a)
    return
}

func pow(x, y int64) (r int64) {
    if x == r || y < r {
        return
    }
    r = 1
    if x == r {
        return
    }
    if x < 0 {
        x = -x
        if y&r == r {
            r = -r
        }
    }
    for y > 0 {
        if y&1 == 1 {
            r *= x
        }
        x *= x
        y >>= 1
    }
    return
}

func solve(a []int) (ans int) {
    m := a[0]
    n := a[1]
    row := (m + n - 2)
    col := (m-1)

    ans = binomialMod(row, col)
    return
}

func solveBig(a []int) (ans int64) {
    m := a[0]
    n := a[1]
    row := int64(m + n - 2)
    col := int64(m-1)

    M := big.NewInt(1000000007)
    b := big.NewInt(0)
    b.Binomial(row, col)

    b.Mod(b, M)

    ans = b.Int64()
    return
}

func main() {
    Primes = sieve(2000000)
    input := NewInputParser(os.Stdin)
    var T int
    input.Scan(&T)

    for t:=0; t<T; t++ {
        a := input.ParseIntArray(2)
        ans := solve(a)
        fmt.Println(ans)
    }

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
