package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var xDigits []int
	var currDigit int
	for ; x != 0; x /= 10 {
		currDigit = x % 10
		xDigits = append(xDigits, currDigit)
	}

	n := len(xDigits)
	for i := range xDigits {
		if xDigits[i] != xDigits[n-1-i] {
			return false
		}
	}

	return true
}

func main() {
	x := 121
	fmt.Printf("isPalindrome(%v) = %v\n", x, isPalindrome(x))

	x = -121
	fmt.Printf("isPalindrome(%v) = %v\n", x, isPalindrome(x))

	x = 10
	fmt.Printf("isPalindrome(%v) = %v\n", x, isPalindrome(x))

	x = 12345
	fmt.Printf("isPalindrome(%v) = %v\n", x, isPalindrome(x))

	x = 1239321
	fmt.Printf("isPalindrome(%v) = %v\n", x, isPalindrome(x))
}
