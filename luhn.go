package main

/*
This is a Luhn Algorithm implemented in GoLang
Author: Prabesh Bista
*/
func Luhn(creditNumber []int) bool {

	// Double the every alternate digit in the array
	//update the index of which the digit has been doubled
	//Currently the time complexity for this algorithm is O(n) and space complexity is O(1)

	for i := len(creditNumber) - 2; i > 0; i -= 2 {
		currentSum := creditNumber[i] * 2

		if currentSum > 9 {
			fDigit := currentSum / 10
			sDigit := currentSum % 10
			creditNumber[i] = fDigit + sDigit

		} else {
			creditNumber[i] = currentSum
		}

	}

	var totalSum int = 0

	for _, num := range creditNumber {
		totalSum += num
	}

	return totalSum%10 == 0
}
