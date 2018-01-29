package utils

import "testing"

func TestSplitEmailAddress(t *testing.T) {
	emailAddress := "kanishk.dudeja@gmail.com"
	firstExpectedArg := "kanishk.dudeja"
	secondExpectedArg := "gmail.com"

	firstArg, secondArg, err := SplitEmailAddress(emailAddress)

	if err != nil {
		t.Errorf("Test case failed for SplitEmailAddress with argument as %s, got error: %s", emailAddress, err.Error())
	}

	if (firstExpectedArg != firstArg) || (secondExpectedArg != secondArg) {
		t.Errorf("Test case failed for SplitEmailAddress with argument as %s, got: %s and %s, want: %s and %s", emailAddress, firstArg, secondArg, firstExpectedArg, secondExpectedArg)
	}

	emailAddress = "kanishk.dudeja gmail.com"

	_, _, err = SplitEmailAddress(emailAddress)

	if err == nil {
		t.Errorf("Test case failed for SplitEmailAddress with argument as %s, expected error, did not get error", emailAddress)
	}

	emailAddress = "kanishk@dudeja@gmail.com"

	_, _, err = SplitEmailAddress(emailAddress)

	if err == nil {
		t.Errorf("Test case failed for SplitEmailAddress with argument as %s, expected error, did not get error", emailAddress)
	}
}

func TestGenerateSlug(t *testing.T) {
	text := "  攻仏ゃで田大ちのド質供ハテ協属 *generated #sample-slug  "
	expectedSlug := "generated-sample-slug"

	slug := GenerateSlug(text)

	if slug != expectedSlug {
		t.Errorf("Test case failed for GenerateSlug with argument as %s, got: %s, want: %s", text, slug, expectedSlug)
	}
}
