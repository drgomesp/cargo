package definition

import "reflect"

// Definition of a service or an argument
type Definition struct {
	Arguments   []Argument
	MethodCalls []reflect.Value
	Constructor reflect.Value
	Type        reflect.Type
}

// NewDefinition based on factory functions, composite ˝ or pointers
func NewDefinition(arg interface{}, args ...interface{}) (def *Definition, err error) {
	switch reflect.TypeOf(arg).Kind() {
	case reflect.Func:
		if constructor, err := createFromConstructorFunction(reflect.ValueOf(arg)); nil == err {
			def = constructor
			break
		}
	default:
		if constructor, err := createFromPointer(&arg); nil == err {
			def = constructor
			break
		}
	}

	return
}

// AddArgument to the definition
func (d *Definition) AddArgument(arg Argument) *Definition {
	d.Arguments = append(d.Arguments, arg)
	return d
}

func createFromConstructorFunction(fn reflect.Value) (def *Definition, err error) {
	var returnType reflect.Type

	constructor := reflect.MakeFunc(fn.Type(), func(in []reflect.Value) []reflect.Value {
		returnType = reflect.TypeOf(fn.Interface()).Out(0)
		return []reflect.Value{reflect.New(returnType).Elem()}
	})

	def = &Definition{
		Arguments:   make([]Argument, 0),
		Constructor: constructor,
		Type:        reflect.TypeOf(constructor.Interface()).Out(0),
	}

	return
}

func createFromPointer(ptr interface{}, args ...interface{}) (def *Definition, err error) {
	def = &Definition{
		Arguments: make([]Argument, 0),
		Type:      reflect.TypeOf(ptr),
	}

	return
}
