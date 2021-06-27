package mac

import (
	"testing"
)

func TestIsMAC(t *testing.T) {
	if !IsMAC("e4:ce:8f:5b:a7:fe") &&
		!IsMAC("00:28:31:8A:41:C6") &&
		!IsMAC("e4-ce-8f-5b-a7-fe") &&
		!IsMAC("e4:ce:8f:5b:a:e") &&
		!IsMAC("00:28:31:8A:1:6") &&
		!IsMAC("e4-ce-8f-5b-7-e") {
		t.Error("Expected true")
	}

	if IsMAC("e4:ce:8f:5b:a7fe") {
		t.Error("Expected false")
	}
}
