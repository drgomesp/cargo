package reference

import "github.com/drgomesp/cargo/argument"

// Interface that defines a service reference
type Interface interface {
	argument.Interface
	Identifier() string
}
