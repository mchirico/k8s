package simpletree

import (
	"reflect"
	"testing"
)

func TestNewT(t *testing.T) {
	tests := []struct {
		name string
		want *Tree
	}{
		// TODO: Add test cases.
		{name: "Initialize",want: NewT()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewT(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_AddN0(t *testing.T) {
	type fields struct {
		Text string
		N0   *Tree
		N1   *Tree
	}
	type args struct {
		text string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &Tree{
				Text: tt.fields.Text,
				N0:   tt.fields.N0,
				N1:   tt.fields.N1,
			}
			tree.AddN0(tt.args.text)
		})
	}
}
