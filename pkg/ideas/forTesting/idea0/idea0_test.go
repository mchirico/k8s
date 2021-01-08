package idea0

import "testing"

func TestEx0(t *testing.T) {
	type args struct {
		s  string
		s2 string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "Simple test", args: args{s: "this fox", s2: "fox"}, want: "fox"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ex0(tt.args.s, tt.args.s2); got != tt.want {
				t.Errorf("Ex0() = %v, want %v", got, tt.want)
			}
		})
	}
}
