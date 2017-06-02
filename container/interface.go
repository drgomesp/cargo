package container

import "github.com/drgomesp/cargo/definition"

// Interface providing the basic API for a service container
type Interface interface {
	Register(id string, arg interface{}) (def *definition.Definition, err error)
	Set(id string, arg interface{}) (service *interface{}, err error)
	Get(id string) (service *interface{}, err error)
	MustGet(id string) interface{}
}
