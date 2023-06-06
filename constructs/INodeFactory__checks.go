//go:build !no_runtime_type_checking

package constructs

import (
	"fmt"
)

func (i *jsiiProxy_INodeFactory) validateCreateNodeParameters(host Construct, scope IConstruct, id *string) error {
	if host == nil {
		return fmt.Errorf("parameter host is required, but nil was provided")
	}

	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}

