package cat

import "testing"

func TestGetCatsOfTekton(t *testing.T) {
	cats := GetCatsOfTekton()
	if cats[0].Name != "Acadia" {
		t.Errorf("Expected first cat to be Acadia but was %s", cats[0].Name)
	}
}
