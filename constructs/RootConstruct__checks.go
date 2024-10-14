//go:build !no_runtime_type_checking

package constructs

import (
	"fmt"
)

func validateRootConstruct_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

