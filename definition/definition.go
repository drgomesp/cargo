package definition

// Definition represents a service definition
type Definition struct {
	Arguments   []interface{}
	Factory     string
	ID          string
	Lazy        bool
	MethodCalls []string
	Tags        []string
	Type        string
}

// NewDefinition creates a new definition
func NewDefinition(id string, t string, args []interface{}) Definition {
	return Definition{
		ID:        id,
		Type:      t,
		Arguments: args,
	}
}
