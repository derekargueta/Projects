package main

import "os"
import "fmt"
import "strconv"

/**
 * FOR USAGE SEE README.md
 */

func RunCollatz(n int, count int) (res int) {

    if n == 1 {
        return count
    }

    count++
    if n % 2 == 0 {
        return RunCollatz(n / 2, count)
    } else {
        return RunCollatz(n * 3 + 1, count)
    }
}

func CollatzInit(n int) (res int) {
    return RunCollatz(n, 0)
}

func main() {

    if len(os.Args) == 1 {
        fmt.Println("Please enter a number")
        return
    }

    value, err := strconv.ParseInt(os.Args[1], 10, 64)
    if err != nil {
        fmt.Println("Error reading integer")
        return
    }

    if value <= 1 {
        fmt.Println("Please enter a valid number")
    }

    fmt.Println(CollatzInit(int(value)))
}