package container

import (
	"reflect"
	"strings"

	"github.com/drgomesp/cargo"
)

const (
	// ServiceNotFound represents a not found service for the given id
	ServiceNotFound = -1

	// ServiceInvalid represents an invalid interface value was provided
	ServiceInvalid = -2

	// ServiceNotRegistered represents the register process failed unexpectedly
	ServiceNotRegistered = -4
)

// Container is a service that handles instances
type Container struct {
	aliases  map[string]reflect.Value
	services map[string]reflect.Value
}

// NewContainer creates a new container
func NewContainer() *Container {
	return &Container{
		aliases:  make(map[string]reflect.Value),
		services: make(map[string]reflect.Value),
	}
}

// Get an instance by its identifier
func (c *Container) Get(id string) (service interface{}, err error) {
	/* Since identifiers are always in lower case and this method can be
	   called thousands of times during the execution of a program, a first
	   attempt to retrieve the service is made without ToLower() unless necessary */
	for i := 0; i < 2; i++ {
		if service, ok := c.aliases[id]; ok {
			return service.Interface(), nil
		}

		if service, ok := c.services[id]; ok {
			return service.Interface(), nil
		}

		id = strings.ToLower(id)
	}

	return nil, cargo.NewError(ServiceNotFound, "Service \"%s\" not found", id)
}

// Register an instance with an identifier
func (c *Container) Register(id string, service interface{}) (err error) {
	if service == nil {
		return cargo.NewError(ServiceInvalid, "Service \"%s\" must not be nil", id)
	}

	id = strings.ToLower(id)

	if _, ok := c.aliases[id]; ok {
		delete(c.aliases, id)
	}

	s := reflect.ValueOf(service)

	if s.Kind() == reflect.Ptr {
		s = s.Elem()
		c.services[id] = s
		return
	}

	return cargo.NewError(ServiceNotRegistered, "Service \"%s\" was not registered", id)
}
