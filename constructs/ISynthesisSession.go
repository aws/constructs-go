package constructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

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

