package argument

// Argument that can be passed into a service constructor
type Argument struct {
	value interface{}
}

// Value carried by the argument
func (a *Argument) Value() interface{} {
	return a.value
}

// New argument to be used in definitions of services
func New(arg interface{}) *Argument {
	return &Argument{arg}
}
