// A programming model for composable configuration
package constructs


// An entry in the construct metadata table.
type MetadataEntry struct {
	// The data.
	Data interface{} `field:"required" json:"data" yaml:"data"`
	// The metadata entry type.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Stack trace.
	//
	// Can be omitted by setting the context key
	// `ConstructMetadata.DISABLE_STACK_TRACE_IN_METADATA` to 1.
	Trace *[]*string `field:"optional" json:"trace" yaml:"trace"`
}

