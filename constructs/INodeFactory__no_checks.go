//go:build no_runtime_type_checking

package constructs

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_INodeFactory) validateCreateNodeParameters(host Construct, scope IConstruct, id *string) error {
	return nil
}

