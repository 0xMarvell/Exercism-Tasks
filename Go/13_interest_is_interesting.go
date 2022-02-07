/* Instructions
In this exercise you'll be working with savings accounts. Each year, the balance of your savings account is updated based on its interest rate. The interest rate your bank gives you depends on the amount of money in your account (its balance):

3.213% for a negative balance (balance gets more negative).
0.5% for a positive balance less than 1000 dollars.
1.621% for a positive balance greater or equal than 1000 dollars and less than 5000 dollars.
2.475% for a positive balance greater or equal than 5000 dollars.
You have four tasks, each of which will deal your balance and its interest rate.

1. Calculate the interest rate
Implement the InterestRate() function to calculate the interest rate based on the specified balance:

InterestRate(200.75)
// Output: 0.5
Note that the value returned is a float32.

2. Calculate the interest
Implement the Interest() function to calculate the interest based on the specified balance:

Interest(200.75)
// Output: 1.003750
Note that the value returned is a float64.

3. Calculate the annual balance update
Implement the AnnualBalanceUpdate() function to calculate the annual balance update, taking into account the interest rate:

AnnualBalanceUpdate(200.75)
// Output: 201.75375
Note that the value returned is a float64.

4. Calculate the years before reaching the desired balance
Implement the YearsBeforeDesiredBalance() function to calculate the minimum number of years required to reach the desired balance, taking into account that each year, interest is added to the balance. This means that the balance after one year is: start balance + interest for start balance. The balance after the second year is: balance after one year + interest for balance after one year. And so on, until the current year's balance is greater than or equal to the target balance.

balance := 200.75
targetBalance := 214.88
YearsBeforeDesiredBalance(balance, targetBalance)
// Output: 14
Note that the value returned is an int. */

package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	panic("Please implement the InterestRate function")
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	panic("Please implement the Interest function")
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	panic("Please implement the AnnualBalanceUpdate function")
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	panic("Please implement the YearsBeforeDesiredBalance function")
}
