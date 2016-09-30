package definition

import (
	"github.com/drgomesp/cargo/method"
	"github.com/drgomesp/cargo/reference"
)

// Interface providing the basic API for a definition
type Interface interface {
	AddArguments(arg ...reference.Reference) *Definition
	AddMethodCall(method method.Method) *Definition
}
