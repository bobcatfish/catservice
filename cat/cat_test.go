package cat

import (
	"fmt"
	"testing"
)

func TestAgeInCatYears(t *testing.T) {
	tests := []struct {
		humanYears int
		catYears   int
	}{{
		humanYears: 20,
		catYears:   140,
	}, {
		humanYears: 1,
		catYears:   7,
	}, {
		humanYears: 4,
		catYears:   28,
	}}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%d", tc.humanYears), func(t *testing.T) {
			actual := AgeInCatYears(tc.humanYears)
			if actual != tc.catYears {
				t.Errorf("Expected %d in cat years when %d in human years but got %d", tc.catYears, tc.humanYears, actual)
			}
		})
	}
}
