package reference

// Reference to a service
type Reference struct {
	value      interface{}
	identifier string
}

// Value carried by the argument
func (r *Reference) Value() interface{} {
	return r.value
}

// Identifier of the referenced service
func (r *Reference) Identifier() string {
	return r.identifier
}

// New reference of a service
func New(id string) Reference {
	return Reference{
		identifier: id,
	}
}
