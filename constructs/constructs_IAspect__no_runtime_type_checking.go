//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// A programming model for composable configuration
package constructs

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IAspect) validateVisitParameters(node IConstruct) error {
	return nil
}

