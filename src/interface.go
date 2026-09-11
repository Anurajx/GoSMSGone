package main

import (
	"fmt"
	"strings"
	"time"
)






type plainText struct {
	message string
}
func  (p plainText) format() (string) {
	return p.message
}

type bold struct {
	message string
}
func  (b bold) format() (string) {
	return fmt.Sprintf("**%v**",b.message)
}

type code struct {
	message string
}

func  (c code) format() (string) {
	return fmt.Sprintf("`%v`",c.message)
}

type formatter interface {
	format() string
}


func getExpenseReport(e expense) (string, float64) {
	switch v := e.(type){
	case email:
		return v.toAddress, float64(e.cost())
	case sms:
		return v.toPhoneNumber, float64(e.cost())
	default:
		return "",0.0
	}
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



type email struct {
	isSubscribed bool
	body         string
	toAddress	string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
	
	id      string
	content string
	tags    []string
}



func sendMessage(msg message) (string, int) {
	messageContent := msg.getMessage()
	messageLen := len(messageContent)*3
	return messageContent, messageLen

	// ?

}

type message interface {
	
	getMessage() string


	// ?
}

// don't edit below this line

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}



//Tag Message
func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	// ?
	// messagesWithTags := []sms{}
	for i, msg := range messages {
		msg.tags = tagger(msg)
		messages[i] = msg

	}
	return messages
}

func tagger(msg sms) []string {
	tags := []string{}
	msg.content = strings.ToLower(msg.content)

	if strings.Contains(msg.content, "urgent") {
		tags = append(tags, "Urgent")
	}
	if strings.Contains(msg.content, "sale") {
		tags = append(tags, "Promo")
	}
	if strings.Contains(msg.content, "important") {
		tags = append(tags, "Important")
	}
	

	return tags
	// ?
}
