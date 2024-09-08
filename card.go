package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type card_num struct {
	Number string `json:"number"`
}

func contains(s string, target rune) bool {
	for _, val := range s {
		if target == val {
			return true
		}
	}
	return false
}

func removeSpaces(cardNum string) string {
	splitNum := strings.Split(cardNum, " ")
	return strings.Join(splitNum, "")
}

func returnCard(r *http.Request) card_num {
	var newCard card_num

	dec := json.NewDecoder(r.Body)

	err := dec.Decode(&newCard)

	if contains(newCard.Number, ' ') {
		newCard.Number = removeSpaces(newCard.Number)
	}

	if err != nil {
		return card_num{}
	}

	return newCard
}

// The bool is whether its a providers name or if it's an industry (is industry)
func providerCheck(potNum string) (string, bool) {
	firstNum := potNum[0]

	switch firstNum {
	case '1':
		return "Airlines and Financial Services", true
	case '2', '5':
		return "Mastercard or Airlines", false
	case '3':
		return "American Express", false
	case '4':
		return "Visa", false
	case '6':
		return "Discover", false
	case '7':
		return "Petroleum", true
	case '8':
		return "Healthcare and Communications", true
	case '9':
		return "Government and Other", true
	}

	return "", false
}

func luhnCheck(potNum string) (int, int) {
	var total int

	intCardNum, err := strconv.Atoi(potNum)

	if err != nil {
		return -1, -1
	}

	checkDigit := intCardNum % 10

	//Ensures the algorithm works with odd and even numbers
	parity := len(potNum) % 2

	// len(potNum) - 2 drops the check digit, just the "payload"
	for i := len(potNum) - 2; i >= 0; i-- {
		num := int(potNum[i] - '0') //'0' cancels out the extra stuff added by the rune

		var result int

		if i%2 == parity {
			result = num * 2
		} else {
			result = num
		}

		//Technically add the digits if greater than 10 BUT can just do 1 (no 2 numbers under 10 sum to greater than 20) + the last number (sum % 10)

		if result > 9 {
			total += 1 + result%10
		} else {
			total += result
		}
	}

	return total, checkDigit
}

func returnCDs(card card_num) (int, int) {
	total, checkDigit := luhnCheck(card.Number)

	realCD := 10 - total%10

	return realCD, checkDigit
}
