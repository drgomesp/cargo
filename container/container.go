package container

import (
	"fmt"
	"reflect"
	"strings"

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
		err = fmt.Errorf("Definition for \"%s\" already exists", id)
		return
	}

	def, err = definition.NewDefinition(arg)
	c.definitions[id] = def

	return
}

// Set a new service
func (c *Container) Set(id string, arg interface{}) (err error) {
	if _, ok := c.definitions[id]; ok {
		err = fmt.Errorf("Definition for \"%s\" already exists", id)
		return
	}

	if c.definitions[id], err = definition.NewDefinition(arg); err != nil {
		err = fmt.Errorf("Could not create definition")
		return
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
			if service, err = createService(def); err == nil {
				c.services[id] = service
				return
			}
		}

		id = strings.ToLower(id)
	}

	err = fmt.Errorf("No service \"%s\" was found", id)
	return
}

func createService(def *definition.Definition) (service interface{}, err error) {
	if !def.Constructor.IsValid() {
		err = fmt.Errorf("Constructor is not valid")
		return
	}

	obj, err := callConstructor(def)

	if err != nil {
		return nil, err
	}

	if len(def.MethodCalls) > 0 {
		callMethods(def, &obj)
	}

	return obj.Interface(), nil
}

func callConstructor(def *definition.Definition) (obj reflect.Value, err error) {
	if len(def.Arguments) > 0 {
		args := make([]reflect.Value, len(def.Arguments))

		for i, arg := range def.Arguments {
			args[i] = reflect.ValueOf(arg.Value)
		}

		obj = def.Constructor.Call(args)[0]

		for i, arg := range def.Arguments {
			field := (obj).Elem().Field(i)

			if field.IsValid() && field.CanSet() {
				field.Set(reflect.ValueOf(arg.Value))
			}
		}
	} else {
		obj = def.Constructor.Call(make([]reflect.Value, 0))[0]
	}

	return obj, nil
}

func callMethods(def *definition.Definition, obj *reflect.Value) (err error) {
	for _, method := range def.MethodCalls {
		if m, ok := obj.Type().MethodByName(method.Name); ok {
			if m.Func.Type().NumIn() > 0 {
				numArgs := m.Func.Type().NumIn() - 1

				if len(method.Args) != numArgs {
					err = fmt.Errorf("Method \"%s\" expects arguments", method.Name)
					return
				}

				args := make([]reflect.Value, numArgs)

				for i, arg := range method.Args {
					args[i] = reflect.ValueOf(arg.Value)
				}

				args = append([]reflect.Value{*obj}, args...)

				m.Func.Call(args)
			} else {
				m.Func.Call([]reflect.Value{*obj})
			}
		}
	}

	return
}
