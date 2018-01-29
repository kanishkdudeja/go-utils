package utils

import "testing"

func TestIsEmailValid(t *testing.T) {
	emailAddress := "kanishk.dudeja@gmail.com"
	isExpectedResult := true

	isValid := IsEmailValid(emailAddress)

	if isValid != true {
		t.Errorf("Test case failed for IsEmailValid with argument as %s, got: %t, want: %t", emailAddress, isValid, isExpectedResult)
	}

	emailAddress = "kanishk.dudeja gmail.com"
	isExpectedResult = false

	isValid = IsEmailValid(emailAddress)

	if isValid != false {
		t.Errorf("Test case failed for IsEmailValid with argument as %s, got: %t, want: %t", emailAddress, isValid, isExpectedResult)
	}

	emailAddress = "kanishk.dudejagmail.com"
	isExpectedResult = false

	isValid = IsEmailValid(emailAddress)

	if isValid != false {
		t.Errorf("Test case failed for IsEmailValid with argument as %s, got: %t, want: %t", emailAddress, isValid, isExpectedResult)
	}

	emailAddress = "kanishk@dudeja@gmail.com"
	isExpectedResult = false

	isValid = IsEmailValid(emailAddress)

	if isValid != false {
		t.Errorf("Test case failed for IsEmailValid with argument as %s, got: %t, want: %t", emailAddress, isValid, isExpectedResult)
	}
}
