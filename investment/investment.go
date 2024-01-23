package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64 = 10

	fmt.Print("Please enter investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Please enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years) 
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Printf("Future value %.0f \nFuture real value %v", futureValue, futureRealValue)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, frv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years) 
	frv = fv /math.Pow(1+inflationRate/100, years)

	return fv, frv
}