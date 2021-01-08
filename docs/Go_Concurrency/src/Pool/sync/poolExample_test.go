package poolexample

import (
	"testing"
	"time"
)

/*
Ref: https://play.golang.org/p/PRa3SOaIXif
*/

func Test_basic_generator(t *testing.T) {

	for i:=0;i<3000; i++ {
	go func() {
		Example1()
	}()

	go func() {
		Example2()
	}()
	}

	time.Sleep(300 * time.Second)
	

}