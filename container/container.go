package container

import (
	"reflect"
	"strings"

	"github.com/drgomesp/cargo"
	"github.com/drgomesp/cargo/definition"
)

const (
	// ServiceNotFound represents a not found service for the given id
	ServiceNotFound = -1

	// ServiceInvalid represents an invalid interface value was provided
	ServiceInvalid = -2

	// ServiceNotRegistered represents the register process failed unexpectedly
	ServiceNotRegistered = -4

	// DefinitionConflicted represents that the definition was already previously created
	DefinitionConflicted = -1

	// DefinitionNotFound represents a not found definition for the given id
	DefinitionNotFound = -2
)

// Container is a service that handles instances
type Container struct {
	aliases     map[string]reflect.Value
	definitions map[string]definition.Definition
	services    map[string]reflect.Value
}

// NewContainer creates a new container
func NewContainer() *Container {
	return &Container{
		aliases:     make(map[string]reflect.Value),
		definitions: make(map[string]definition.Definition),
		services:    make(map[string]reflect.Value),
	}
}

// Get an instance by its identifier
func (c *Container) Get(id string) (service interface{}, err error) {
	/* Since identifiers are always in lower case and this method can be
	   called thousands of times during the execution of a program, a first
	   attempt to retrieve the service is made without ToLower() unless necessary */
	for i := 0; i < 2; i++ {
		if def, ok := c.definitions[id]; ok {
			return c.createService(id, def)
		}

		if service, ok := c.aliases[id]; ok {
			return service.Elem(), nil
		}

		if service, ok := c.services[id]; ok {
			return service.Elem(), nil
		}

		id = strings.ToLower(id)
	}

	return nil, cargo.NewError(ServiceNotFound, "Service \"%s\" was not found", id)
}

// Set an instance with an identifier
func (c *Container) Set(id string, service interface{}) (err error) {
	if service == nil {
		return cargo.NewError(ServiceInvalid, "Service \"%s\" must not be nil", id)
	}

	id = strings.ToLower(id)

	if _, ok := c.aliases[id]; ok {
		delete(c.aliases, id)
	}

	if def, ok := c.definitions[id]; ok {
		if s, err := c.createService(id, def); err != nil {
			c.services[id] = reflect.ValueOf(s)
			return nil
		}
	}

	def := definition.NewDefinition(service)
	c.definitions[id] = def

	if s, err := c.createService(id, def); err != nil {
		c.services[id] = reflect.ValueOf(s)
	}

	return cargo.NewError(ServiceNotRegistered, "Service \"%s\" was not registered", id)
}

// HasDefinition checks if a given definiition exists on the container
func (c *Container) HasDefinition(id string) bool {
	_, ok := c.definitions[id]
	return ok
}

// GetDefinition retrieves a definition from the container
func (c *Container) GetDefinition(id string) (def definition.Definition, err error) {
	for i := 0; i < 2; i++ {
		if def, ok := c.definitions[id]; ok {
			return def, nil
		}

		if def, ok := c.definitions[strings.ToLower(id)]; ok {
			return def, nil
		}
	}

	return def, cargo.NewError(DefinitionNotFound, "Definition \"%s\" was not found", id)
}

// Register a definition to the container
func (c *Container) Register(id string, def definition.Definition) (container *Container, rerr error) {
	if _, ok := c.definitions[id]; ok {
		return c, cargo.NewError(DefinitionConflicted, "Definition \"%s\" already exists", id)
	}

	c.definitions[id] = def
	return c, nil
}

func (c *Container) createService(id string, def definition.Definition) (service interface{}, err error) {
	if _, ok := c.definitions[id]; !ok {
		return
	}

	value := reflect.ValueOf(service)

	if !value.IsValid() {
		s := reflect.New(c.definitions[id].Type)
		service = reflect.Indirect(s).Interface()
		return
	}

	c.services[id] = value
	service = value

	return
}
