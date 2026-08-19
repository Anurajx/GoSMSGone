package main

import "fmt"

func concat(s1 string, s2 string) string {
	return s1 + s2
}

// don't touch below this line

// func main() {
// 	e := email{
// 		isSubscribed: true,
// 		toAddress: "anurajupadhyay6@gmail.com",
// 		body: "netra",
// 	}
// 	to,cost := getExpenseReport(e)
// 	fmt.Printf("the email is for %v and costs %v Rs",to,cost)
// 	// test("Lane,", " happy birthday!")
// 	// test("Zuck,", " hope that Metaverse thing works out")
// 	// test("Go", " is fantastic")

	

	

// }

func main() {
	e := email{
		isSubscribed: true,
		body:         "Whoa there!",
		toAddress:    "soldier@monty.com",
	}

	fmt.Println(getExpenseReport(e))
	finalString:=getSMSErrorString(5, "is this correct? it is")
	fmt.Print(finalString)
}

func test(s1 string, s2 string) {
	fmt.Println(concat(s1, s2))
	testies(true, true)
	testies(false, true)
	fmt.Println("--------")
	fmt.Println(placeOrder("1", 5, 10.00))
	fmt.Println(getMonthlyPrice("basic"))
}



