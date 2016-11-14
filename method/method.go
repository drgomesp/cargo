package method

import "github.com/drgomesp/cargo/argument"

// Method represents a method for a service definition
type Method struct {
	Name string
	Args []*argument.Argument
}

// New method reference
func New(name string, args ...interface{}) *Method {
	arguments := make([]*argument.Argument, len(args))

	for i, arg := range args {
		arguments[i] = argument.New(arg)
	}

	return &Method{
		Name: name,
		Args: arguments,
	}
}
