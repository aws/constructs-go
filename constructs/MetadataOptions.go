package constructs


// Options for `construct.addMetadata()`.
type MetadataOptions struct {
	// Include stack trace with metadata entry.
	// Default: false.
	//
	StackTrace *bool `field:"optional" json:"stackTrace" yaml:"stackTrace"`
	// The actual stack trace to be added to the metadata.
	//
	// If this
	// parameter is passed, the stackTrace parameter is ignored.
	StackTraceOverride *[]*string `field:"optional" json:"stackTraceOverride" yaml:"stackTraceOverride"`
	// A JavaScript function to begin tracing from.
	//
	// This option is ignored unless `stackTrace` is `true`.
	// Default: addMetadata().
	//
	TraceFromFunction interface{} `field:"optional" json:"traceFromFunction" yaml:"traceFromFunction"`
}

