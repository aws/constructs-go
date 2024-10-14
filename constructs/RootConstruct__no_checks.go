//go:build no_runtime_type_checking

package constructs

// Building without runtime type checking enabled, so all the below just return nil

func validateRootConstruct_IsConstructParameters(x interface{}) error {
	return nil
}

