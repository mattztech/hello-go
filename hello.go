package main

import "fmt"

var packageScope int
var packageScopeWithVal = 10
var notUsed bool = true

func main() {
	/* Multiple returns in Go - _ ignores the return, as you may not declare a variable which is not used in Go
	Note that package scope variables must use var keyword, and cannot use := operator.
	Also, package level variables do not have to be used
	 */
	bytes, _ := fmt.Println("Hello Go!")
	fmt.Println(bytes)
	fmt.Println(packageScope)
	fmt.Println(packageScopeWithVal)
	fmt.Printf("%T\n", packageScope)
	// Included for completeness only - see docs, this actually prints to stderr
	println("Thing")
}
