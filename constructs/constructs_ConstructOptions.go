// A programming model for composable configuration
package constructs


// Options for creating constructs.
type ConstructOptions struct {
	// A factory for attaching `Node`s to the construct.
	NodeFactory INodeFactory `field:"optional" json:"nodeFactory" yaml:"nodeFactory"`
}

