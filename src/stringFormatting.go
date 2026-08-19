package main

import (
	"fmt"
)

func getSMSErrorString(cost float64, recipient string) string {
	// ?
	return fmt.Sprintf("SMS that costs $%0.2f to be sent to '%v' cannot be sent",cost ,recipient)
}

// func main(){
// 	getSMSErrorString(5, "123456789")
// }