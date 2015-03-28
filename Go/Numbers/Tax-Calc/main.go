package main

import "fmt"
import "os"
import "strconv"
import "math"
import "errors"

const FLOAT_BIT_SIZE = 64

/**
 * this function courtesy of https://gist.github.com/DavidVaini/10308388
 */
func Round(val float64) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(2))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= .5 {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

func GetPercentage(arg string) (amt float64, err error) {

	taxAmt := 0.0
	err = errors.New("")

	if string(arg[len(arg)-1]) == "%" {
		taxAmt, err = strconv.ParseFloat(string(arg[:len(arg)-1]), FLOAT_BIT_SIZE)
	} else {
		taxAmt, err = strconv.ParseFloat(arg, FLOAT_BIT_SIZE)
	}

	return taxAmt, err
}

func GetDollarAmt(arg string) (amt float64, err error) {

	dollarAmt := 0.0
	err = errors.New("")

	if string(arg[0]) == "$" {
		dollarAmt, err = strconv.ParseFloat(string(arg[1:]), FLOAT_BIT_SIZE)
	} else {
		dollarAmt, err = strconv.ParseFloat(arg, FLOAT_BIT_SIZE)
	}

	return dollarAmt, err
}

func Compute(dollarAmt float64, taxPercent float64) (totalTax float64, total float64) {
	taxPercent /= 100.00
	totalTax = dollarAmt * taxPercent
	total = totalTax + dollarAmt
	return Round(totalTax), Round(total)
}

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage:\narg0 - Dollar amount\narg1 - tax percentage\nResult - Tax amount and total cost")
	}

	dollarAmt, err1 := GetDollarAmt(os.Args[1])
	if err1 != nil {
		fmt.Println("Error converting dollar amount into float64")
		os.Exit(1)
	}

	taxAmt, err2 := GetPercentage(os.Args[2])
	if err2 != nil {
		fmt.Println("Error converting tax amount into float64")
		os.Exit(1)
	}

	totalTax, total := Compute(dollarAmt, taxAmt)

	fmt.Println("Tax amount is $", totalTax)
	fmt.Println("total is $", total)
}
