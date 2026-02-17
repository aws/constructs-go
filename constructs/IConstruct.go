package constructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a construct.
type IConstruct interface {
	IDependable
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	//
	// Returns: This construct for chaining.
	With(mixins ...IMixin) IConstruct
	// The tree node.
	Node() Node
}

// The jsii proxy for IConstruct
type jsiiProxy_IConstruct struct {
	jsiiProxy_IDependable
}

func (i *jsiiProxy_IConstruct) With(mixins ...IMixin) IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IConstruct) Node() Node {
	var returns Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

