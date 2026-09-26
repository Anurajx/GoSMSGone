package main

import (
	"fmt"
	"time"
)

func sendEmail(message string) {
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
	}()
	fmt.Printf("Email sent: '%s'\n", message)
}

// Don't touch below this line

// type email struct {
// 	body string
// 	date time.Time
// }

func checkEmailAge(emails [3]email) [3]bool {
	isOldChan := make(chan bool)
	go sendIsOld(isOldChan, emails)

	isOld := [3]bool{}
	isOld[0] = <-isOldChan
	isOld[1] = <-isOldChan
	isOld[2] = <-isOldChan
	return isOld
}

// don't touch below this line

func sendIsOld(isOldChan chan<- bool, emails [3]email) {
	for _, e := range emails {
		
		if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
			
			isOldChan <- true
			continue
		}
		isOldChan <- false
	}
}

//DB connection establised

func waitForDBs(numDBs int, dbChan chan struct{}) {
	for i:=0;i < numDBs;i++ {
		<-dbChan
	} 
}

// don't touch below this line

func getDBsChannel(numDBs int) (chan struct{}, *int) {
	count := 0
	ch := make(chan struct{})

	go func() {
		for i := 0; i < numDBs; i++ {
			ch <- struct{}{}
			fmt.Printf("Database %v is online\n", i+1)
			count++
		}
	}()

	return ch, &count
}


//Closing channels ------------------
func countReports(numSentCh chan int) int {
	iter := 0
	for {
		_, ok := <- numSentCh
		if !ok {
			break
		}
		iter++
	}
	return iter
	// ?
}

// don't touch below this line

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}

//Concurrent Fibonacci-----------

func concurrentFib(n int) []int {
	// ?
	ch := make(chan int)
	result := []int{}
	go fibonacci(n,ch)
	for val := range ch {
		result = append(result, val)

	}
	return result
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
