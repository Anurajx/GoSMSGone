package main

import "fmt"

func (e email) cost() int {
	charCost := 0
	if e.isSubscribed  {
		charCost = 2
	} else {
		charCost = 5
	}
	
	print(charCost)
	return len(e.body) * charCost
}

func (e email) format() string {
	if e.isSubscribed {
		return fmt.Sprintf("'%v' | SUBSCRIBED ",e.body)
	}
	return fmt.Sprintf("'%v' | NOT SUBSCRIBED ",e.body)
	// ?
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}
