package main

import (
	"fmt"
	"math"
)

// returns the amount required for you to retire safely with a given salary.
func calculateAmountRequiredForDesiredSalary(s float64) float64 {
	// the equation used for this is based on the 4% rule. this rule says that if you want
	// to retire with specific yearly salary safely, you need to multiple that number by 25
	// and you would safely be able to withdraw 4% from the account per year.
	return s * 25
}

// returns the amount of years until the total amount invested would equal the desired amount.
func calculateYearsUntilAmount(desiredAmount float64, monthlyContribution float64, yearlyROR float64) float64 {
	monthlyROR := yearlyROR / 100 / 12

	numMonths := math.Log((desiredAmount*monthlyROR)/monthlyContribution+1) / math.Log(1+monthlyROR)

	numYears := numMonths / 12

	return numYears
}

// prompts the user with the prompt and writes the response to the pointer.
func promptUser(prompt string, pointer *float64) {
	fmt.Print(prompt)
	_, err := fmt.Scan(pointer)
	if err != nil {
		panic(err)
	}
}

func main() {
	var desiredSalary float64
	promptUser("Enter your desired retirement salary: ", &desiredSalary)

	var monthlyContribution float64
	promptUser("Enter accurate monthly contribution amount: ", &monthlyContribution)

	var expectedROR float64
	promptUser("Enter you expected yearly rate of return (a percentage number between 1 and 100, a commonly used value is anywhere between 7 - 10):", &expectedROR)

	var currentAge float64
	promptUser("Enter your current age: ", &currentAge)

	amountRequiredForDesiredSalary := calculateAmountRequiredForDesiredSalary(desiredSalary)

	yearsUntilRetirement := calculateYearsUntilAmount(amountRequiredForDesiredSalary, monthlyContribution, expectedROR)

	retirementAge := currentAge + yearsUntilRetirement

	fmt.Printf("With a monthly contribution of $%.2f, an expected yearly rate of return of %.2f%%, and a desired salary of $%.2f, you will need to work for another %.2f years.\n", monthlyContribution, expectedROR, desiredSalary, yearsUntilRetirement)
	fmt.Printf("This means you could retire at age %.2f.\n", retirementAge)
}
