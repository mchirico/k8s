package pool2

import (
	"fmt"
	"sync"
)

// Pool for our struct A
var pool *sync.Pool

// A dummy struct with a member
type dataContainer struct {
	Name string
	H map[string]string
}





// Func to init pool
func initPool() {
	pool = &sync.Pool{
		New: func() interface{} {
			fmt.Println("Returning new dataContainer")
			return new(dataContainer)
		},
	}
}


func InitPutA() {
	// Initializing pool
	initPool()
	// Get hold of instance one
	one := pool.Get().(*dataContainer)
	one.Name = "first"
	one.H = map[string]string{}
	one.H["bob"] = "one"
	fmt.Printf("one.Name = %s\n", one.Name)
	fmt.Printf("one.H = %s\n", one.H)
	// Submit back the instance after using
	pool.Put(one)
	// Now the same instance becomes usable by another routine without allocating it again
}