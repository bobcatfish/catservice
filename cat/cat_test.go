package cat

import "testing"

func TestAgeInCatYears(t *testing.T) {
	age := AgeInCatYears(200)
	expectedAge := 200
	if age != expectedAge {
		t.Errorf("Expected age %d but got %d", age, expectedAge)
	}
}
