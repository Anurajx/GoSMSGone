package main

import (
	"fmt"
)

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	// ?
	customerCost, customerErr := sendSMS(msgToCustomer)
	if customerErr != nil {
		fmt.Printf("could not send mesage: %v",customerErr)
		return 0,customerErr
	}
	spouseCost,spouseErr := sendSMS(msgToSpouse)
	if spouseErr != nil {
		fmt.Printf("could not send your spouse message: %v", spouseErr)
		return 0,spouseErr
	}
	totalCost := spouseCost+customerCost
	return totalCost, nil

}

// don't edit below this line

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}
