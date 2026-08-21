package main

import (
	"errors"
)

func validateStatus(status string) error {
	if status == "" {
		return errors.New("Status message cannot be empty")
	}
	if len(status) >= 140 {
		return errors.New("length of status cant be more than 140 chars")
	}
	return nil
	// ?
}
