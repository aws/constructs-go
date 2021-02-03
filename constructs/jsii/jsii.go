package jsii

import (
	rt "github.com/aws/jsii-runtime-go"
	"sync"
)

var once sync.Once

// Initialize performs the necessary work for the enclosing
// module to be loaded in the jsii kernel.
func Initialize() {
	once.Do(func(){
		// Load this library into the kernel
		rt.Load("constructs", "3.3.6", tarball)
	})
}
