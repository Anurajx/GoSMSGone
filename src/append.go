package main

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	// ?
	//var dayCost []float64
	// dayCost := []float64{}
	dayCost := make([]float64,0)
	for _, i := range costs {
		if i.day == day {
			dayCost = append(dayCost, i.value)
		}
	}
	return dayCost
}
