package main
import (
    "fmt"
    "strings"
    "bufio"
    "os"
    "strconv"
)

func printArray(a []int) {
    fmt.Println(strings.Trim(fmt.Sprint(a),"[]"))
}

func atoi(s string) (i int) {
    i,err := strconv.Atoi(s)
    if err != nil {
        panic(err)
    }
    return
}

func strToArray(s string) (a []int) {
    for _,w := range strings.Split(s, " ") {
        a = append(a, atoi(w))
    }
    return
}

func countLessThanX(a []int, x int) (count int){
    for _,v := range a {
        if x < v {
            count++
        }
    }
    return
}

func countShifts(a []int) (count int) {
    for i:=1; i<len(a); i++ {
        count += countLessThanX(a[:i], a[i])
    }
    return
}

func main() {
    scanner := bufio.NewReader(os.Stdin)
    s,_ := scanner.ReadString('\n')
    s = strings.TrimSpace(s)
    cases := atoi(s)

    for c:=cases; c>0; c-- {
        s,_ = scanner.ReadString('\n')
        s = strings.TrimSpace(s)
        size := atoi(s)

        s,_ = scanner.ReadString('\n')
        s = strings.TrimSpace(s)
        a := strToArray(s)

        n := countShifts(a)
        fmt.Println(n)
    }
}
