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
		t.Errorf("Test case failed for IsPrivateSubnet with argument as %s, got: %t, want: %t", ip.String(), result, expectedResult)
	}

	ip = net.ParseIP("200.10.10.10")
	expectedResult = false

	result = isPrivateSubnet(ip)

	if result != expectedResult {
		t.Errorf("Test case failed for IsPrivateSubnet with argument as %s, got: %t, want: %t", ip.String(), result, expectedResult)
	}

	ip = net.ParseIP("fd8f:a840:e5ab:95a3:1234:1234:1234:1234")
	expectedResult = true

	result = isPrivateSubnet(ip)

	if result != expectedResult {
		t.Errorf("Test case failed for IsPrivateSubnet with argument as %s, got: %t, want: %t", ip.String(), result, expectedResult)
	}

	ip = net.ParseIP("2001:0db8:85a3:0000:0000:8a2e:0370:7334")
	expectedResult = false

	result = isPrivateSubnet(ip)

	if result != expectedResult {
		t.Errorf("Test case failed for IsPrivateSubnet with argument as %s, got: %t, want: %t", ip.String(), result, expectedResult)
	}
}
