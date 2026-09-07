package main

func getMessageWithTries(primary, secondary, tertiary string) ([3]string, [3]int) {
	message := [3]string{primary, secondary, tertiary}
	cost := 0
	tries := [3]int{0, 0, 0}
	for i := 0; i< len(message); i++ {
		cost += len(message[i])
		tries[i]= cost
	}
	return message, tries
}