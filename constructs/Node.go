// A programming model for composable configuration
package constructs

import (
	_init_ "github.com/aws/constructs-go/constructs/v3/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents the construct node in the scope tree.
type Node interface {
	// Returns an opaque tree-unique address for this construct.
	//
	// Addresses are 42 characters hexadecimal strings. They begin with "c8"
	// followed by 40 lowercase hexadecimal characters (0-9a-f).
	//
	// Addresses are calculated using a SHA-1 of the components of the construct
	// path.
	//
	// To enable refactorings of construct trees, constructs with the ID `Default`
	// will be excluded from the calculation. In those cases constructs in the
	// same tree may have the same addreess.
	//
	// Example:
	//   c83a2846e506bcc5f10682b564084bca2d275709ee
	//
	Addr() *string
	// All direct children of this construct.
	Children() *[]IConstruct
	// Returns the child construct that has the id `Default` or `Resource"`.
	//
	// This is usually the construct that provides the bulk of the underlying functionality.
	// Useful for modifications of the underlying construct that are not available at the higher levels.
	// Override the defaultChild property.
	//
	// This should only be used in the cases where the correct
	// default child is not named 'Resource' or 'Default' as it
	// should be.
	//
	// If you set this to undefined, the default behavior of finding
	// the child named 'Resource' or 'Default' will be used.
	//
	// Returns: a construct or undefined if there is no default child.
	DefaultChild() IConstruct
	SetDefaultChild(val IConstruct)
	// Return all dependencies registered on this node or any of its children.
	Dependencies() *[]*Dependency
	// The id of this construct within the current scope.
	//
	// This is a a scope-unique id. To obtain an app-unique id for this construct, use `uniqueId`.
	Id() *string
	// Returns true if this construct or the scopes in which it is defined are locked.
	Locked() *bool
	// An immutable array of metadata objects associated with this construct.
	//
	// This can be used, for example, to implement support for deprecation notices, source mapping, etc.
	Metadata() *[]*MetadataEntry
	// The full, absolute path of this construct in the tree.
	//
	// Components are separated by '/'.
	Path() *string
	// Returns the root of the construct tree.
	//
	// Returns: The root of the construct tree.
	Root() IConstruct
	// Returns the scope in which this construct is defined.
	//
	// The value is `undefined` at the root of the construct scope tree.
	Scope() IConstruct
	// All parent scopes of this construct.
	//
	// Returns: a list of parent scopes. The last element in the list will always
	// be the current construct and the first element will be the root of the
	// tree.
	Scopes() *[]IConstruct
	// A tree-global unique alphanumeric identifier for this construct.
	//
	// Includes
	// all components of the tree.
	// Deprecated: please avoid using this property and use `addr` to form unique names.
	// This algorithm uses MD5, which is not FIPS-complient and also excludes the
	// identity of the root construct from the calculation.
	UniqueId() *string
	// Add an ordering dependency on another Construct.
	//
	// All constructs in the dependency's scope will be deployed before any
	// construct in this construct's scope.
	AddDependency(dependencies ...IConstruct)
	// Adds an { "error": <message> } metadata entry to this construct.
	//
	// The toolkit will fail synthesis when errors are reported.
	AddError(message *string)
	// Adds a { "info": <message> } metadata entry to this construct.
	//
	// The toolkit will display the info message when apps are synthesized.
	AddInfo(message *string)
	// Adds a metadata entry to this construct.
	//
	// Entries are arbitrary values and will also include a stack trace to allow tracing back to
	// the code location for when the entry was added. It can be used, for example, to include source
	// mapping in CloudFormation templates to improve diagnostics.
	AddMetadata(type_ *string, data interface{}, fromFunction interface{})
	// Adds a validation to this construct.
	//
	// When `node.validate()` is called, the `validate()` method will be called on
	// all validations and all errors will be returned.
	AddValidation(validation IValidation)
	// Adds a { "warning": <message> } metadata entry to this construct.
	//
	// The toolkit will display the warning when an app is synthesized, or fail
	// if run in --strict mode.
	AddWarning(message *string)
	// Applies the aspect to this Constructs node.
	ApplyAspect(aspect IAspect)
	// Return this construct and all of its children in the given order.
	FindAll(order ConstructOrder) *[]IConstruct
	// Return a direct child by id.
	//
	// Throws an error if the child is not found.
	//
	// Returns: Child with the given id.
	FindChild(id *string) IConstruct
	// Invokes "prepare" on all constructs (depth-first, post-order) in the tree under `node`.
	Prepare()
	// This can be used to set contextual values.
	//
	// Context must be set before any children are added, since children may consult context info during construction.
	// If the key already exists, it will be overridden.
	SetContext(key *string, value interface{})
	// Synthesizes a CloudAssembly from a construct tree.
	Synthesize(options *SynthesisOptions)
	// Return a direct child by id, or undefined.
	//
	// Returns: the child if found, or undefined.
	TryFindChild(id *string) IConstruct
	// Retrieves a value from tree context.
	//
	// Context is usually initialized at the root, but can be overridden at any point in the tree.
	//
	// Returns: The context value or `undefined` if there is no context value for thie key.
	TryGetContext(key *string) interface{}
	// Remove the child with the given name, if present.
	//
	// Returns: Whether a child with the given name was deleted.
	// Experimental.
	TryRemoveChild(childName *string) *bool
	// Validates tree (depth-first, pre-order) and returns the list of all errors.
	//
	// An empty list indicates that there are no errors.
	Validate() *[]*ValidationError
}

// The jsii proxy struct for Node
type jsiiProxy_Node struct {
	_ byte // padding
}

func (j *jsiiProxy_Node) Addr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Node) Children() *[]IConstruct {
	var returns *[]IConstruct
	_jsii_.Get(
		j,
		"children",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Node) DefaultChild() IConstruct {
	var returns IConstruct
	_jsii_.Get(
		j,
		"defaultChild",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Node) Dependencies() *[]*Dependency {
	var returns *[]*Dependency
	_jsii_.Get(
		j,
		"dependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Node) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Node) Locked() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"locked",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Node) Metadata() *[]*MetadataEntry {
	var returns *[]*MetadataEntry
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Node) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Node) Root() IConstruct {
	var returns IConstruct
	_jsii_.Get(
		j,
		"root",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Node) Scope() IConstruct {
	var returns IConstruct
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Node) Scopes() *[]IConstruct {
	var returns *[]IConstruct
	_jsii_.Get(
		j,
		"scopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Node) UniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uniqueId",
		&returns,
	)
	return returns
}


func NewNode(host Construct, scope IConstruct, id *string) Node {
	_init_.Initialize()

	if err := validateNewNodeParameters(host, scope, id); err != nil {
		panic(err)
	}
	j := jsiiProxy_Node{}

	_jsii_.Create(
		"constructs.Node",
		[]interface{}{host, scope, id},
		&j,
	)

	return &j
}

func NewNode_Override(n Node, host Construct, scope IConstruct, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"constructs.Node",
		[]interface{}{host, scope, id},
		n,
	)
}

func (j *jsiiProxy_Node)SetDefaultChild(val IConstruct) {
	_jsii_.Set(
		j,
		"defaultChild",
		val,
	)
}

// Returns the node associated with a construct.
func Node_Of(construct IConstruct) Node {
	_init_.Initialize()

	if err := validateNode_OfParameters(construct); err != nil {
		panic(err)
	}
	var returns Node

	_jsii_.StaticInvoke(
		"constructs.Node",
		"of",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func Node_PATH_SEP() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"constructs.Node",
		"PATH_SEP",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_Node) AddDependency(dependencies ...IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addDependency",
		args,
	)
}

func (n *jsiiProxy_Node) AddError(message *string) {
	if err := n.validateAddErrorParameters(message); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addError",
		[]interface{}{message},
	)
}

func (n *jsiiProxy_Node) AddInfo(message *string) {
	if err := n.validateAddInfoParameters(message); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addInfo",
		[]interface{}{message},
	)
}

func (n *jsiiProxy_Node) AddMetadata(type_ *string, data interface{}, fromFunction interface{}) {
	if err := n.validateAddMetadataParameters(type_, data); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addMetadata",
		[]interface{}{type_, data, fromFunction},
	)
}

func (n *jsiiProxy_Node) AddValidation(validation IValidation) {
	if err := n.validateAddValidationParameters(validation); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addValidation",
		[]interface{}{validation},
	)
}

func (n *jsiiProxy_Node) AddWarning(message *string) {
	if err := n.validateAddWarningParameters(message); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addWarning",
		[]interface{}{message},
	)
}

func (n *jsiiProxy_Node) ApplyAspect(aspect IAspect) {
	if err := n.validateApplyAspectParameters(aspect); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"applyAspect",
		[]interface{}{aspect},
	)
}

func (n *jsiiProxy_Node) FindAll(order ConstructOrder) *[]IConstruct {
	var returns *[]IConstruct

	_jsii_.Invoke(
		n,
		"findAll",
		[]interface{}{order},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Node) FindChild(id *string) IConstruct {
	if err := n.validateFindChildParameters(id); err != nil {
		panic(err)
	}
	var returns IConstruct

	_jsii_.Invoke(
		n,
		"findChild",
		[]interface{}{id},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Node) Prepare() {
	_jsii_.InvokeVoid(
		n,
		"prepare",
		nil, // no parameters
	)
}

func (n *jsiiProxy_Node) SetContext(key *string, value interface{}) {
	if err := n.validateSetContextParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"setContext",
		[]interface{}{key, value},
	)
}

func (n *jsiiProxy_Node) Synthesize(options *SynthesisOptions) {
	if err := n.validateSynthesizeParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"synthesize",
		[]interface{}{options},
	)
}

func (n *jsiiProxy_Node) TryFindChild(id *string) IConstruct {
	if err := n.validateTryFindChildParameters(id); err != nil {
		panic(err)
	}
	var returns IConstruct

	_jsii_.Invoke(
		n,
		"tryFindChild",
		[]interface{}{id},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Node) TryGetContext(key *string) interface{} {
	if err := n.validateTryGetContextParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"tryGetContext",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Node) TryRemoveChild(childName *string) *bool {
	if err := n.validateTryRemoveChildParameters(childName); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		n,
		"tryRemoveChild",
		[]interface{}{childName},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Node) Validate() *[]*ValidationError {
	var returns *[]*ValidationError

	_jsii_.Invoke(
		n,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

