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
	Node() Node
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
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
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

// Returns a string representation of this construct.
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
	ConstructOrder_PREORDER ConstructOrder = "PREORDER"
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
// TODO: EXAMPLE
//
// Experimental.
type Dependable interface {
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
// Deprecated: use `of`
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

// Add a construct to the dependency roots.
// Experimental.
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
// an `IDependableTrait` implementation.
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
//
// Implement this interface in order for the construct to be able to validate itself.
type IValidation interface {
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
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
	Data interface{} `json:"data"`
	// The metadata entry type.
	Type *string `json:"type"`
	// Stack trace at the point of adding the metadata.
	//
	// Only available if `addMetadata()` is called with `stackTrace: true`.
	Trace *[]*string `json:"trace"`
}

// Options for `construct.addMetadata()`.
type MetadataOptions struct {
	// Include stack trace with metadata entry.
	StackTrace *bool `json:"stackTrace"`
	// A JavaScript function to begin tracing from.
	//
	// This option is ignored unless `stackTrace` is `true`.
	TraceFromFunction interface{} `json:"traceFromFunction"`
}

// Represents the construct node in the scope tree.
type Node interface {
	Addr() *string
	Children() *[]IConstruct
	DefaultChild() IConstruct
	SetDefaultChild(val IConstruct)
	Dependencies() *[]IConstruct
	Id() *string
	Locked() *bool
	Metadata() *[]*MetadataEntry
	Path() *string
	Root() IConstruct
	Scope() IConstruct
	Scopes() *[]IConstruct
	AddDependency(deps ...IDependable)
	AddMetadata(type_ *string, data interface{}, options *MetadataOptions)
	AddValidation(validation IValidation)
	FindAll(order ConstructOrder) *[]IConstruct
	FindChild(id *string) IConstruct
	Lock()
	SetContext(key *string, value interface{})
	TryFindChild(id *string) IConstruct
	TryGetContext(key *string) interface{}
	TryRemoveChild(childName *string) *bool
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

// Add an ordering dependency on another construct.
//
// An `IDependable`
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

// Adds a metadata entry to this construct.
//
// Entries are arbitrary values and will also include a stack trace to allow tracing back to
// the code location for when the entry was added. It can be used, for example, to include source
// mapping in CloudFormation templates to improve diagnostics.
func (n *jsiiProxy_Node) AddMetadata(type_ *string, data interface{}, options *MetadataOptions) {
	_jsii_.InvokeVoid(
		n,
		"addMetadata",
		[]interface{}{type_, data, options},
	)
}

// Adds a validation to this construct.
//
// When `node.validate()` is called, the `validate()` method will be called on
// all validations and all errors will be returned.
func (n *jsiiProxy_Node) AddValidation(validation IValidation) {
	_jsii_.InvokeVoid(
		n,
		"addValidation",
		[]interface{}{validation},
	)
}

// Return this construct and all of its children in the given order.
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

// Return a direct child by id.
//
// Throws an error if the child is not found.
//
// Returns: Child with the given id.
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

// Locks this construct from allowing more children to be added.
//
// After this
// call, no more children can be added to this construct or to any children.
func (n *jsiiProxy_Node) Lock() {
	_jsii_.InvokeVoid(
		n,
		"lock",
		nil, // no parameters
	)
}

// This can be used to set contextual values.
//
// Context must be set before any children are added, since children may consult context info during construction.
// If the key already exists, it will be overridden.
func (n *jsiiProxy_Node) SetContext(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		n,
		"setContext",
		[]interface{}{key, value},
	)
}

// Return a direct child by id, or undefined.
//
// Returns: the child if found, or undefined
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

// Retrieves a value from tree context.
//
// Context is usually initialized at the root, but can be overridden at any point in the tree.
//
// Returns: The context value or `undefined` if there is no context value for thie key.
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

// Remove the child with the given name, if present.
//
// Returns: Whether a child with the given name was deleted.
// Experimental.
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

// Validates this construct.
//
// Invokes the `validate()` method on all validations added through
// `addValidation()`.
//
// Returns: an array of validation error messages associated with this
// construct.
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

