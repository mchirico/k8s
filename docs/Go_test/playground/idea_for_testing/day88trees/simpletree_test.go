package simpletree

import (
	"reflect"
	"testing"
)

func Test_subset(t *testing.T) {
	type args struct {
		seq1 []int
		seq2 []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "Test true", args: args{seq1: []int{1, 2, 3, 4, 5}, seq2: []int{4, 2}}, want: true},
		{name: "Test false", args: args{seq1: []int{1, 2, 3, 4, 5}, seq2: []int{4, 2, 7}}, want: false},
		{name: "Test false", args: args{seq1: []int{1, 2, 3, 4, 5}, seq2: []int{14, 12, 7}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subset(tt.args.seq1, tt.args.seq2); got != tt.want {
				t.Errorf("subset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numbersMissing(t *testing.T) {
	type args struct {
		seq1 []int
		seq2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "Test true", args: args{seq1: []int{1, 2, 3, 4, 5}, seq2: []int{4, 2}}, want: []int{}},
		{name: "Test false", args: args{seq1: []int{1, 2, 3, 4, 5}, seq2: []int{4, 2, 7}}, want: []int{7}},
		{name: "Test false", args: args{seq1: []int{1, 2, 3, 4, 5}, seq2: []int{44, 22, 7}}, want: []int{44,22,7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numbersMissing(tt.args.seq1, tt.args.seq2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numbersMissing() = %v, want %v", got, tt.want)
			}
		})
	}
}
