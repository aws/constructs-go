// A programming model for software-defined state
package constructs

import (
	_init_ "github.com/aws/constructs-go/constructs/v10/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents the building block of the construct graph.
//
// All constructs besides the root construct must be created within the scope of
// another construct.
type Construct interface {
	IConstruct
	// The tree node.
	Node() Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Construct
type jsiiProxy_Construct struct {
	jsiiProxy_IConstruct
}

func (j *jsiiProxy_Construct) Node() Node {
	var returns Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Creates a new construct node.
func NewConstruct(scope Construct, id *string) Construct {
	_init_.Initialize()

	j := jsiiProxy_Construct{}

	_jsii_.Create(
		"constructs.Construct",
		[]interface{}{scope, id},
		&j,
	)

	return &j
}

// Creates a new construct node.
func NewConstruct_Override(c Construct, scope Construct, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"constructs.Construct",
		[]interface{}{scope, id},
		c,
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
func Construct_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"constructs.Construct",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Construct) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// In what order to return constructs.
type ConstructOrder string

const (
	// Depth-first, pre-order.
	ConstructOrder_PREORDER ConstructOrder = "PREORDER"
	// Depth-first, post-order (leaf nodes first).
	ConstructOrder_POSTORDER ConstructOrder = "POSTORDER"
)

// Trait for IDependable.
//
// Traits are interfaces that are privately implemented by objects. Instead of
// showing up in the public interface of a class, they need to be queried
// explicitly. This is used to implement certain framework features that are
// not intended to be used by Construct consumers, and so should be hidden
// from accidental use.
//
// Example:
//   // Usage
//   const roots = Dependable.of(construct).dependencyRoots;
//
//   // Definition
//   Dependable.implement(construct, {
//         dependencyRoots: [construct],
//   });
//
// Experimental.
type Dependable interface {
	// The set of constructs that form the root of this dependable.
	//
	// All resources under all returned constructs are included in the ordering
	// dependency.
	// Experimental.
	DependencyRoots() *[]IConstruct
}

// The jsii proxy struct for Dependable
type jsiiProxy_Dependable struct {
	_ byte // padding
}

func (j *jsiiProxy_Dependable) DependencyRoots() *[]IConstruct {
	var returns *[]IConstruct
	_jsii_.Get(
		j,
		"dependencyRoots",
		&returns,
	)
	return returns
}


// Experimental.
func NewDependable_Override(d Dependable) {
	_init_.Initialize()

	_jsii_.Create(
		"constructs.Dependable",
		nil, // no parameters
		d,
	)
}

// Return the matching Dependable for the given class instance.
// Deprecated: use `of`.
func Dependable_Get(instance IDependable) Dependable {
	_init_.Initialize()

	var returns Dependable

	_jsii_.StaticInvoke(
		"constructs.Dependable",
		"get",
		[]interface{}{instance},
		&returns,
	)

	return returns
}

// Turn any object into an IDependable.
// Experimental.
func Dependable_Implement(instance IDependable, trait Dependable) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"constructs.Dependable",
		"implement",
		[]interface{}{instance, trait},
	)
}

// Return the matching Dependable for the given class instance.
// Experimental.
func Dependable_Of(instance IDependable) Dependable {
	_init_.Initialize()

	var returns Dependable

	_jsii_.StaticInvoke(
		"constructs.Dependable",
		"of",
		[]interface{}{instance},
		&returns,
	)

	return returns
}

// A set of constructs to be used as a dependable.
//
// This class can be used when a set of constructs which are disjoint in the
// construct tree needs to be combined to be used as a single dependable.
// Experimental.
type DependencyGroup interface {
	IDependable
	// Add a construct to the dependency roots.
	// Experimental.
	Add(scopes ...IDependable)
}

// The jsii proxy struct for DependencyGroup
type jsiiProxy_DependencyGroup struct {
	jsiiProxy_IDependable
}

// Experimental.
func NewDependencyGroup(deps ...IDependable) DependencyGroup {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	j := jsiiProxy_DependencyGroup{}

	_jsii_.Create(
		"constructs.DependencyGroup",
		args,
		&j,
	)

	return &j
}

// Experimental.
func NewDependencyGroup_Override(d DependencyGroup, deps ...IDependable) {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.Create(
		"constructs.DependencyGroup",
		args,
		d,
	)
}

func (d *jsiiProxy_DependencyGroup) Add(scopes ...IDependable) {
	args := []interface{}{}
	for _, a := range scopes {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		d,
		"add",
		args,
	)
}

// Represents a construct.
type IConstruct interface {
	IDependable
	// The tree node.
	Node() Node
}

// The jsii proxy for IConstruct
type jsiiProxy_IConstruct struct {
	jsiiProxy_IDependable
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

// Trait marker for classes that can be depended upon.
//
// The presence of this interface indicates that an object has
// an `IDependable` implementation.
//
// This interface can be used to take an (ordering) dependency on a set of
// constructs. An ordering dependency implies that the resources represented by
// those constructs are deployed before the resources depending ON them are
// deployed.
type IDependable interface {
}

// The jsii proxy for IDependable
type jsiiProxy_IDependable struct {
	_ byte // padding
}

// Implement this interface in order for the construct to be able to validate itself.
type IValidation interface {
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if there the construct is valid.
	Validate() *[]*string
}

// The jsii proxy for IValidation
type jsiiProxy_IValidation struct {
	_ byte // padding
}

func (i *jsiiProxy_IValidation) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// An entry in the construct metadata table.
type MetadataEntry struct {
	// The data.
	Data interface{} `field:"required" json:"data" yaml:"data"`
	// The metadata entry type.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Stack trace at the point of adding the metadata.
	//
	// Only available if `addMetadata()` is called with `stackTrace: true`.
	Trace *[]*string `field:"optional" json:"trace" yaml:"trace"`
}

// Options for `construct.addMetadata()`.
type MetadataOptions struct {
	// Include stack trace with metadata entry.
	StackTrace *bool `field:"optional" json:"stackTrace" yaml:"stackTrace"`
	// A JavaScript function to begin tracing from.
	//
	// This option is ignored unless `stackTrace` is `true`.
	TraceFromFunction interface{} `field:"optional" json:"traceFromFunction" yaml:"traceFromFunction"`
}

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
	// Return all dependencies registered on this node (non-recursive).
	Dependencies() *[]IConstruct
	// The id of this construct within the current scope.
	//
	// This is a a scope-unique id. To obtain an app-unique id for this construct, use `addr`.
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
	// Add an ordering dependency on another construct.
	//
	// An `IDependable`.
	AddDependency(deps ...IDependable)
	// Adds a metadata entry to this construct.
	//
	// Entries are arbitrary values and will also include a stack trace to allow tracing back to
	// the code location for when the entry was added. It can be used, for example, to include source
	// mapping in CloudFormation templates to improve diagnostics.
	AddMetadata(type_ *string, data interface{}, options *MetadataOptions)
	// Adds a validation to this construct.
	//
	// When `node.validate()` is called, the `validate()` method will be called on
	// all validations and all errors will be returned.
	AddValidation(validation IValidation)
	// Return this construct and all of its children in the given order.
	FindAll(order ConstructOrder) *[]IConstruct
	// Return a direct child by id.
	//
	// Throws an error if the child is not found.
	//
	// Returns: Child with the given id.
	FindChild(id *string) IConstruct
	// Locks this construct from allowing more children to be added.
	//
	// After this
	// call, no more children can be added to this construct or to any children.
	Lock()
	// This can be used to set contextual values.
	//
	// Context must be set before any children are added, since children may consult context info during construction.
	// If the key already exists, it will be overridden.
	SetContext(key *string, value interface{})
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
	// Validates this construct.
	//
	// Invokes the `validate()` method on all validations added through
	// `addValidation()`.
	//
	// Returns: an array of validation error messages associated with this
	// construct.
	Validate() *[]*string
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

func (j *jsiiProxy_Node) Dependencies() *[]IConstruct {
	var returns *[]IConstruct
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


func NewNode(host Construct, scope IConstruct, id *string) Node {
	_init_.Initialize()

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

func (j *jsiiProxy_Node) SetDefaultChild(val IConstruct) {
	_jsii_.Set(
		j,
		"defaultChild",
		val,
	)
}

// Returns the node associated with a construct.
// Deprecated: use `construct.node` instead
func Node_Of(construct IConstruct) Node {
	_init_.Initialize()

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

func (n *jsiiProxy_Node) AddDependency(deps ...IDependable) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addDependency",
		args,
	)
}

func (n *jsiiProxy_Node) AddMetadata(type_ *string, data interface{}, options *MetadataOptions) {
	_jsii_.InvokeVoid(
		n,
		"addMetadata",
		[]interface{}{type_, data, options},
	)
}

func (n *jsiiProxy_Node) AddValidation(validation IValidation) {
	_jsii_.InvokeVoid(
		n,
		"addValidation",
		[]interface{}{validation},
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
	var returns IConstruct

	_jsii_.Invoke(
		n,
		"findChild",
		[]interface{}{id},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Node) Lock() {
	_jsii_.InvokeVoid(
		n,
		"lock",
		nil, // no parameters
	)
}

func (n *jsiiProxy_Node) SetContext(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		n,
		"setContext",
		[]interface{}{key, value},
	)
}

func (n *jsiiProxy_Node) TryFindChild(id *string) IConstruct {
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
	var returns *bool

	_jsii_.Invoke(
		n,
		"tryRemoveChild",
		[]interface{}{childName},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Node) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

