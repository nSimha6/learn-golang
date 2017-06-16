package main

import "testing"

func TestHalf(t *testing.T) {
	actualResult := Half(1)
	var expectedResult = "0 false"

	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}
