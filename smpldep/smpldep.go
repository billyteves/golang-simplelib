package smpldep

import (
    "fmt"
)

type alertCounter int

// NewAlertCounter creates and returns objects of
// the unexported type alertCounter.
func NewAlertCounter(value int) alertCounter {
	return alertCounter(value)
}
