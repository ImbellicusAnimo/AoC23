package day1

import "testing"

func TestGetAlphaNumCoords(t *testing.T) {
	// input
	testInput := "8four419eighteight1bpv"
	expectedOutput := "8,1"
	// test
	x, y := getAlphaNumCoordinates(testInput)
	got := x + "," + y
	if got != expectedOutput {
		t.Errorf("getAlphaNumCoords(%s) = %s,%s; want %s", testInput, x, y, expectedOutput)
	}
}
