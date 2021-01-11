package goroutine

import (
	"reflect"
	"testing"
)

func Test_ex0(t *testing.T) {
	tests := []struct {
		name    string
		want    map[string]string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ex0()
			if (err != nil) != tt.wantErr {
				t.Errorf("ex0() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ex0() = %v, want %v", got, tt.want)
			}
		})
	}
}
