package ds

import "testing"

func TestUniq(t *testing.T) {
	aSlice := []interface{}{1, 2, 3, 3, 4}
	expected := []interface{}{1, 2, 3, 4}

	uniqueSlice := Uniq(aSlice)

	for i, element := range uniqueSlice {
		if expected[i] != element {
			t.Errorf("Expected %v but got %v", expected[i], element)
		}
	}
}
