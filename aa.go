package main

import (
    "fmt"
    "utf8"
)

func main() {
    for i:=1;i<21;i++ {
        for j:=1;j<i;j++ {
            fmt.Printf("A")
        }
        fmt.Printf("\n")
    }


    s:= "asSASA ddd dsjkdsjs dk"
    l:= len([]byte(s))
    total:=utf8.RuneCount([]byte(s))

    fmt.Printf("length %d",l)
    fmt.Printf("total bytes %d",total)

    a:= []byte(s)
    for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
        a[i], a[j] = a[j], a[i]
    }
    fmt.Printf("%s\n", string(a))
}
