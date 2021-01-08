package idea0

import (
	"testing"
)

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
		{name: "test valid", args: args{s: "this is fox", sub: "fox"}, want: "match"},
		{name: "test valid", args: args{s: "this is", sub: "fox"}, want: "no match"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fun(tt.args.s, tt.args.sub); got != tt.want {
				t.Errorf("Fun() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFun2(t *testing.T) {
	type args struct {
		s   string
		sub string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		// TODO: Add test cases.
		{name: "test valid", args: args{s: "this is fox", sub: "fox"}, want: "match",wantErr: false},
		{name: "test valid", args: args{s: "this is", sub: "fox"}, want: "no match", wantErr: false},
		{name: "test Error", args: args{s: "", sub: "fox"}, want: "", wantErr: true},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Fun2(tt.args.s, tt.args.sub)
			if (err != nil) != tt.wantErr {
				t.Errorf("Fun2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Fun2() = %v, want %v", got, tt.want)
			}
		})
	}
}
