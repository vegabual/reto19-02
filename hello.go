package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	quantityOfTests := requestQuantity()
	validateAndRunTests(quantityOfTests)
}

func requestQuantity() int {
	var quantity int
	fmt.Scan(&quantity)
	return quantity
}

func validateAndRunTests(quant int) {
	if checkQuantityOfTests(quant) {
		requestTests(quant)
	} else {
		fmt.Println("Invalid quantity of tests")
	}
}

func checkQuantityOfTests(quant int) bool {
	return (quant > 1 && quant <= 1000)
}

func requestTests(quant int) {
	for i := 1; i <= quant; i++ {
		test(i)
	}
}

func test(testNum int) {
	i, j := float64(rand.Intn(150000)), float64(rand.Intn(150000))
	remainder := math.Remainder(i, j)
	fmt.Println(testNum, i, j, remainder)
}
