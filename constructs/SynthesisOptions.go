// A programming model for composable configuration
package constructs


// Options for synthesis.
type SynthesisOptions struct {
	// The output directory into which to synthesize the cloud assembly.
	Outdir *string `field:"required" json:"outdir" yaml:"outdir"`
	// Additional context passed into the synthesis session object when `construct.synth` is called.
	SessionContext *map[string]interface{} `field:"optional" json:"sessionContext" yaml:"sessionContext"`
	// Whether synthesis should skip the validation phase.
	SkipValidation *bool `field:"optional" json:"skipValidation" yaml:"skipValidation"`
}

