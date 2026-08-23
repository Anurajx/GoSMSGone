package main

func bulkSend(numMessages int) float64 {
	totalCost := 0.0
	for i := 0; i < numMessages; i++ {
		extraFess := float64(i) * 0.01
		totalCost += 1 + extraFess
	}
	return float64(totalCost)
	// ?
}

func maxMessages(thresh int) int {
	// ?
	cost := 0
	for i := 0; ; i++ {

		if cost+100+i > thresh {
			return i
		} else {
			cost += 100 + i
		}

	}
}


func getMaxMessagesToSend(costMultiplier float64, budgetInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 1
	balance := float64(budgetInPennies) - actualCostInPennies
	for balance>0{
		actualCostInPennies *= costMultiplier
		balance -= actualCostInPennies
		maxMessagesToSend++
	}
	if balance < 0 {
		maxMessagesToSend--
	}
	return maxMessagesToSend
}
