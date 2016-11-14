package definition

import (
	"fmt"
	"reflect"

	"github.com/drgomesp/cargo/argument"
	"github.com/drgomesp/cargo/method"
)

// Definition of a service or an argument
type Definition struct {
	arguments   []argument.Interface
	methodCalls []*method.Method
	constructor reflect.Value
	t           reflect.Type
}

// New definition based on factory functions or pointers
func New(arg interface{}, args ...interface{}) (def Interface, err error) {
	switch reflect.TypeOf(arg).Kind() {
	case reflect.Func:
		if reflect.TypeOf(arg).NumOut() == 0 {
			err = fmt.Errorf("Constructor function must have a return type")
			return
		}

		if constructor, err := createFromConstructorFunction(reflect.ValueOf(arg)); nil == err {
			def = constructor
		}
	case reflect.Ptr:
		if constructor, err := createFromPointer(&arg); nil == err {
			def = constructor
		}
	default:
		err = fmt.Errorf("A definition must be created from a pointer to a struct or a constructor function")
	}

	return
}

// AddArguments to the definition
func (d *Definition) AddArguments(arg ...argument.Interface) Interface {
	d.arguments = append(d.arguments, arg...)
	return Interface(d)
}

// AddMethodCall to the definition
func (d *Definition) AddMethodCall(method *method.Method) Interface {
	d.methodCalls = append(d.methodCalls, method)
	return Interface(d)
}

// Arguments of the definition
func (d *Definition) Arguments() []argument.Interface {
	return d.arguments
}

// Method calls of the definition
func (d *Definition) MethodCalls() []*method.Method {
	return d.methodCalls
}

// Constructor for the definition
func (d *Definition) Constructor() reflect.Value {
	return d.constructor
}

// Type for the definition
func (d *Definition) Type() reflect.Type {
	return d.t
}

func createFromConstructorFunction(fn reflect.Value) (def Interface, err error) {
	var returnType reflect.Type

	constructor := reflect.MakeFunc(fn.Type(), func(in []reflect.Value) []reflect.Value {
		returnType = reflect.TypeOf(fn.Interface()).Out(0)

		return []reflect.Value{reflect.New(returnType.Elem())}
	})

	def = &Definition{
		arguments:   make([]argument.Interface, 0),
		methodCalls: make([]*method.Method, 0),
		constructor: constructor,
		t:           reflect.TypeOf(constructor.Interface()).Out(0),
	}

	return
}

func createFromPointer(ptr interface{}, args ...interface{}) (def Interface, err error) {
	def = &Definition{
		arguments:   make([]argument.Interface, 0),
		methodCalls: make([]*method.Method, 0),
		t:           reflect.TypeOf(ptr),
	}

	return
}
