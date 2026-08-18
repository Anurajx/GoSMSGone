package main

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

func (d directMessage) importance() int {
	if d.isUrgent {
		return 50
	}
	return d.priorityLevel
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

func (g groupMessage) importance() int {
	return g.priorityLevel
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (s systemAlert) importance() int {
	return 100
}

// ?

func processNotification(n notification) (string, int) {
	switch new := n.(type) {
	case systemAlert:
		return new.alertCode, new.importance()
	case directMessage:
		return new.senderUsername, new.importance()
	case groupMessage:
		return new.groupName, new.importance()
	default:
		return "", 0
	}
	// ?
}
