// A programming model for composable configuration
package constructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go"
	_init_ "github.com/aws/constructs-go/constructs/v3/jsii"
	"reflect"
)

// Class interface
type ConstructIface interface {
	IConstructIface
	OnPrepare()
	OnSynthesize(session ISynthesisSessionIface)
	OnValidate() []string
	ToString() string
}

// Represents the building block of the construct graph.
//
// All constructs besides the root construct must be created within the scope of
// another construct.
// Struct proxy
type Construct struct {
}

// Creates a new construct node.
func NewConstruct(scope ConstructIface, id string, options ConstructOptionsIface) ConstructIface {
	_init_.Initialize()
	self := Construct{}
	_jsii_.Create(
		"constructs.Construct",
		[]interface{}{scope, id, options},
		[]_jsii_.FQN{"constructs.IConstruct"},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (c *Construct) OnPrepare() {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"onPrepare",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *Construct) OnSynthesize(session ISynthesisSessionIface) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"onSynthesize",
		[]interface{}{session},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *Construct) OnValidate() []string {
	var returns []string
	_jsii_.Invoke(
		c,
		"onValidate",
		[]interface{}{},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *Construct) ToString() string {
	var returns string
	_jsii_.Invoke(
		c,
		"toString",
		[]interface{}{},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

// Class interface
type ConstructMetadataIface interface {
}

// Metadata keys used by constructs.
// Struct proxy
type ConstructMetadata struct {
}

func ConstructMetadata_DisableStackTraceInMetadata() string {
	_init_.Initialize()
	var returns string
	_jsii_.StaticGet(
		"constructs.ConstructMetadata",
		"DISABLE_STACK_TRACE_IN_METADATA",
		&returns,
		map[reflect.Type]reflect.Type{},
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
		map[reflect.Type]reflect.Type{},
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
		map[reflect.Type]reflect.Type{},
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
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

// ConstructOptionsIface is the public interface for the custom type ConstructOptions
type ConstructOptionsIface interface {
	GetNodeFactory() INodeFactoryIface
}

// Options for creating constructs.
// Struct proxy
type ConstructOptions struct {
	// A factory for attaching `Node`s to the construct.
	NodeFactory INodeFactoryIface `json:"nodeFactory"`
}

func (c *ConstructOptions) GetNodeFactory() INodeFactoryIface {
	var returns INodeFactoryIface
	_jsii_.Get(
		c,
		"nodeFactory",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*INodeFactoryIface)(nil)).Elem(): reflect.TypeOf((*INodeFactory)(nil)).Elem(),
		},
	)
	return returns
}


// In what order to return constructs.
type ConstructOrder string

const (
	ConstructOrderPreorder ConstructOrder = "PREORDER"
	ConstructOrderPostorder ConstructOrder = "POSTORDER"
)

// DependencyIface is the public interface for the custom type Dependency
type DependencyIface interface {
	GetSource() IConstructIface
	GetTarget() IConstructIface
}

// A single dependency.
// Struct proxy
type Dependency struct {
	// Source the dependency.
	Source IConstructIface `json:"source"`
	// Target of the dependency.
	Target IConstructIface `json:"target"`
}

func (d *Dependency) GetSource() IConstructIface {
	var returns IConstructIface
	_jsii_.Get(
		d,
		"source",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IConstructIface)(nil)).Elem(): reflect.TypeOf((*IConstruct)(nil)).Elem(),
		},
	)
	return returns
}

func (d *Dependency) GetTarget() IConstructIface {
	var returns IConstructIface
	_jsii_.Get(
		d,
		"target",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IConstructIface)(nil)).Elem(): reflect.TypeOf((*IConstruct)(nil)).Elem(),
		},
	)
	return returns
}


// Represents an Aspect.
type IAspectIface interface {
	// All aspects can visit an IConstruct.
	Visit(node IConstructIface)
}

type IAspect struct {}

func (i *IAspect) Visit(node IConstructIface) {
	var returns interface{}
	_jsii_.Invoke(
		i,
		"visit",
		[]interface{}{node},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// Represents a construct.
type IConstructIface interface {
}

type IConstruct struct {}

// A factory for attaching `Node`s to the construct.
type INodeFactoryIface interface {
	// Returns a new `Node` associated with `host`.
	CreateNode(host ConstructIface, scope IConstructIface, id string) NodeIface
}

type INodeFactory struct {}

func (i *INodeFactory) CreateNode(host ConstructIface, scope IConstructIface, id string) NodeIface {
	var returns NodeIface
	_jsii_.Invoke(
		i,
		"createNode",
		[]interface{}{host, scope, id},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NodeIface)(nil)).Elem(): reflect.TypeOf((*Node)(nil)).Elem(),
		},
	)
	return returns
}

// Represents a single session of synthesis.
//
// Passed into `construct.onSynthesize()` methods.
type ISynthesisSessionIface interface {
	// The output directory for this synthesis session.
	GetOutdir() string
}

type ISynthesisSession struct {}

func (i *ISynthesisSession) GetOutdir() string {
	var returns string
	_jsii_.Get(
		i,
		"outdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

// Implement this interface in order for the construct to be able to validate itself.
type IValidationIface interface {
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if there the construct is valid.
	Validate() []string
}

type IValidation struct {}

func (i *IValidation) Validate() []string {
	var returns []string
	_jsii_.Invoke(
		i,
		"validate",
		[]interface{}{},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

// MetadataEntryIface is the public interface for the custom type MetadataEntry
type MetadataEntryIface interface {
	GetData() interface{}
	GetType() string
	GetTrace() []string
}

// An entry in the construct metadata table.
// Struct proxy
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

func (m *MetadataEntry) GetData() interface{} {
	var returns interface{}
	_jsii_.Get(
		m,
		"data",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (m *MetadataEntry) GetType() string {
	var returns string
	_jsii_.Get(
		m,
		"type",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (m *MetadataEntry) GetTrace() []string {
	var returns []string
	_jsii_.Get(
		m,
		"trace",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}


// Class interface
type NodeIface interface {
	GetAddr() string
	GetChildren() []IConstructIface
	GetDependencies() []DependencyIface
	GetId() string
	GetLocked() bool
	GetMetadata() []MetadataEntryIface
	GetPath() string
	GetRoot() IConstructIface
	GetScopes() []IConstructIface
	GetUniqueId() string
	GetScope() IConstructIface
	GetDefaultChild() IConstructIface
	SetDefaultChild(val IConstructIface)
	AddDependency(dependencies IConstructIface)
	AddError(message string)
	AddInfo(message string)
	AddMetadata(type_ string, data interface{}, fromFunction interface{})
	AddValidation(validation IValidationIface)
	AddWarning(message string)
	ApplyAspect(aspect IAspectIface)
	FindAll(order ConstructOrder) []IConstructIface
	FindChild(id string) IConstructIface
	Prepare()
	SetContext(key string, value interface{})
	Synthesize(options SynthesisOptionsIface)
	TryFindChild(id string) IConstructIface
	TryGetContext(key string) interface{}
	TryRemoveChild(childName string) bool
	Validate() []ValidationErrorIface
}

// Represents the construct node in the scope tree.
// Struct proxy
type Node struct {
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
	// TODO: EXAMPLE
	//
	Addr string `json:"addr"`
	// All direct children of this construct.
	Children []IConstructIface `json:"children"`
	// Return all dependencies registered on this node or any of its children.
	Dependencies []DependencyIface `json:"dependencies"`
	// The id of this construct within the current scope.
	//
	// This is a a scope-unique id. To obtain an app-unique id for this construct, use `uniqueId`.
	Id string `json:"id"`
	// Returns true if this construct or the scopes in which it is defined are locked.
	Locked bool `json:"locked"`
	// An immutable array of metadata objects associated with this construct.
	//
	// This can be used, for example, to implement support for deprecation notices, source mapping, etc.
	Metadata []MetadataEntryIface `json:"metadata"`
	// The full, absolute path of this construct in the tree.
	//
	// Components are separated by '/'.
	Path string `json:"path"`
	// Returns the root of the construct tree.
	//
	// Returns: The root of the construct tree.
	Root IConstructIface `json:"root"`
	// All parent scopes of this construct.
	//
	// Returns: a list of parent scopes. The last element in the list will always
	// be the current construct and the first element will be the root of the
	// tree.
	Scopes []IConstructIface `json:"scopes"`
	// A tree-global unique alphanumeric identifier for this construct.
	//
	// Includes
	// all components of the tree.
	// Deprecated: please avoid using this property and use `uid` instead. This
	// algorithm uses MD5, which is not FIPS-complient and also excludes the
	// identity of the root construct from the calculation.
	UniqueId string `json:"uniqueId"`
	// Returns the scope in which this construct is defined.
	//
	// The value is `undefined` at the root of the construct scope tree.
	Scope IConstructIface `json:"scope"`
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
	// Returns: a construct or undefined if there is no default child
	DefaultChild IConstructIface `json:"defaultChild"`
}

func (n *Node) GetAddr() string {
	var returns string
	_jsii_.Get(
		n,
		"addr",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *Node) GetChildren() []IConstructIface {
	var returns []IConstructIface
	_jsii_.Get(
		n,
		"children",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IConstructIface)(nil)).Elem(): reflect.TypeOf((*IConstruct)(nil)).Elem(),
		},
	)
	return returns
}

func (n *Node) GetDependencies() []DependencyIface {
	var returns []DependencyIface
	_jsii_.Get(
		n,
		"dependencies",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*DependencyIface)(nil)).Elem(): reflect.TypeOf((*Dependency)(nil)).Elem(),
		},
	)
	return returns
}

func (n *Node) GetId() string {
	var returns string
	_jsii_.Get(
		n,
		"id",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *Node) GetLocked() bool {
	var returns bool
	_jsii_.Get(
		n,
		"locked",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *Node) GetMetadata() []MetadataEntryIface {
	var returns []MetadataEntryIface
	_jsii_.Get(
		n,
		"metadata",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*MetadataEntryIface)(nil)).Elem(): reflect.TypeOf((*MetadataEntry)(nil)).Elem(),
		},
	)
	return returns
}

func (n *Node) GetPath() string {
	var returns string
	_jsii_.Get(
		n,
		"path",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *Node) GetRoot() IConstructIface {
	var returns IConstructIface
	_jsii_.Get(
		n,
		"root",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IConstructIface)(nil)).Elem(): reflect.TypeOf((*IConstruct)(nil)).Elem(),
		},
	)
	return returns
}

func (n *Node) GetScopes() []IConstructIface {
	var returns []IConstructIface
	_jsii_.Get(
		n,
		"scopes",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IConstructIface)(nil)).Elem(): reflect.TypeOf((*IConstruct)(nil)).Elem(),
		},
	)
	return returns
}

func (n *Node) GetUniqueId() string {
	var returns string
	_jsii_.Get(
		n,
		"uniqueId",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *Node) GetScope() IConstructIface {
	var returns IConstructIface
	_jsii_.Get(
		n,
		"scope",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IConstructIface)(nil)).Elem(): reflect.TypeOf((*IConstruct)(nil)).Elem(),
		},
	)
	return returns
}

func (n *Node) GetDefaultChild() IConstructIface {
	var returns IConstructIface
	_jsii_.Get(
		n,
		"defaultChild",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IConstructIface)(nil)).Elem(): reflect.TypeOf((*IConstruct)(nil)).Elem(),
		},
	)
	return returns
}


func NewNode(host ConstructIface, scope IConstructIface, id string) NodeIface {
	_init_.Initialize()
	self := Node{}
	_jsii_.Create(
		"constructs.Node",
		[]interface{}{host, scope, id},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (n *Node) SetDefaultChild(val IConstructIface) {
	_jsii_.Set(
		n,
		"defaultChild",
		val,
	)
}

func Node_Of(construct IConstructIface) NodeIface {
	_init_.Initialize()
	var returns NodeIface
	_jsii_.InvokeStatic(
		"constructs.Node",
		"of",
		[]interface{}{construct},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*NodeIface)(nil)).Elem(): reflect.TypeOf((*Node)(nil)).Elem(),
		},
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
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *Node) AddDependency(dependencies IConstructIface) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addDependency",
		[]interface{}{dependencies},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *Node) AddError(message string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addError",
		[]interface{}{message},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *Node) AddInfo(message string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addInfo",
		[]interface{}{message},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *Node) AddMetadata(type_ string, data interface{}, fromFunction interface{}) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addMetadata",
		[]interface{}{type_, data, fromFunction},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *Node) AddValidation(validation IValidationIface) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addValidation",
		[]interface{}{validation},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *Node) AddWarning(message string) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"addWarning",
		[]interface{}{message},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *Node) ApplyAspect(aspect IAspectIface) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"applyAspect",
		[]interface{}{aspect},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *Node) FindAll(order ConstructOrder) []IConstructIface {
	var returns []IConstructIface
	_jsii_.Invoke(
		n,
		"findAll",
		[]interface{}{order},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IConstructIface)(nil)).Elem(): reflect.TypeOf((*IConstruct)(nil)).Elem(),
		},
	)
	return returns
}

func (n *Node) FindChild(id string) IConstructIface {
	var returns IConstructIface
	_jsii_.Invoke(
		n,
		"findChild",
		[]interface{}{id},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IConstructIface)(nil)).Elem(): reflect.TypeOf((*IConstruct)(nil)).Elem(),
		},
	)
	return returns
}

func (n *Node) Prepare() {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"prepare",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *Node) SetContext(key string, value interface{}) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"setContext",
		[]interface{}{key, value},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *Node) Synthesize(options SynthesisOptionsIface) {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"synthesize",
		[]interface{}{options},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (n *Node) TryFindChild(id string) IConstructIface {
	var returns IConstructIface
	_jsii_.Invoke(
		n,
		"tryFindChild",
		[]interface{}{id},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IConstructIface)(nil)).Elem(): reflect.TypeOf((*IConstruct)(nil)).Elem(),
		},
	)
	return returns
}

func (n *Node) TryGetContext(key string) interface{} {
	var returns interface{}
	_jsii_.Invoke(
		n,
		"tryGetContext",
		[]interface{}{key},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *Node) TryRemoveChild(childName string) bool {
	var returns bool
	_jsii_.Invoke(
		n,
		"tryRemoveChild",
		[]interface{}{childName},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (n *Node) Validate() []ValidationErrorIface {
	var returns []ValidationErrorIface
	_jsii_.Invoke(
		n,
		"validate",
		[]interface{}{},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ValidationErrorIface)(nil)).Elem(): reflect.TypeOf((*ValidationError)(nil)).Elem(),
		},
	)
	return returns
}

// SynthesisOptionsIface is the public interface for the custom type SynthesisOptions
type SynthesisOptionsIface interface {
	GetOutdir() string
	GetSessionContext() map[string]interface{}
	GetSkipValidation() bool
}

// Options for synthesis.
// Struct proxy
type SynthesisOptions struct {
	// The output directory into which to synthesize the cloud assembly.
	Outdir string `json:"outdir"`
	// Additional context passed into the synthesis session object when `construct.synth` is called.
	SessionContext map[string]interface{} `json:"sessionContext"`
	// Whether synthesis should skip the validation phase.
	SkipValidation bool `json:"skipValidation"`
}

func (s *SynthesisOptions) GetOutdir() string {
	var returns string
	_jsii_.Get(
		s,
		"outdir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (s *SynthesisOptions) GetSessionContext() map[string]interface{} {
	var returns map[string]interface{}
	_jsii_.Get(
		s,
		"sessionContext",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*interface{})(nil)).Elem(): reflect.TypeOf((*interface{})(nil)).Elem(),
		},
	)
	return returns
}

func (s *SynthesisOptions) GetSkipValidation() bool {
	var returns bool
	_jsii_.Get(
		s,
		"skipValidation",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// ValidationErrorIface is the public interface for the custom type ValidationError
type ValidationErrorIface interface {
	GetMessage() string
	GetSource() ConstructIface
}

// An error returned during the validation phase.
// Struct proxy
type ValidationError struct {
	// The error message.
	Message string `json:"message"`
	// The construct which emitted the error.
	Source ConstructIface `json:"source"`
}

func (v *ValidationError) GetMessage() string {
	var returns string
	_jsii_.Get(
		v,
		"message",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (v *ValidationError) GetSource() ConstructIface {
	var returns ConstructIface
	_jsii_.Get(
		v,
		"source",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ConstructIface)(nil)).Elem(): reflect.TypeOf((*Construct)(nil)).Elem(),
		},
	)
	return returns
}


