package container

import (
	"reflect"
	"testing"

	"github.com/drgomesp/cargo/definition"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewContainer(t *testing.T) {
	Convey("Given a new container is created", t, func() {
		container := NewContainer()

		Convey("Then a new container instance should be returned", func() {
			So(container, ShouldHaveSameTypeAs, &Container{})
		})
	})
}

func TestRegisterAlreadyExistingDefinition(t *testing.T) {
	Convey("Given a constructor function for an arbitraty type", t, func() {
		type Foo struct{}
		NewFoo := func() *Foo {
			return &Foo{}
		}

		Convey("And a service container instance", func() {
			container := NewContainer()

			Convey("When the function is used to register a service on the container", func() {
				container.Register("foo", NewFoo)

				Convey("And the function is used to register the service again", func() {
					_, err := container.Register("foo", NewFoo)

					Convey("Then it should return an error", func() {
						So(err, ShouldNotBeNil)
						So(err.Error(), ShouldEqual, "Definition for \"foo\" already exists")
					})
				})
			})

		})
	})
}

func TestRegisterWithConstructorFunction(t *testing.T) {
	Convey("Given a constructor function for an arbitraty type", t, func() {
		type Foo struct{}
		NewFoo := func() *Foo {
			return &Foo{}
		}

		Convey("And a service container instance", func() {
			container := NewContainer()

			Convey("When the function is used to register a service on the container", func() {
				def, err := container.Register("foo", NewFoo)

				Convey("Then it should return an empty error", func() {
					So(err, ShouldBeNil)
				})

				Convey("And it should return a definition of that service", func() {
					So(def, ShouldHaveSameTypeAs, &definition.Definition{})
					So(def.Arguments, ShouldHaveLength, 0)
					So(def.Type, ShouldHaveSameTypeAs, reflect.TypeOf(&Foo{}))
					So(def.Constructor, ShouldHaveSameTypeAs, reflect.Value{})
				})
			})

		})
	})
}

func TestRegisterWithInstance(t *testing.T) {
	Convey("Given an instance of an arbitraty type", t, func() {
		type Foo struct{}
		foo := &Foo{}

		Convey("And a service container instance", func() {
			container := NewContainer()

			Convey("When the instance is used to register a service on the container", func() {
				def, err := container.Register("foo", foo)

				Convey("Then it should return an empty error", func() {
					So(err, ShouldBeNil)
				})

				Convey("And it should return a definition of that service", func() {
					So(def, ShouldHaveSameTypeAs, &definition.Definition{})
					So(def.Arguments, ShouldHaveLength, 0)
					So(def.Type, ShouldHaveSameTypeAs, reflect.TypeOf(&Foo{}))
					So(def.Constructor, ShouldHaveSameTypeAs, reflect.Value{})
				})
			})

		})
	})
}

func TestGetServiceRegisteredWithConstructorFunction(t *testing.T) {
	Convey("Given a service container instance", t, func() {
		container := NewContainer()

		Convey("And an arbitrary type with a constructor function", func() {
			type Foo struct{}
			NewFoo := func() *Foo {
				return &Foo{}
			}

			Convey("And service of that type registered as \"foo\" in the container", func() {
				def, err := container.Register("foo", NewFoo)

				Convey("Then it should return an empty error", func() {
					So(err, ShouldBeNil)
				})

				Convey("And it should return a definition of that service", func() {
					So(def, ShouldHaveSameTypeAs, &definition.Definition{})
					So(def.Arguments, ShouldHaveLength, 0)
					So(def.Type, ShouldHaveSameTypeAs, reflect.TypeOf(&Foo{}))
					So(def.Constructor, ShouldHaveSameTypeAs, reflect.Value{})

					Convey("And when requesting for that service named \"foo\" from the container", func() {
						foo, err := container.Get("foo")

						Convey("Then it should return an empty error", func() {
							So(err, ShouldBeNil)
						})

						Convey("And it should return an instance of that service", func() {
							So(*foo, ShouldHaveSameTypeAs, NewFoo())
						})
					})
				})
			})
		})
	})
}
