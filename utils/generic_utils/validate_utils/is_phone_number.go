package validate_utils

import "regexp"

var phoneNumberRegexString = "^1[0-9]{10}$"
var phoneNumberRegex = regexp.MustCompile(phoneNumberRegexString)

func IsPhoneNumber(phoneNumber string) bool {
	return phoneNumberRegex.MatchString(phoneNumber)
}
