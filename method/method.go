package method

import "github.com/drgomesp/cargo/argument"

// Method represents a method for a service definition
type Method struct {
	Name string
	Args []*argument.Argument
}

// NewMethod reference
func NewMethod(name string, args ...interface{}) *Method {
	arguments := make([]*argument.Argument, len(args))

	for i, arg := range args {
		arguments[i] = argument.NewArgument(arg)
	}

	return &Method{
		Name: name,
		Args: arguments,
	}
}
