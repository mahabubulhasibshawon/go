
package main

import "fmt"

func main() {
    var n, a int
    fmt.Scan(&n)
    for i:=0; i< n; i++ {
        fmt.Scan(&a)
        fmt.Println(a-1)
    }
}