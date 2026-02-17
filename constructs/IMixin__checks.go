//go:build !no_runtime_type_checking

package constructs

import (
	"fmt"
)

func (i *jsiiProxy_IMixin) validateApplyToParameters(construct IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IMixin) validateSupportsParameters(construct IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

