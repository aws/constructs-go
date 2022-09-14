//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// A programming model for composable configuration
package constructs

import (
	"fmt"
)

func (i *jsiiProxy_IAspect) validateVisitParameters(node IConstruct) error {
	if node == nil {
		return fmt.Errorf("parameter node is required, but nil was provided")
	}

	return nil
}

