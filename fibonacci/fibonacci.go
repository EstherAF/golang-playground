package main

// fibonacci is a function that returns
// a function that returns an int.
func Fibonacci() func() int {
	oneBefore := 0
	twoBefore := 0
	return func() int {
		new := oneBefore + twoBefore
		twoBefore = oneBefore
		oneBefore = new

		if oneBefore == 0 && twoBefore == 0 {
			twoBefore = 1
		}
		return new
	}
}
