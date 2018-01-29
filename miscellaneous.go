package utils

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

// GetCurrentUTCTime functions returns the UTC time in the YYYY-MM-DD HH:MM:SS format
func GetCurrentUTCTime() string {
	t := time.Now().UTC()
	return t.Format("2006-01-02 15:04:05")
}

// GenerateSlug function can be used to generate a pretty-URL string from a UNICODE string
// It replaces spaces with dashes, removes all non-ASCII characters
// and then trims the string for any starting/ending dashes
func GenerateSlug(text string) string {
	var re = regexp.MustCompile("[^a-z0-9-]+")

	text = strings.Replace(text, " ", "-", -1)

	replacedString := re.ReplaceAllString(strings.ToLower(text), "")
	return strings.Trim(replacedString, "-")
}

// SplitEmailAddress takes an email address as argument
// and returns the recepient name (part before the @ sign)
// and the domain name (part after the @ sign)
func SplitEmailAddress(emailAddress string) (string, string, error) {
	components := strings.Split(emailAddress, "@")

	if len(components) < 2 {
		return "", "", errors.New("Invalid Email Address. Does not contain @ sign")
	}

	if len(components) > 2 {
		return "", "", errors.New("Invalid Email Address. Contains >1 @ signs")
	}

	return components[0], components[1], nil
}
