package main

import "os"
import "fmt"
import "strconv"

func fib(value int64) (res []int64) {
    var list = make([]int64, value)
    count := int64(2)
    numOne := 0
    numTwo := 1
    list[0] = 0
    list[1] = 1
    fmt.Print(" 0 1")
    for count < value {
        result := numOne + numTwo
        numOne = numTwo
        numTwo = result
        list[count] = int64(result)
        count++
        fmt.Print(" ", result)
    }

    return list
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

    if value < 2 {
        fmt.Println("Please enter a valid number above 2")
    }

    fmt.Println("\n", fib(value))
}