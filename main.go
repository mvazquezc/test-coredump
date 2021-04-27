package main

import "fmt"
import "time"

func main() {

    a := 10
    b := 20
    c := a + b
    fmt.Printf("a: %d, b: %d, c: %d\n", a, b ,c)
    i := 10
    for {
        if (i <= 0) {
            panic("Program crashed")
        }
        fmt.Printf("Crashing in %d\n", i)
        time.Sleep(1 * time.Second)
        i--
    }
}
