package p009

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	div := 1
	for ; x/div >= 10; {
		div *= 10
	}
	var left int
	var right int
	for ; x > 0; {
		right = x % 10
		left = x / div
		if left != right {
			return false
		}
		x = (x % div) / 10
		div /= 100
	}
	return true
}
