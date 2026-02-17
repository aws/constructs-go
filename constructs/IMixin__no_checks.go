//go:build no_runtime_type_checking

package constructs

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IMixin) validateApplyToParameters(construct IConstruct) error {
	return nil
}

func (i *jsiiProxy_IMixin) validateSupportsParameters(construct IConstruct) error {
	return nil
}

