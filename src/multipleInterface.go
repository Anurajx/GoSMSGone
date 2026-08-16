package main

import "fmt"


func main() {
	e := email{
		isSubscribed: true,
		body:         "Whoa there!",
		toAddress:    "soldier@monty.com",
	}

	fmt.Println(getExpenseReport(e))
}


func getExpenseReport(e expense) (string, float64) {
	e1, ok := e.(email)
	if ok {
		return e1.toAddress, float64(e1.cost())
	}
	e2, ok1 := e.(sms)
	if ok1{
		return e2.toPhoneNumber, e2.cost()
	}
	return "",0.0
	// ?
}

func (e email) cost() float64 {
	charCost := 0
	if e.isSubscribed  {
		charCost = 2
	} else {
		charCost = 5
	}
	
	print(charCost)
	return float64(len(e.body) * charCost)
}

func (e email) format() string {
	if e.isSubscribed {
		return fmt.Sprintf("'%v' | SUBSCRIBED ",e.body)
	}
	return fmt.Sprintf("'%v' | NOT SUBSCRIBED ",e.body)
	// ?
}



func (sm sms) cost() float64 {
	if !sm.isSubscribed {
		return float64(len(sm.body)) * .1
	}
	return float64(len(sm.body)) * .03
}

// func (inv inva) cost() float64 {
// 	return 0.0
// }

type expense interface {
	cost() float64
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
	toAddress	string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

