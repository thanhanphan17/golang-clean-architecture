package hashprovider

import "testing"

func TestMD5Hash_Hash(t *testing.T) {
	h := &md5Hash{}

	// Test case 1: Empty string
	input1 := ""
	expectedOutput1 := "d41d8cd98f00b204e9800998ecf8427e"
	output1 := h.Hash(input1)
	if output1 != expectedOutput1 {
		t.Errorf("Expected %s, but got %s", expectedOutput1, output1)
	}

	// Test case 2: Non-empty string
	input2 := "Hello, World!"
	expectedOutput2 := "65a8e27d8879283831b664bd8b7f0ad4"
	output2 := h.Hash(input2)
	if output2 != expectedOutput2 {
		t.Errorf("Expected %s, but got %s", expectedOutput2, output2)
	}

	// Test case 3: Special characters
	input3 := "!@#$%^&*()"
	expectedOutput3 := "05b28d17a7b6e7024b6e5d8cc43a8bf7"
	output3 := h.Hash(input3)
	if output3 != expectedOutput3 {
		t.Errorf("Expected %s, but got %s", expectedOutput3, output3)
	}
}
