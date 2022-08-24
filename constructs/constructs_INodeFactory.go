// A programming model for composable configuration
package constructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A factory for attaching `Node`s to the construct.
type INodeFactory interface {
	// Returns a new `Node` associated with `host`.
	CreateNode(host Construct, scope IConstruct, id *string) Node
}

// The jsii proxy for INodeFactory
type jsiiProxy_INodeFactory struct {
	_ byte // padding
}

func (i *jsiiProxy_INodeFactory) CreateNode(host Construct, scope IConstruct, id *string) Node {
	var returns Node

	_jsii_.Invoke(
		i,
		"createNode",
		[]interface{}{host, scope, id},
		&returns,
	)

	return returns
}

