package utils

import (
	"net"
	"testing"
)

func TestIsPrivateSubnet(t *testing.T) {
	ip := net.ParseIP("10.10.10.10")
	expectedResult := true

	result := isPrivateSubnet(ip)

	if result != expectedResult {
		t.Errorf("Test case failed for IsPrivateSubnet with argument as %s, got: %s, want: %s", ip.String(), result, expectedResult)
	}

	ip = net.ParseIP("200.10.10.10")
	expectedResult = false

	result = isPrivateSubnet(ip)

	if result != expectedResult {
		t.Errorf("Test case failed for IsPrivateSubnet with argument as %s, got: %s, want: %s", ip.String(), result, expectedResult)
	}
}
