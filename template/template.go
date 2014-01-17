package main
import (
    "fmt"
    "strings"
    "io"
    "bufio"
    "os"
    "strconv"
)

func solve(a []int) (ans int) {
    return
}

func main() {
    input := NewInputParser(os.Stdin)
    cases := input.ParseInt()

    for c:=0; c<cases; c++ {
        size := input.ParseInt()
        a := input.ParseIntArray(size)

        n := solve(a)
        fmt.Println(n)
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
