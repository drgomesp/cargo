package container

import (
	"reflect"
	"strings"

	"github.com/drgomesp/cargo"
	"github.com/drgomesp/cargo/definition"
)

const (
	// DefinitionConflicted represents that the definition was already previously created
	DefinitionConflicted = -1

	// DefinitionNotFound represents a not found definition for the given id
	DefinitionNotFound = -2
)

// Builder provides an API to easily build service definitions
type Builder struct {
	Container

	typeRegistry map[string]reflect.Type
}

// NewBuilder creates a new builder
func NewBuilder() *Builder {
	return &Builder{
		Container:    *NewContainer(),
		typeRegistry: make(map[string]reflect.Type),
	}
}

// Register a new service definition
func (b *Builder) Register(id string, fn interface{}) (builder *Builder, err error) {
	if _, ok := b.Container.definitions[id]; ok {
		return builder, cargo.NewError(DefinitionConflicted, "Definition \"%s\" already exists", id)
	}

	var t reflect.Type
	var v reflect.Value

	v = reflect.ValueOf(fn)
	t = reflect.TypeOf(fn)

	switch v.Kind() {
	case reflect.Ptr:
		b.Container.definitions[id] = definition.NewDefinition(id, v.Interface())
	case reflect.Func:
		t = reflect.TypeOf(v.Call(nil)[0].Interface())
		b.Container.definitions[id] = definition.NewDefinition(id, (v.Call(nil)[0]).Interface())
	}

	b.typeRegistry[id] = t
	return
}

// Get an instance by its identifier
func (b *Builder) Get(id string) (service interface{}, err error) {
	id = strings.ToLower(id)

	if service, ok := b.Container.services[id]; ok {
		return service.Interface(), nil
	}

	if definition, ok := b.definitions[id]; ok {
		return b.createService(definition, id)
	}

	return service, cargo.NewError(DefinitionConflicted, "Service \"%s\" was not found", id)
}

func (b *Builder) createService(def definition.Definition, id string) (service interface{}, err error) {
	if _, ok := b.Container.definitions[id]; !ok {
		return
	}

	value := reflect.ValueOf(service)

	if !value.IsValid() {
		s := reflect.New(b.Container.definitions[id].Type)
		service = reflect.Indirect(s).Interface()
		return
	}

	b.Container.services[id] = value
	service = value

	return
}
