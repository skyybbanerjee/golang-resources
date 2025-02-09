//todo: Build A Profit Calculator

//1. Ask for revenue, expenses and tax-rate
//2. Calculate EARNINGS BEFORE TAX (EBT) and EARNINGS AFTER TAX (profit)
//3. Calculate ratio (EBT/profit)
//4. Output EBT, profit and the ratio

package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("💵 Enter revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("🛒 Enter expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("📈 Enter tax rate (%): ")
	fmt.Scan(&taxRate)

	

	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - taxRate/100)
	ratio := earningsBeforeTax / profit

	//!Same results //('Sprintf' stores strings, for later use)
	// formattedEBT := fmt.Sprintf("\nEBT: $%.2f",earningsBeforeTax)
	// formattedProfit := fmt.Sprintf("\nProfit: $%.2f", profit)
	// formattedRatio := fmt.Sprintf("\nRatio (EBT/profit): %.2f", ratio)
	// fmt.Print(formattedEBT, formattedProfit,formattedRatio)

	fmt.Printf("\nEBT: $%.2f",earningsBeforeTax)
	fmt.Printf("\nProfit: $%.2f", profit)
	fmt.Printf("\nRatio (EBT/profit): $%.2f", ratio)

}