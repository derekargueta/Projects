package main

import "os"
import "fmt"
import "strconv"

func FibSeq(value int64) (res []int64) {
    
    if value == 1 { return []int64{0} }

    var list = make([]int64, value)
    list[0] = 0
    list[1] = 1
    if value == 2 { return list }
    
    for count := int64(2); count < value; count++ {
        list[count] = int64(list[count-2] + list[count-1])
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

    if value < 1 {
        fmt.Println("Please enter a valid number")
    }

    fmt.Println("\n", FibSeq(value))
}