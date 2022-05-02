// A programming model for composable configuration
package constructs

import (
	_init_ "github.com/aws/constructs-go/constructs/v3/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents the building block of the construct graph.
//
// All constructs besides the root construct must be created within the scope of
// another construct.
type Construct interface {
	IConstruct
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	OnSynthesize(session ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if there the construct is valid.
	// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
	// instead of overriding this method.
	OnValidate() *[]*string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Construct
type jsiiProxy_Construct struct {
	jsiiProxy_IConstruct
}

// Creates a new construct node.
func NewConstruct(scope Construct, id *string, options *ConstructOptions) Construct {
	_init_.Initialize()

	j := jsiiProxy_Construct{}

	_jsii_.Create(
		"constructs.Construct",
		[]interface{}{scope, id, options},
		&j,
	)

	return &j
}

// Creates a new construct node.
func NewConstruct_Override(c Construct, scope Construct, id *string, options *ConstructOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"constructs.Construct",
		[]interface{}{scope, id, options},
		c,
	)
}

func (c *jsiiProxy_Construct) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Construct) OnSynthesize(session ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_Construct) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
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

// Metadata keys used by constructs.
type ConstructMetadata interface {
}

// The jsii proxy struct for ConstructMetadata
type jsiiProxy_ConstructMetadata struct {
	_ byte // padding
}

func ConstructMetadata_DISABLE_STACK_TRACE_IN_METADATA() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"constructs.ConstructMetadata",
		"DISABLE_STACK_TRACE_IN_METADATA",
		&returns,
	)
	return returns
}

func ConstructMetadata_ERROR_METADATA_KEY() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"constructs.ConstructMetadata",
		"ERROR_METADATA_KEY",
		&returns,
	)
	return returns
}

func ConstructMetadata_INFO_METADATA_KEY() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"constructs.ConstructMetadata",
		"INFO_METADATA_KEY",
		&returns,
	)
	return returns
}

func ConstructMetadata_WARNING_METADATA_KEY() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"constructs.ConstructMetadata",
		"WARNING_METADATA_KEY",
		&returns,
	)
	return returns
}

// Options for creating constructs.
type ConstructOptions struct {
	// A factory for attaching `Node`s to the construct.
	NodeFactory INodeFactory `json:"nodeFactory" yaml:"nodeFactory"`
}

// In what order to return constructs.
type ConstructOrder string

const (
	// Depth-first, pre-order.
	ConstructOrder_PREORDER ConstructOrder = "PREORDER"
	// Depth-first, post-order (leaf nodes first).
	ConstructOrder_POSTORDER ConstructOrder = "POSTORDER"
)

// A single dependency.
type Dependency struct {
	// Source the dependency.
	Source IConstruct `json:"source" yaml:"source"`
	// Target of the dependency.
	Target IConstruct `json:"target" yaml:"target"`
}

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
	_jsii_.InvokeVoid(
		i,
		"visit",
		[]interface{}{node},
	)
}

// Represents a construct.
type IConstruct interface {
}

// The jsii proxy for IConstruct
type jsiiProxy_IConstruct struct {
	_ byte // padding
}

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

// Represents a single session of synthesis.
//
// Passed into `construct.onSynthesize()` methods.
type ISynthesisSession interface {
	// The output directory for this synthesis session.
	Outdir() *string
}

// The jsii proxy for ISynthesisSession
type jsiiProxy_ISynthesisSession struct {
	_ byte // padding
}

func (j *jsiiProxy_ISynthesisSession) Outdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
	)
	return returns
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
	Data interface{} `json:"data" yaml:"data"`
	// The metadata entry type.
	Type *string `json:"type" yaml:"type"`
	// Stack trace.
	//
	// Can be omitted by setting the context key
	// `ConstructMetadata.DISABLE_STACK_TRACE_IN_METADATA` to 1.
	Trace *[]*string `json:"trace" yaml:"trace"`
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
	_jsii_.InvokeVoid(
		n,
		"addError",
		[]interface{}{message},
	)
}

func (n *jsiiProxy_Node) AddInfo(message *string) {
	_jsii_.InvokeVoid(
		n,
		"addInfo",
		[]interface{}{message},
	)
}

func (n *jsiiProxy_Node) AddMetadata(type_ *string, data interface{}, fromFunction interface{}) {
	_jsii_.InvokeVoid(
		n,
		"addMetadata",
		[]interface{}{type_, data, fromFunction},
	)
}

func (n *jsiiProxy_Node) AddValidation(validation IValidation) {
	_jsii_.InvokeVoid(
		n,
		"addValidation",
		[]interface{}{validation},
	)
}

func (n *jsiiProxy_Node) AddWarning(message *string) {
	_jsii_.InvokeVoid(
		n,
		"addWarning",
		[]interface{}{message},
	)
}

func (n *jsiiProxy_Node) ApplyAspect(aspect IAspect) {
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
	_jsii_.InvokeVoid(
		n,
		"setContext",
		[]interface{}{key, value},
	)
}

func (n *jsiiProxy_Node) Synthesize(options *SynthesisOptions) {
	_jsii_.InvokeVoid(
		n,
		"synthesize",
		[]interface{}{options},
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

// Options for synthesis.
type SynthesisOptions struct {
	// The output directory into which to synthesize the cloud assembly.
	Outdir *string `json:"outdir" yaml:"outdir"`
	// Additional context passed into the synthesis session object when `construct.synth` is called.
	SessionContext *map[string]interface{} `json:"sessionContext" yaml:"sessionContext"`
	// Whether synthesis should skip the validation phase.
	SkipValidation *bool `json:"skipValidation" yaml:"skipValidation"`
}

// An error returned during the validation phase.
type ValidationError struct {
	// The error message.
	Message *string `json:"message" yaml:"message"`
	// The construct which emitted the error.
	Source Construct `json:"source" yaml:"source"`
}

