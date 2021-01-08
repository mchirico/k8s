package generator

import (
	"strings"
	"testing"
)

func Test_basic_generator(t *testing.T) {

	expected := "test"

	result := generator("test")
	out := <-result
	if !strings.Contains(out, expected) {
		t.Fatalf("We got: ->%v<-\nExpected: %s\n", out, expected)
	}

}

func Test_errorOnThis(t *testing.T) {
	errorOnThis()
	t.Log("\n\n *********** \n\n")
}

func Test_GoClosure(t *testing.T) {
	goClosure()
}




