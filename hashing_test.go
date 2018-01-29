package utils

import "testing"

func TestHashPassword(t *testing.T) {
	password := "something#random"

	passwordHash, err := HashPassword(password)

	if err != nil {
		t.Errorf("Test case failed for HashPassword with argument: %s, error returned: %s", password, err.Error())
	}

	if passwordHash == "" {
		t.Errorf("Test case failed for HashPassword with argument: %s, got: %s", password, passwordHash)
	}

	doesPassMatch := CheckPasswordHash(password, passwordHash)
	isExpectedResult := true

	if doesPassMatch != isExpectedResult {
		t.Errorf("Test case failed for CheckPasswordHarsh with arguments: %s, %s, got: %t, want: %t", password, passwordHash, isExpectedResult, doesPassMatch)
	}

	passwordHash = "something#wrong"

	doesPassMatch = CheckPasswordHash(password, passwordHash)
	isExpectedResult = false

	if doesPassMatch != isExpectedResult {
		t.Errorf("Test case failed for CheckPasswordHarsh with arguments: %s, %s, got: %t, want: %t", password, passwordHash, isExpectedResult, doesPassMatch)
	}
}
