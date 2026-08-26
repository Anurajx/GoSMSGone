package main

func countConnections(groupSize int) int {
	numConnections := 0
	for i:=1;i< groupSize;i++{
		numConnections +=i
	}
	return numConnections
	// ?
}
