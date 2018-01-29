package utils

import "testing"

func TestEncryption(t *testing.T) {
	plainText := "Some Random Text Here"
	key := "16digitkeyhere:)"

	cipherText, err := EncryptString(plainText, key)

	if err != nil {
		t.Errorf("Test case failed for EncryptString with arguments: %s, %s, error returned: %s", plainText, key, err.Error())
	}

	if cipherText == "" {
		t.Errorf("Test case failed for EncryptString with argument: %s, %s, got: %s", plainText, key, cipherText)
	}

	plainTextReturned, err := DecryptString(cipherText, key)

	if err != nil {
		t.Errorf("Test case failed for Decrypt String with arguments: %s, %s, error returned: %s", cipherText, key, err.Error())
	}

	if plainTextReturned == "" {
		t.Errorf("Test case failed for Decrypt String with argument: %s, %s, got: %s", cipherText, key, plainTextReturned)
	}

	if plainText != plainTextReturned {
		t.Errorf("Test case failed for EncryptionDecryption with arguments: %s, %s, got: %s, want: %s", cipherText, key, plainTextReturned, plainText)
	}
}
