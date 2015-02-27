package main

import "os"
import "fmt"
import "bufio"
import "strconv"
import "math"

/**
 * this function courtesy of https://gist.github.com/DavidVaini/10308388
 */
func Round(val float64, roundOn float64, places int ) (newVal float64) {
    var round float64
    pow := math.Pow(10, float64(places))
    digit := pow * val
    _, div := math.Modf(digit)
    if div >= roundOn {
        round = math.Ceil(digit)
    } else {
        round = math.Floor(digit)
    }
    newVal = round / pow
    return
}

/**
 * Function to read in a $###.## format. Returns a decimal rounded
 * to 2 places
 */
func GetMoney()(money float64) {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    token := scanner.Text()
    flo, err := strconv.ParseFloat(token[1:len(token)], 32)
    if err == nil {
        return Round(flo, .5, 2)
    } else {
        fmt.Println("Try again!")
        return GetMoney()
    }
}

/**
 * Takes in the name of a money value, the total to break down, and the amount
 * to be broken by i.e. (twenties, 100.00, 20). Counts how many amt's fit in
 * total
 */
func Reducer(name string, total float64, amt float64) (/*num uint8,*/ remainder float64) {
    if amt <= 0 || amt > total {
        return total
    }

    count := 0
    for total >= amt {
        total -= amt
        count++
    }
    fmt.Println(count, name)
    return total
}

func main() {

    // grab the total
    fmt.Println("Enter the total cost")
    total := GetMoney()

    fmt.Println("Enter the amount paid")
    paid := GetMoney()

    diff := paid - total
    fmt.Printf("Difference: $%.2f\n", diff)

    diff = Reducer("twenties", diff, 20)
    diff = Reducer("tens", diff, 10)
    diff = Reducer("fives", diff, 5)
    diff = Reducer("ones", diff, 1)
    diff = Reducer("quarters", diff, .25)
    diff = Reducer("dimes", diff, .10)
    diff = Reducer("nickels", diff, .05)
    diff = Reducer("pennies", diff, .01)
}