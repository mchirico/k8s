package selectexample

import "testing"

func Test_simpleSelect(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "Test example"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			simpleSelect()
		})
	}
}
