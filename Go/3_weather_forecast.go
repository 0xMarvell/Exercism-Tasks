/* Instructions
Being hired by a big weather forecast company, your job is to maintain their code base. Scrolling through the code, you find it hard to follow what the code is actually doing. Good thing you know just the right thing to do!

1. Document package weather
Write a package comment for package weather that describes its contents. The package comment should introduce the package and provide information relevant to the package as a whole.

2. Document the CurrentCondition variable
Write a comment for the package variable CurrentCondition.

This should tell any user of the package what information the variable stores, and what they can do with it.

3. Document the CurrentLocation variable
Just like the previous step, write a comment for CurrentLocation.

4. Document the Forecast() function
Write a function comment for Forecast(). Your comments must describe what the function does, but not how it does it. */

// Package weather provides tools to return the weather forecast for any location it is provided with.
package weather

// CurrentCondition represents the current weather condition for a specific provided location.
var CurrentCondition string

// CurrentLocation represents the porvided location to read the weather forecast from.
var CurrentLocation string

//Forecast takes two string arguments "city" and "condition" and returns a string value that represents the weather forecast for a specific location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
