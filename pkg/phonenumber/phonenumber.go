package phonenumber

import (
	"fmt"
	"strconv"
)

func IsValid(phoneNumber string) bool {
	// TODO - regular expression
	fmt.Println(len(phoneNumber))
	if len(phoneNumber) != 11 {
		return false
	}
	fmt.Println("phoneNumber:", phoneNumber[:2])
	if phoneNumber[:2] != "09" {
		return false
	}
	fmt.Println("phoneNumber:", phoneNumber[2:])
	if _, err := strconv.Atoi(phoneNumber[2:]); err != nil {
		return false
	}
	return true
}
