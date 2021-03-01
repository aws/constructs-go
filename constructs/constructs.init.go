package constructs

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go"
)

func init() {
	_jsii_.RegisterClass(
		"constructs.Construct",
		reflect.TypeOf((*Construct)(nil)).Elem(),
		func() interface{} {
			c := construct{}
			_jsii_.InitJsiiProxy(&c.iConstruct)
			return &c
		},
	)
	_jsii_.RegisterClass(
		"constructs.ConstructMetadata",
		reflect.TypeOf((*ConstructMetadata)(nil)).Elem(),
		func() interface{} {
			return &constructMetadata{}
		},
	)
	_jsii_.RegisterStruct(
		"constructs.ConstructOptions",
		reflect.TypeOf((*ConstructOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"constructs.ConstructOrder",
		reflect.TypeOf((*ConstructOrder)(nil)).Elem(),
		map[string]interface{}{
			"PREORDER": ConstructOrder_PREORDER,
			"POSTORDER": ConstructOrder_POSTORDER,
		},
	)
	_jsii_.RegisterStruct(
		"constructs.Dependency",
		reflect.TypeOf((*Dependency)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"constructs.IAspect",
		reflect.TypeOf((*IAspect)(nil)).Elem(),
		func() interface{} {
			return &iAspect{}
		},
	)
	_jsii_.RegisterInterface(
		"constructs.IConstruct",
		reflect.TypeOf((*IConstruct)(nil)).Elem(),
		func() interface{} {
			return &iConstruct{}
		},
	)
	_jsii_.RegisterInterface(
		"constructs.INodeFactory",
		reflect.TypeOf((*INodeFactory)(nil)).Elem(),
		func() interface{} {
			return &iNodeFactory{}
		},
	)
	_jsii_.RegisterInterface(
		"constructs.ISynthesisSession",
		reflect.TypeOf((*ISynthesisSession)(nil)).Elem(),
		func() interface{} {
			return &iSynthesisSession{}
		},
	)
	_jsii_.RegisterInterface(
		"constructs.IValidation",
		reflect.TypeOf((*IValidation)(nil)).Elem(),
		func() interface{} {
			return &iValidation{}
		},
	)
	_jsii_.RegisterStruct(
		"constructs.MetadataEntry",
		reflect.TypeOf((*MetadataEntry)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"constructs.Node",
		reflect.TypeOf((*Node)(nil)).Elem(),
		func() interface{} {
			return &node{}
		},
	)
	_jsii_.RegisterStruct(
		"constructs.SynthesisOptions",
		reflect.TypeOf((*SynthesisOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"constructs.ValidationError",
		reflect.TypeOf((*ValidationError)(nil)).Elem(),
	)
}
