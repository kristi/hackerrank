package main
import (
    "fmt"
    "strings"
    "io"
    "bufio"
    "os"
    "strconv"
)

func solve(k int, a []int) (ans int) {
    m := make(map[int]bool, len(a))
    for _,v := range a {
        m[v] = true
    }
    for x := range m {
        if _, ok := m[x+k]; ok {
            ans++
            continue
        }
    }
    return
}

func main() {
    input := NewInputParser(os.Stdin)
    var n,k int

    input.Scan(&n, &k)
    a := input.ParseIntArray(n)

    ans := solve(k, a)
    fmt.Println(ans)
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
