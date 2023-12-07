package main

import (
	"fmt"
	"regexp"
)

func validateCommand(msg string) bool {

	msg1 := msg
	regex1 := regexp.MustCompile(`^/note\s(.+)\|(.+)\|(\d+)\|(.+)$`)
	valid1 := regex1.MatchString(msg1)
	fmt.Println("valid1", msg1, valid1)

	msg2 := msg
	regex2 := regexp.MustCompile(`^/note\s(.+)\|(.+)\|(\d+)$`)
	valid2 := regex2.MatchString(msg2)
	fmt.Println("valid2", msg2, valid2)
	return valid1 || valid2
}
