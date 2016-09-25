package definition

// Reference to a service
type Reference struct {
	Identifier string
}

// NewReference of a service
func NewReference(id string) Reference {
	return Reference{id}
}
