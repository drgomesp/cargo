package definition

import (
	"fmt"
	"reflect"
)

// Definition represents a service definition
type Definition struct {
	Arguments []interface{}
	ID        string
	Type      reflect.Type
}

// NewDefinition creates a new definition
func NewDefinition(id string, fn interface{}) Definition {
	var v reflect.Value
	var t reflect.Type

	t = reflect.TypeOf(fn)
	v = reflect.ValueOf(fn)

	switch v.Kind() {
	case reflect.Ptr:
		v = v.Elem()
	case reflect.Func:
		fmt.Printf("%s => reflect.Func (%s) \n\n", id, t)

	}

	return Definition{
		Arguments: make([]interface{}, 0),
		ID:        id,
		Type:      t,
	}
}

// AddArgument to the definition
func (d *Definition) AddArgument(arg interface{}) {
	d.Arguments = append(d.Arguments, arg)
}
