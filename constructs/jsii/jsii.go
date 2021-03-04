package jsii

import (
	_      "embed"
	"sync"

	_jsii_ "github.com/aws/jsii-runtime-go"
)

//go:embed constructs-3.3.55.tgz
var tarball []byte
var once    sync.Once

// Initialize performs the necessary work for the enclosing
// module to be loaded in the jsii kernel.
func Initialize() {
	once.Do(func(){
		// Load this library into the kernel
		_jsii_.Load("constructs", "3.3.55", tarball)
	})
}
