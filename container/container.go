package container

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/drgomesp/cargo"
	"github.com/drgomesp/cargo/definition"
)

// Container for dependency injection
type Container struct {
	definitions map[string]*definition.Definition
	services    map[string]interface{}
}

// NewContainer instance
func NewContainer() *Container {
	return &Container{
		definitions: make(map[string]*definition.Definition, 0),
		services:    make(map[string]interface{}, 0),
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

// Set a new service
func (c *Container) Set(id string, arg interface{}) (err error) {
	if _, ok := c.definitions[id]; ok {
		err = cargo.NewError(fmt.Sprintf("Definition for \"%s\" already exists", id))
		return
	}

	if c.definitions[id], err = definition.NewDefinition(arg); err != nil {
		err = cargo.NewError("Could not create definition")
	}

	c.services[id] = arg

	return
}

// Get a service
func (c *Container) Get(id string) (service interface{}, err error) {
	for i := 0; i < 2; i++ {
		if s, ok := c.services[id]; ok {
			service = s
			return
		}

		if def, ok := c.definitions[id]; ok {
			if service, err = createServiceFromDefinition(def); err == nil {
				c.services[id] = service
				return
			}
		}

		id = strings.ToLower(id)
	}

	err = cargo.NewError(fmt.Sprintf("No service \"%s\" was found", id))
	return
}

func createServiceFromDefinition(def *definition.Definition) (service interface{}, err error) {
	var ret reflect.Value

	if def.Constructor.IsValid() {
		if len(def.Arguments) > 0 {
			args := make([]reflect.Value, len(def.Arguments))

			for i, arg := range def.Arguments {
				args[i] = reflect.ValueOf(arg.Value)
			}

			ret = def.Constructor.Call(args)[0]

			for i, arg := range def.Arguments {
				field := (&ret).Elem().Field(i)

				if field.IsValid() && field.CanSet() {
					field.Set(reflect.ValueOf(arg.Value))
				}
			}
		} else {
			ret = def.Constructor.Call(make([]reflect.Value, 0))[0]
		}

		if len(def.MethodCalls) > 0 {
			for _, method := range def.MethodCalls {
				if m, ok := ret.Type().MethodByName(method.Name); ok {
					if m.Func.Type().NumIn() > 0 {
						numArgs := m.Func.Type().NumIn() - 1

						if len(method.Args) != numArgs {
							err = cargo.NewError("Method \"%s\" expects arguments")
							return
						}

						args := make([]reflect.Value, numArgs)

						for i, arg := range method.Args {
							args[i] = reflect.ValueOf(arg.Value)
						}

						args = append([]reflect.Value{ret}, args...)

						m.Func.Call(args)
					} else {
						m.Func.Call([]reflect.Value{ret})
					}
				}
			}
		}

		service = ret.Interface()
	}

	return
}
