package definition

// Argument that can be passed into a service constructor
type Argument struct {
	Value interface{}
}

// NewArgument to be used in definitions of services
func NewArgument(arg interface{}) *Argument {
	return &Argument{arg}
}
