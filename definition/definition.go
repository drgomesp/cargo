package definition

import (
	"reflect"

	"github.com/drgomesp/cargo"
	"github.com/drgomesp/cargo/argument"
	"github.com/drgomesp/cargo/method"
)

// Definition of a service or an argument
type Definition struct {
	Arguments   []*argument.Argument
	MethodCalls []*method.Method
	Constructor reflect.Value
	Type        reflect.Type
}

// NewDefinition based on factory functions, composite ˝ or pointers
func NewDefinition(arg interface{}, args ...interface{}) (def *Definition, err error) {
	switch reflect.TypeOf(arg).Kind() {
	case reflect.Func:
		if constructor, err := createFromConstructorFunction(reflect.ValueOf(arg)); nil == err {
			def = constructor
		}
	case reflect.Ptr:
		if constructor, err := createFromPointer(&arg); nil == err {
			def = constructor
		}
	default:
		err = cargo.NewError("Could not create definition")
	}

	return
}

// AddArguments to the definition
func (d *Definition) AddArguments(arg ...*argument.Argument) *Definition {
	d.Arguments = append(d.Arguments, arg...)
	return d
}

// AddMethodCall to the definition
func (d *Definition) AddMethodCall(method *method.Method) *Definition {
	d.MethodCalls = append(d.MethodCalls, method)
	return d
}

func createFromConstructorFunction(fn reflect.Value) (def *Definition, err error) {
	var returnType reflect.Type

	constructor := reflect.MakeFunc(fn.Type(), func(in []reflect.Value) []reflect.Value {
		returnType = reflect.TypeOf(fn.Interface()).Out(0)

		return []reflect.Value{reflect.New(returnType.Elem())}
	})

	def = &Definition{
		Arguments:   make([]*argument.Argument, 0),
		MethodCalls: make([]*method.Method, 0),
		Constructor: constructor,
		Type:        reflect.TypeOf(constructor.Interface()).Out(0),
	}

	return
}

func createFromPointer(ptr interface{}, args ...interface{}) (def *Definition, err error) {
	def = &Definition{
		Arguments:   make([]*argument.Argument, 0),
		MethodCalls: make([]*method.Method, 0),
		Type:        reflect.TypeOf(ptr),
	}

	return
}
