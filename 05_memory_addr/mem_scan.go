package main

import "fmt"

const conversion float64 = 1.60000

func main() {
	var meter float64
	fmt.Println("Enter distance in KM")
	fmt.Scan(&meter)
	result := meter * conversion
	fmt.Println(result)

}
