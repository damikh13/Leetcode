Given int x, return if it's a palindrome or not.

- Get digits into a slice
- Go through that slice and check if `digit[current idx] == digit[n - current idx]`

```go
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
```
