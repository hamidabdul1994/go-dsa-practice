package main

func divide(dividend int, divisor int) int {
	// Check for divide by zero
	if divisor == 0 {
		return 0
	}

	// Check for overflow
	if dividend == -(1<<31) && divisor == -1 {
		return 1<<31 - 1
	}

	// Determine the sign of the result
	sign := 1
	if (dividend < 0) != (divisor < 0) {
		sign = -1
	}

	// Convert both dividend and divisor to positive numbers
	dividend = abs(dividend)
	divisor = abs(divisor)

	result := 0
	for dividend >= divisor {
		temp := divisor
		multiple := 1
		for dividend >= (temp << 1) {
			temp <<= 1
			multiple <<= 1
		}
		dividend -= temp
		result += multiple
	}

	return result * sign
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
