package playground

import "testing"

func TestPlayground1(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "test 1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Playground1()
		})
	}
}
