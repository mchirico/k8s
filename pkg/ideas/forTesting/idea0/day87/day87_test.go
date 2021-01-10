package day87

import "testing"

func Test_example1(t *testing.T) {
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
		{name: "Test match",args: args{s: "This is a fox",sub: "fox"}, want: "match", wantErr: false},
		{name: "Test no match",args: args{s: "This is a fox",sub: "bunny"}, want: "no match", wantErr: false},
		{name: "Test error",args: args{s: "This is a fox",sub: ""}, want: "no match", wantErr: true},
		{name: "Test error",args: args{s: "",sub: "fox"}, want: "no match", wantErr: true},
	
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := example1(tt.args.s, tt.args.sub)
			if (err != nil) != tt.wantErr {
				t.Errorf("example1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("example1() = %v, want %v", got, tt.want)
			}
		})
	}
}
