package main
import (
    "fmt"
    "strings"
)

func readInt() (x int){
    fmt.Scan(&x)
    return
}

func readArray() (a []int){
    s := readInt()
    for i:=0; i<s; i++ {
        a = append(a, readInt())
    }
    return
}

func printArray(a []int) {
    fmt.Println(strings.Trim(fmt.Sprint(a),"[]"))
}

func insertionSort(a []int) {
    L := len(a)
    v := a[L-1]
    for i:=L-2; i>=0; i-- {
        if a[i] > v {
            a[i+1] = a[i]
        } else {
            a[i+1] = v
            break
        }
    }
    if v < a[1] {
        a[0] = v
    }
}

func main() {
    //a := readArray()
    a := []int{1,4,3,5,6,2}
    for i:=1; i<len(a); i++ {
        insertionSort(a[:i+1])
        printArray(a)
    }
}
