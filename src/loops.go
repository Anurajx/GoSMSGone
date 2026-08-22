package main

func bulkSend(numMessages int) float64 {
	totalCost := 0.0
	for i := 1; i <= numMessages; i++ {
		extraFess := float64(i-1) * 0.01
		totalCost += 1 + extraFess
	}
	return float64(totalCost)
	// ?
}
