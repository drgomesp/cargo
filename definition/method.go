package definition

// Method represents a method for a service definition
type Method struct {
	Name string
	Args []*Argument
}

// NewMethod reference
func NewMethod(name string, args ...interface{}) *Method {
	arguments := make([]*Argument, len(args))

	for i, arg := range args {
		arguments[i] = NewArgument(arg)
	}

	return &Method{
		Name: name,
		Args: arguments,
	}
}
