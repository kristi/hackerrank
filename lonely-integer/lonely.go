package main
import (
    "fmt"
)

func readInt() (x int){
    fmt.Scan(&x)
    return
}

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    size := readInt()
    a := make([]bool, size + 1)
    for i:=0; i<size; i++ {
        x := readInt()
        a[x] = ! a[x]
    }
}
