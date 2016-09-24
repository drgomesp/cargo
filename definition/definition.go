package definition

import "reflect"

// Definition represents a service definition
type Definition struct {
	Arguments []interface{}
	ID        string
	Type      reflect.Type
}

// NewDefinition creates a new definition
func NewDefinition(reference interface{}) Definition {
	var v reflect.Value
	var t reflect.Type

	v = reflect.ValueOf(reference)
	t = reflect.TypeOf(reference)

	if v.Kind() == reflect.Func {
		t = reflect.TypeOf(v.Call(nil)[0].Interface())
	}

	return Definition{
		Arguments: make([]interface{}, 0),
		Type:      t,
	}
}

// AddArguments to the definition
func (d *Definition) AddArguments(arg ...interface{}) {
	d.Arguments = append(d.Arguments, arg...)
}
