package main
import (
    "fmt"
    "strings"
    "io"
    "bufio"
    "os"
    "strconv"
    "math"
)

func nChoose2(n int) int {
    if n <= 1 { return 0 }
    return (n * (n-1)) / 2
}

func sumInt(x []int) (ans int) {
    for _,v := range x {
        ans += v
    }
    return
}

func minInt(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func sum(x []float64) (ans float64) {
    for _,v := range x {
        ans += v
    }
    return
}

func solve(nSwaps int, nReverses int, d []int) (ans float64) {
    N := len(d)

    C := nChoose2(N)
    n := float64(N)
    c := float64(C)

    // Swap
    // d2[i] = z * d[i] + b
    d2 := make([]float64, N)
    {
        z := math.Pow((c-n)/c, float64(nSwaps))
        b := (1-z) / n * float64(sumInt(d))
        for i,di := range(d) {
            d2[i] = z * float64(di) + b
        }
    }

    // Reverse
    d3 := make([]float64, N)
    printArray(d2)
    for r:=0; r<nReverses; r++ {
        s := sum(d2)
        for i,di := range d2 {
            d3[i] = s - di + di * float64(nChoose2(i) + nChoose2(N-i-1))
        }
        for e:=1; e < N/2; e++ {
            s -= d2[e-1]
            s -= d2[N-e]
            for i:=e; i<N-e; i++ {
                d3[i] += s
            }
        }
        if N % 2 == 1 {
            d3[N/2] += d2[N/2]
        }
        for i := range d3 {
            d3[i] /= c
        }
        printArray(d3)
        copy(d2, d3)
    }

    // Sum
    for i,di := range d3 {
        g := C - nChoose2(i) - nChoose2(N-i-1)
        ans += float64(g) * di
    }

    ans /= c

    return
}

func main() {
    input := NewInputParser(os.Stdin)
    cases := 1

    for c:=0; c<cases; c++ {
        nab := input.ParseIntArray(3)
        n := nab[0]
        a := nab[1]
        b := nab[2]
        d := input.ParseIntArray(n)

        ans := solve(a,b,d)
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
    if err != nil { panic(err) }
    s = strings.TrimSpace(s)
    return s
}

/* Parses a line of input as a single integer 
   If there are multiple integers on the same line, this only returns
   the first integer.
*/
func (in *InputParser) ParseInt() int {
    s := in.ParseString()
    i,err := strconv.Atoi(s)
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

