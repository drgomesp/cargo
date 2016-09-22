package cargo

// Container is a service that handles instances
type Container struct {
	singletons map[string]*interface{}
}

// Get an instance by its identifier
func (c *Container) Get(identifier string) (singleton *interface{}, err error) {
	if singleton, ok := c.singletons[identifier]; !ok {
		return singleton, nil
	}

	return nil, Error{}
}
