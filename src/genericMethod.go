package main

type email struct {
	recipient string
	subject   string
}

type receipt struct {
	amount int
}

func (in inbox) Emails() []email {
	result := []email{}
	for _, event := range in.events {
		if e, ok := event.(email); ok {
			result = append(result, e)
		}
	}
	return result
}

func (in inbox) Receipts() []receipt {
	result := []receipt{}
	for _, event := range in.events {
		if r, ok := event.(receipt); ok {
			result = append(result, r)
		}
	}
	return result
}

type inbox struct {
	events []any
}

func Select[T any](in inbox) []T {
	result := []T{}
	for _, event := range in.events {
		if r, ok := event.(T); ok {
			result = append(result, r)
		}
	}
	return result
}
