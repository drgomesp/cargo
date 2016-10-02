package definition

import (
	"reflect"

	"github.com/drgomesp/cargo/argument"
	"github.com/drgomesp/cargo/method"
)

// Interface providing the basic API for a definition
type Interface interface {
	AddArguments(arg ...*argument.Argument) Interface
	AddMethodCall(method *method.Method) Interface

	Arguments() []*argument.Argument
	MethodCalls() []*method.Method
	Constructor() reflect.Value
	Type() reflect.Type
}
