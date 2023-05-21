//go:build no_runtime_type_checking

package constructs

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_Construct) validateOnSynthesizeParameters(session ISynthesisSession) error {
	return nil
}

func validateNewConstructParameters(scope Construct, id *string, options *ConstructOptions) error {
	return nil
}

