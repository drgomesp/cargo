package container

import (
	"fmt"
	"reflect"

	"github.com/drgomesp/cargo"
	"github.com/drgomesp/cargo/definition"
)

// Container for dependency injection
type Container struct {
	definitions map[string]*definition.Definition
	services    map[string]*interface{}
}

// NewContainer instance
func NewContainer() *Container {
	return &Container{
		definitions: make(map[string]*definition.Definition, 0),
	}
}

// Register a new service definition
func (c *Container) Register(id string, arg interface{}) (def *definition.Definition, err error) {
	if _, ok := c.definitions[id]; ok {
		err = cargo.NewError(fmt.Sprintf("Definition for \"%s\" already exists", id))
		return
	}

	def, err = definition.NewDefinition(arg)
	c.definitions[id] = def

	return
}

// Get a service
func (c *Container) Get(id string) (service *interface{}, err error) {
	if s, ok := c.services[id]; ok {
		service = s
		return
	}

	if def, ok := c.definitions[id]; ok {
		if service, err = createServiceFromDefinition(def); err != nil {
			err = cargo.NewError(fmt.Sprintf("No service \"%s\" was found", id))
		}
	}

	return
}

func createServiceFromDefinition(def *definition.Definition) (service *interface{}, err error) {
	if def.Constructor.Kind() == reflect.Func {
		ret := def.Constructor.Call(nil)[0]
		ptr := ret.Interface()
		service = &ptr
		return service, nil
	}

	return
}
