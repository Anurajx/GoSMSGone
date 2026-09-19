package main

import "fmt"

type Messages struct {
	Recipient string
	Text      string
	Success   bool
}

func getMessageText(m Messages) string {
	return fmt.Sprintf(`
To: %v
Message: %v
`, m.Recipient, m.Text)
}
