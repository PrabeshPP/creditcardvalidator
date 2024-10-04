package main

import "fmt"


//API that will take the credit card number from the user and it will check whether the credit card number is valid or not
//If the provided payload is in-valid then the API will return Invalid error message in a JSON Format
func main() {
	creditCardNumber := []int{7, 9, 9, 2, 7, 3, 9, 8, 7, 1, 3}
	fmt.Println(Luhn(creditCardNumber))
}
