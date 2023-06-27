package constructs

import (
	_init_ "github.com/aws/constructs-go/constructs/v3/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

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

