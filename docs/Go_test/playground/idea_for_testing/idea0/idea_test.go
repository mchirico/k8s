package idea0

import "testing"

func TestFun(t *testing.T) {
	type args struct {
		s   string
		sub string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "test valid", args: args{s: "this is fox",sub: "fox"}, want: "match"},
		{name: "test valid", args: args{s: "this is",sub: "fox"}, want: "no match"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fun(tt.args.s, tt.args.sub); got != tt.want {
				t.Errorf("Fun() = %v, want %v", got, tt.want)
			}
		})
	}
}
