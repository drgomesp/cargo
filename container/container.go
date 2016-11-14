package container

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/drgomesp/cargo/definition"
	"github.com/drgomesp/cargo/reference"
)

// Container for dependency injection
type Container struct {
	definitions map[string]definition.Interface
	services    map[string]interface{}
}

// New continer instance
func New() *Container {
	return &Container{
		definitions: make(map[string]definition.Interface, 0),
		services:    make(map[string]interface{}, 0),
	}
}

// Register a new service definition
func (c *Container) Register(id string, arg interface{}) (def definition.Interface, err error) {
	if _, ok := c.definitions[id]; ok {
		err = fmt.Errorf(`Definition for "%s" already exists`, id)
		return
	}

	def, err = definition.New(arg)
	c.definitions[id] = def

	return
}

// Set a new service
func (c *Container) Set(id string, arg interface{}) (err error) {
	if _, ok := c.definitions[id]; ok {
		err = fmt.Errorf(`Definition for "%s" already exists`, id)
		return
	}

	if c.definitions[id], err = definition.New(arg); err != nil {
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
			service, err = c.createService(def)

			if err != nil {
				return
			}

			c.services[id] = service
			return
		}

		id = strings.ToLower(id)
	}

	err = fmt.Errorf(`No service "%s" was found`, id)
	return
}

// MustGet is a wrapper for Get that panics if service was not found
func (c *Container) MustGet(id string) interface{} {
	if service, err := c.Get(id); err != nil {
		panic(err)
	} else {
		return service
	}
}

func (c *Container) createService(def definition.Interface) (service interface{}, err error) {
	obj, _ := c.callConstructor(def)

	if len(def.MethodCalls()) > 0 {
		if err = callMethods(def, &obj); err != nil {
			return
		}
	}

	return obj.Interface(), nil
}

func (c *Container) callConstructor(def definition.Interface) (obj reflect.Value, err error) {
	if len(def.Arguments()) > 0 {
		args := make([]reflect.Value, len(def.Arguments()))

		for i, arg := range def.Arguments() {
			if reference, ok := arg.(reference.Interface); ok {
				var found interface{}
				found, err = c.Get(reference.Identifier())

				if err != nil {
					return
				}

				args[i] = reflect.ValueOf(found)
			} else {
				args[i] = reflect.ValueOf(arg.Value())
			}
		}

		obj = def.Constructor().Call(args)[0]

		for i, arg := range args {
			field := obj.Elem().Field(i)

			if field.IsValid() && field.CanSet() {
				field.Set(arg)
			}
		}
	} else {
		obj = def.Constructor().Call(make([]reflect.Value, 0))[0]
	}

	return obj, nil
}

func callMethods(def definition.Interface, obj *reflect.Value) (err error) {
	for _, method := range def.MethodCalls() {
		if m, ok := obj.Type().MethodByName(method.Name); ok {
			if m.Func.Type().NumIn() > 0 {
				numArgs := m.Func.Type().NumIn() - 1

				if len(method.Args) != numArgs {
					err = fmt.Errorf(`Method "%s" expects %d arguments`, method.Name, numArgs)
					return
				}

				args := make([]reflect.Value, numArgs)

				for i, arg := range method.Args {
					args[i] = reflect.ValueOf(arg.Value())
				}

				args = append([]reflect.Value{*obj}, args...)

				m.Func.Call(args)
			} else {
				m.Func.Call(nil)
			}
		}
	}

	return
}
