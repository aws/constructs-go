// A programming model for composable configuration
package constructs

import (
	_init_ "github.com/aws/constructs-go/constructs/v3/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go"
)

// Represents the building block of the construct graph.
//
// All constructs besides the root construct must be created within the scope of
// another construct.
type Construct interface {
	IConstruct
	OnPrepare()
	OnSynthesize(session ISynthesisSession)
	OnValidate() []string
	ToString() string
}

// The jsii proxy struct for Construct
type construct_jsiiProxy struct {
	iConstruct_jsiiProxy // implements constructs.IConstruct
}

// Creates a new construct node.
func NewConstruct(scope Construct, id string, options ConstructOptions) Construct {
	_init_.Initialize()

	c := construct_jsiiProxy{}

	_jsii_.Create(
		"constructs.Construct",
		[]interface{}{scope, id, options},
		[]_jsii_.FQN{"constructs.IConstruct"},
		nil, // no overrides
		&c,
	)

	return &c
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
func (c *construct_jsiiProxy) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil /* no parameters */,
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
func (c *construct_jsiiProxy) OnSynthesize(session ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (c *construct_jsiiProxy) OnValidate() []string {
	var returns []string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil /* no parameters */,
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (c *construct_jsiiProxy) ToString() string {
	var returns string

	_jsii_.Invoke(
		c,
		"toString",
		nil /* no parameters */,
		&returns,
	)

	return returns
}

// Metadata keys used by constructs.
type ConstructMetadata interface {
}

// The jsii proxy struct for ConstructMetadata
type constructMetadata_jsiiProxy struct {
	_ byte // padding
}

func ConstructMetadata_DisableStackTraceInMetadata() string {
	_init_.Initialize()
	var returns string
	_jsii_.StaticGet(
		"constructs.ConstructMetadata",
		"DISABLE_STACK_TRACE_IN_METADATA",
		&returns,
	)
	return returns
}

func ConstructMetadata_ErrorMetadataKey() string {
	_init_.Initialize()
	var returns string
	_jsii_.StaticGet(
		"constructs.ConstructMetadata",
		"ERROR_METADATA_KEY",
		&returns,
	)
	return returns
}

func ConstructMetadata_InfoMetadataKey() string {
	_init_.Initialize()
	var returns string
	_jsii_.StaticGet(
		"constructs.ConstructMetadata",
		"INFO_METADATA_KEY",
		&returns,
	)
	return returns
}

func ConstructMetadata_WarningMetadataKey() string {
	_init_.Initialize()
	var returns string
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
	NodeFactory INodeFactory `json:"nodeFactory"`
}

// In what order to return constructs.
type ConstructOrder string

const (
	ConstructOrder_PREORDER ConstructOrder = "PREORDER"
	ConstructOrder_POSTORDER ConstructOrder = "POSTORDER"
)

// A single dependency.
type Dependency struct {
	// Source the dependency.
	Source IConstruct `json:"source"`
	// Target of the dependency.
	Target IConstruct `json:"target"`
}

// Represents an Aspect.
type IAspect interface {
	// All aspects can visit an IConstruct.
	Visit(node IConstruct)
}

// The jsii proxy for IAspect
type iAspect_jsiiProxy struct {
	_ byte // padding
}

func (i *iAspect_jsiiProxy) Visit(node IConstruct) {
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
type iConstruct_jsiiProxy struct {
	_ byte // padding
}

// A factory for attaching `Node`s to the construct.
type INodeFactory interface {
	// Returns a new `Node` associated with `host`.
	CreateNode(host Construct, scope IConstruct, id string) Node
}

// The jsii proxy for INodeFactory
type iNodeFactory_jsiiProxy struct {
	_ byte // padding
}

func (i *iNodeFactory_jsiiProxy) CreateNode(host Construct, scope IConstruct, id string) Node {
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
	Outdir() string
}

// The jsii proxy for ISynthesisSession
type iSynthesisSession_jsiiProxy struct {
	_ byte // padding
}

func (i *iSynthesisSession_jsiiProxy) Outdir() string {
	var returns string
	_jsii_.Get(
		i,
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
	Validate() []string
}

// The jsii proxy for IValidation
type iValidation_jsiiProxy struct {
	_ byte // padding
}

func (i *iValidation_jsiiProxy) Validate() []string {
	var returns []string

	_jsii_.Invoke(
		i,
		"validate",
		nil /* no parameters */,
		&returns,
	)

	return returns
}

// An entry in the construct metadata table.
type MetadataEntry struct {
	// The data.
	Data interface{} `json:"data"`
	// The metadata entry type.
	Type string `json:"type"`
	// Stack trace.
	//
	// Can be omitted by setting the context key
	// `ConstructMetadata.DISABLE_STACK_TRACE_IN_METADATA` to 1.
	Trace []string `json:"trace"`
}

// Represents the construct node in the scope tree.
type Node interface {
	Addr() string
	Children() []IConstruct
	DefaultChild() IConstruct
	SetDefaultChild(val IConstruct)
	Dependencies() []Dependency
	Id() string
	Locked() bool
	Metadata() []MetadataEntry
	Path() string
	Root() IConstruct
	Scope() IConstruct
	Scopes() []IConstruct
	UniqueId() string
	AddDependency(dependencies IConstruct)
	AddError(message string)
	AddInfo(message string)
	AddMetadata(type_ string, data interface{}, fromFunction interface{})
	AddValidation(validation IValidation)
	AddWarning(message string)
	ApplyAspect(aspect IAspect)
	FindAll(order ConstructOrder) []IConstruct
	FindChild(id string) IConstruct
	Prepare()
	SetContext(key string, value interface{})
	Synthesize(options SynthesisOptions)
	TryFindChild(id string) IConstruct
	TryGetContext(key string) interface{}
	TryRemoveChild(childName string) bool
	Validate() []ValidationError
}

// The jsii proxy struct for Node
type node_jsiiProxy struct {
	_ byte // padding
}

func (n *node_jsiiProxy) Addr() string {
	var returns string
	_jsii_.Get(
		n,
		"addr",
		&returns,
	)
	return returns
}

func (n *node_jsiiProxy) Children() []IConstruct {
	var returns []IConstruct
	_jsii_.Get(
		n,
		"children",
		&returns,
	)
	return returns
}

func (n *node_jsiiProxy) DefaultChild() IConstruct {
	var returns IConstruct
	_jsii_.Get(
		n,
		"defaultChild",
		&returns,
	)
	return returns
}

func (n *node_jsiiProxy) Dependencies() []Dependency {
	var returns []Dependency
	_jsii_.Get(
		n,
		"dependencies",
		&returns,
	)
	return returns
}

func (n *node_jsiiProxy) Id() string {
	var returns string
	_jsii_.Get(
		n,
		"id",
		&returns,
	)
	return returns
}

func (n *node_jsiiProxy) Locked() bool {
	var returns bool
	_jsii_.Get(
		n,
		"locked",
		&returns,
	)
	return returns
}

func (n *node_jsiiProxy) Metadata() []MetadataEntry {
	var returns []MetadataEntry
	_jsii_.Get(
		n,
		"metadata",
		&returns,
	)
	return returns
}

func (n *node_jsiiProxy) Path() string {
	var returns string
	_jsii_.Get(
		n,
		"path",
		&returns,
	)
	return returns
}

func (n *node_jsiiProxy) Root() IConstruct {
	var returns IConstruct
	_jsii_.Get(
		n,
		"root",
		&returns,
	)
	return returns
}

func (n *node_jsiiProxy) Scope() IConstruct {
	var returns IConstruct
	_jsii_.Get(
		n,
		"scope",
		&returns,
	)
	return returns
}

func (n *node_jsiiProxy) Scopes() []IConstruct {
	var returns []IConstruct
	_jsii_.Get(
		n,
		"scopes",
		&returns,
	)
	return returns
}

func (n *node_jsiiProxy) UniqueId() string {
	var returns string
	_jsii_.Get(
		n,
		"uniqueId",
		&returns,
	)
	return returns
}


func NewNode(host Construct, scope IConstruct, id string) Node {
	_init_.Initialize()

	n := node_jsiiProxy{}

	_jsii_.Create(
		"constructs.Node",
		[]interface{}{host, scope, id},
		[]_jsii_.FQN{},
		nil, // no overrides
		&n,
	)

	return &n
}

func (n *node_jsiiProxy) SetDefaultChild(val IConstruct) {
	_jsii_.Set(
		n,
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

func Node_PathSep() string {
	_init_.Initialize()
	var returns string
	_jsii_.StaticGet(
		"constructs.Node",
		"PATH_SEP",
		&returns,
	)
	return returns
}

// Add an ordering dependency on another Construct.
//
// All constructs in the dependency's scope will be deployed before any
// construct in this construct's scope.
func (n *node_jsiiProxy) AddDependency(dependencies IConstruct) {
	_jsii_.InvokeVoid(
		n,
		"addDependency",
		[]interface{}{dependencies},
	)
}

// Adds an { "error": <message> } metadata entry to this construct.
//
// The toolkit will fail synthesis when errors are reported.
func (n *node_jsiiProxy) AddError(message string) {
	_jsii_.InvokeVoid(
		n,
		"addError",
		[]interface{}{message},
	)
}

// Adds a { "info": <message> } metadata entry to this construct.
//
// The toolkit will display the info message when apps are synthesized.
func (n *node_jsiiProxy) AddInfo(message string) {
	_jsii_.InvokeVoid(
		n,
		"addInfo",
		[]interface{}{message},
	)
}

// Adds a metadata entry to this construct.
//
// Entries are arbitrary values and will also include a stack trace to allow tracing back to
// the code location for when the entry was added. It can be used, for example, to include source
// mapping in CloudFormation templates to improve diagnostics.
func (n *node_jsiiProxy) AddMetadata(type_ string, data interface{}, fromFunction interface{}) {
	_jsii_.InvokeVoid(
		n,
		"addMetadata",
		[]interface{}{type_, data, fromFunction},
	)
}

// Adds a validation to this construct.
//
// When `node.validate()` is called, the `validate()` method will be called on
// all validations and all errors will be returned.
func (n *node_jsiiProxy) AddValidation(validation IValidation) {
	_jsii_.InvokeVoid(
		n,
		"addValidation",
		[]interface{}{validation},
	)
}

// Adds a { "warning": <message> } metadata entry to this construct.
//
// The toolkit will display the warning when an app is synthesized, or fail
// if run in --strict mode.
func (n *node_jsiiProxy) AddWarning(message string) {
	_jsii_.InvokeVoid(
		n,
		"addWarning",
		[]interface{}{message},
	)
}

// Applies the aspect to this Constructs node.
func (n *node_jsiiProxy) ApplyAspect(aspect IAspect) {
	_jsii_.InvokeVoid(
		n,
		"applyAspect",
		[]interface{}{aspect},
	)
}

// Return this construct and all of its children in the given order.
func (n *node_jsiiProxy) FindAll(order ConstructOrder) []IConstruct {
	var returns []IConstruct

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
func (n *node_jsiiProxy) FindChild(id string) IConstruct {
	var returns IConstruct

	_jsii_.Invoke(
		n,
		"findChild",
		[]interface{}{id},
		&returns,
	)

	return returns
}

// Invokes "prepare" on all constructs (depth-first, post-order) in the tree under `node`.
func (n *node_jsiiProxy) Prepare() {
	_jsii_.InvokeVoid(
		n,
		"prepare",
		nil /* no parameters */,
	)
}

// This can be used to set contextual values.
//
// Context must be set before any children are added, since children may consult context info during construction.
// If the key already exists, it will be overridden.
func (n *node_jsiiProxy) SetContext(key string, value interface{}) {
	_jsii_.InvokeVoid(
		n,
		"setContext",
		[]interface{}{key, value},
	)
}

// Synthesizes a CloudAssembly from a construct tree.
func (n *node_jsiiProxy) Synthesize(options SynthesisOptions) {
	_jsii_.InvokeVoid(
		n,
		"synthesize",
		[]interface{}{options},
	)
}

// Return a direct child by id, or undefined.
//
// Returns: the child if found, or undefined
func (n *node_jsiiProxy) TryFindChild(id string) IConstruct {
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
func (n *node_jsiiProxy) TryGetContext(key string) interface{} {
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
func (n *node_jsiiProxy) TryRemoveChild(childName string) bool {
	var returns bool

	_jsii_.Invoke(
		n,
		"tryRemoveChild",
		[]interface{}{childName},
		&returns,
	)

	return returns
}

// Validates tree (depth-first, pre-order) and returns the list of all errors.
//
// An empty list indicates that there are no errors.
func (n *node_jsiiProxy) Validate() []ValidationError {
	var returns []ValidationError

	_jsii_.Invoke(
		n,
		"validate",
		nil /* no parameters */,
		&returns,
	)

	return returns
}

// Options for synthesis.
type SynthesisOptions struct {
	// The output directory into which to synthesize the cloud assembly.
	Outdir string `json:"outdir"`
	// Additional context passed into the synthesis session object when `construct.synth` is called.
	SessionContext map[string]interface{} `json:"sessionContext"`
	// Whether synthesis should skip the validation phase.
	SkipValidation bool `json:"skipValidation"`
}

// An error returned during the validation phase.
type ValidationError struct {
	// The error message.
	Message string `json:"message"`
	// The construct which emitted the error.
	Source Construct `json:"source"`
}

