package cat

import "testing"

func TestAgeInCatYears(t *testing.T) {
	age := AgeInCatYears(20)
	expectedAge := 140
	if age != expectedAge {
		t.Errorf("Expected age %d but got %d", age, expectedAge)
	}
}
