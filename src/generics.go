package main

import (
	"errors"
	"fmt"
	"time"
)

func getLast[T any](s []T) T {
	// ?
	if len(s) == 0 {
		var zero T
		return zero
	}

	length := len(s)
	return s[length-1]
}

//constraints

func chargeForLineItem[T lineItem](newItem T, oldItems []T, balance float64) ([]T, float64, error) {
	// ?
	if newItem.GetCost() < balance {
		var genericT []T
		return genericT, 0.0, errors.New("Insfiicient funds")
	}

	oldItems = append(oldItems, newItem)
	newBalance := balance - newItem.GetCost()
	return oldItems, newBalance, nil
}

// don't edit below this line

type lineItem interface {
	GetCost() float64
	GetName() string
}

type subscription struct {
	userEmail string
	startDate time.Time
	interval  string
}

func (s subscription) GetName() string {
	return fmt.Sprintf("%s subscription", s.interval)
}

func (s subscription) GetCost() float64 {
	if s.interval == "monthly" {
		return 25.00
	}
	if s.interval == "yearly" {
		return 250.00
	}
	return 0.0
}

type oneTimeUsagePlan struct {
	userEmail        string
	numEmailsAllowed int
}

func (otup oneTimeUsagePlan) GetName() string {
	return fmt.Sprintf("one time usage plan with %v emails", otup.numEmailsAllowed)
}

func (otup oneTimeUsagePlan) GetCost() float64 {
	const costPerEmail = 0.03
	return float64(otup.numEmailsAllowed) * costPerEmail
}

// GenericStructs
type Envelope[T any] struct {
	recipient string
	payload   T
}

func createEnvelopes[T any](recipients []string, payload T) []Envelope[T] {
	if len(recipients) == 0 {
		return []Envelope[T]{}
	}
	ans := make([]Envelope[T], 0, len(recipients))
	for _, recipient := range recipients {
		ans = append(ans, Envelope[T]{
			recipient: recipient,
			payload:   payload,
		})
	}
	return ans
}
