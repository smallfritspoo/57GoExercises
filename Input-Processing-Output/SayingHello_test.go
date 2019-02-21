package main

import "testing"

func TestStringCRLFToLF(t *testing.T) {
	if StringCRLFToLF("Test\r\n") != "Test" {
		t.Error("name string was not cleaned correctly")
	}
}
