package main

import "fmt"

//CH4,L3
func getMonthlyPrice(tier string) int {
	fmt.Print(tier)
	// ?
	if tier == "basic" {
		return 10000
	} else if tier == "premium" {
		return 15000
	} else if tier == "enterprise" {
		return 50000
	}
	return 0
}


