//! investment_calculator in GoLang 🧮💸

package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmt float64
	var noOfYears float64 
	expectedReturnRate := 5.5 //default value

	//Getting the user input ⌨️
	fmt.Print("💵 Enter the investment amount: ")
	fmt.Scan(&investmentAmt)
	fmt.Print("📅 Enter the number of years: ")
	fmt.Scan(&noOfYears)
	fmt.Print("💭 Enter the expected return rate (%): ")
	fmt.Scan(&expectedReturnRate)
	

	futureVal := investmentAmt * math.Pow(1+ expectedReturnRate/100, noOfYears)
	futureRealVal := futureVal/ math.Pow(1+ inflationRate/100, noOfYears)

	fmt.Println("\n💰 The future value is: $",futureVal,"✅")
	fmt.Println("📈 The future real value considering inflation is: $",futureRealVal,"☑️")
}

//! Alt. ways- 
//var investmentAmt, noOfYears float64 = 1000,10 // Alt. way with same dataType
//investmentAmt, years, expectedRate := 1000.0, 10.0, 5.5 //Inferred
//CHOICE IS OURS