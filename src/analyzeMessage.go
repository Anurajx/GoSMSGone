package main

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}



// don't touch above this line


func analyzeMessage(a *Analytics, m Messages){
	a.MessagesTotal++

	if m.Success {
		a.MessagesSucceeded++
	} else {
		a.MessagesFailed++
	}
}
// ?

