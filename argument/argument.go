package argument

// Argument that can be passed into a service constructor
type Argument struct {
	Value interface{}
}

// New argument to be used in definitions of services
func New(arg interface{}) *Argument {
	return &Argument{arg}
}
