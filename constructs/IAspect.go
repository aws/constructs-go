package constructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents an Aspect.
type IAspect interface {
	// All aspects can visit an IConstruct.
	Visit(node IConstruct)
}

// The jsii proxy for IAspect
type jsiiProxy_IAspect struct {
	_ byte // padding
}

func (i *jsiiProxy_IAspect) Visit(node IConstruct) {
	if err := i.validateVisitParameters(node); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"visit",
		[]interface{}{node},
	)
}

