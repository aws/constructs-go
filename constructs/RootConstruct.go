package constructs

import (
	_init_ "github.com/aws/constructs-go/constructs/v10/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Creates a new root construct node.
//
// The root construct represents the top of the construct tree and is not contained within a parent scope itself.
// For root constructs, the id is optional.
type RootConstruct interface {
	Construct
	// The tree node.
	Node() Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for RootConstruct
type jsiiProxy_RootConstruct struct {
	jsiiProxy_Construct
}

func (j *jsiiProxy_RootConstruct) Node() Node {
	var returns Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Creates a new root construct node.
func NewRootConstruct(id *string) RootConstruct {
	_init_.Initialize()

	j := jsiiProxy_RootConstruct{}

	_jsii_.Create(
		"constructs.RootConstruct",
		[]interface{}{id},
		&j,
	)

	return &j
}

// Creates a new root construct node.
func NewRootConstruct_Override(r RootConstruct, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"constructs.RootConstruct",
		[]interface{}{id},
		r,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func RootConstruct_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRootConstruct_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"constructs.RootConstruct",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RootConstruct) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

