package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

//API that will take the credit card number from the user and it will check whether the credit card number is valid or not
//If the provided payload is in-valid then the API will return Invalid error message in a JSON Format

type CardNumber struct{
	CreditNum string `json:"creditnum`
}

func checkCreditNumber(w http.ResponseWriter, r *http.Request) {
	var creditNum CardNumber
	err := json.NewDecoder(r.Body).Decode(&creditNum)

	if err!= nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		return 
	}

	var cNumber []int

	for _,cNum := range(creditNum.CreditNum){
		val := int(cNum) - 48
		cNumber = append(cNumber, val)
	}
	
	isValid := Luhn(cNumber)

	//returns whether the given number is valid or not
	json.NewEncoder(w).Encode(isValid)
}

func main() {
	http.HandleFunc("/", checkCreditNumber)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
