package container

import (
	"reflect"
	"testing"

	"github.com/drgomesp/cargo/argument"
	"github.com/drgomesp/cargo/definition"
	"github.com/drgomesp/cargo/method"
	"github.com/drgomesp/cargo/reference"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNew(t *testing.T) {
	Convey("Given a new container is created", t, func() {
		container := New()

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
			container := New()

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
			container := New()

			Convey("When the function is used to register a service on the container", func() {
				def, err := container.Register("foo", NewFoo)

				Convey("Then it should return an empty error", func() {
					So(err, ShouldBeNil)
				})

				Convey("And it should return a definition of that service", func() {
					So(def, ShouldHaveSameTypeAs, &definition.Definition{})
					So(def.Arguments(), ShouldHaveLength, 0)
					So(def.Type(), ShouldHaveSameTypeAs, reflect.TypeOf(&Foo{}))
					So(def.Constructor(), ShouldHaveSameTypeAs, reflect.Value{})
				})
			})

		})
	})
}

func TestRegisterWithInstance(t *testing.T) {
	Convey("Given an instance of an arbitraty type", t, func() {
		type Foo struct{}
		foo := new(Foo)

		Convey("And a service container instance", func() {
			container := New()

			Convey("When the instance is used to register a service on the container", func() {
				def, err := container.Register("foo", foo)

				Convey("Then it should return an empty error", func() {
					So(err, ShouldBeNil)
				})

				Convey("And it should return a definition of that service", func() {
					So(def, ShouldHaveSameTypeAs, &definition.Definition{})
					So(def.Arguments(), ShouldHaveLength, 0)
					So(def.Type(), ShouldHaveSameTypeAs, reflect.TypeOf(&Foo{}))
					So(def.Constructor(), ShouldHaveSameTypeAs, reflect.Value{})
				})
			})
		})
	})
}

func TestSetServiceWithInvalidType(t *testing.T) {
	Convey("Given a service container instance", t, func() {
		container := New()

		Convey("When setting a service of invalid type", func() {
			err := container.Set("foo", 1)

			Convey("Then it should return an error", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "Could not create definition")
			})
		})
	})
}

func TestSetServiceWithAlreadyExistingIdentifier(t *testing.T) {
	Convey("Given a service container instance", t, func() {
		container := New()

		Convey("And an existing service \"foo\"", func() {
			type Foo struct{}
			foo := &Foo{}
			container.Set("foo", foo)

			Convey("When setting a service \"foo\"", func() {
				err := container.Set("foo", &Foo{})

				Convey("Then it should return an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "Definition for \"foo\" already exists")
				})
			})
		})
	})
}

func TestGetServiceRegisteredWithConstructorFunction(t *testing.T) {
	Convey("Given a service container instance", t, func() {
		container := New()

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
					So(def.Arguments(), ShouldHaveLength, 0)
					So(def.Type(), ShouldHaveSameTypeAs, reflect.TypeOf(&Foo{}))
					So(def.Constructor(), ShouldHaveSameTypeAs, reflect.Value{})

					Convey("And when requesting for that service named \"foo\" from the container", func() {
						foo, err := container.Get("foo")

						Convey("Then it should return an empty error", func() {
							So(err, ShouldBeNil)
						})

						Convey("And it should return an instance of that service", func() {
							So(foo, ShouldHaveSameTypeAs, NewFoo())
						})
					})
				})
			})
		})
	})
}

func TestGetServiceSetWithInstance(t *testing.T) {
	Convey("Given a service container instance", t, func() {
		container := New()

		Convey("And an instance of an arbitrary type", func() {
			type Foo struct {
				A int
				B string
			}
			foo := &Foo{10, "FOO"}

			Convey("When that instance is registered as a service \"foo\" in the container", func() {
				err := container.Set("foo", foo)

				Convey("Then the container should return an empty error", func() {
					So(err, ShouldBeNil)
				})

				Convey("And when requesting for that service named \"foo\" from the container", func() {
					ret, err := container.Get("foo")

					Convey("Then it should return an empty error", func() {
						So(err, ShouldBeNil)
					})

					Convey("And it should return a pointer to the same service", func() {
						So(ret, ShouldPointTo, foo)

						original := ret.(*Foo)
						So(original.A, ShouldEqual, 10)
						So(original.B, ShouldEqual, "FOO")
					})
				})
			})
		})
	})
}

func TestGetServiceWithDifferentCaseCharactersForIdentifier(t *testing.T) {
	Convey("Given a service container instance", t, func() {
		container := New()

		Convey("And an existing service \"foo\"", func() {
			type Foo struct{}
			foo := &Foo{}
			container.Set("foo", foo)

			Convey("When requesting for a service \"FOO\"", func() {
				ret, err := container.Get("FoO")

				Convey("Then it should return an empty error", func() {
					So(err, ShouldBeNil)
				})

				Convey("And it should a pointer to the same service", func() {
					So(ret, ShouldPointTo, foo)
				})
			})
		})
	})
}

func TestGetServiceRegisteredWithConstructorFunctionAndArguments(t *testing.T) {
	Convey("Given a service container instance", t, func() {
		container := New()

		Convey("And an arbitrary type with a constructor function", func() {
			type Foo struct {
				Number int
				Text   string
			}

			NewFoo := func(number int, text string) *Foo {
				return &Foo{
					Number: number,
					Text:   text,
				}
			}

			Convey("And service of that type registered as \"foo\" in the container", func() {
				def, err := container.Register("foo", NewFoo)
				def.AddArguments(argument.New(100), argument.New("constructor_was_called"))

				Convey("Then it should return an empty error", func() {
					So(err, ShouldBeNil)
				})

				Convey("And it should return a definition of that service", func() {
					So(def, ShouldHaveSameTypeAs, &definition.Definition{})
					So(def.Arguments(), ShouldHaveLength, 2)
					So(def.Type(), ShouldHaveSameTypeAs, reflect.TypeOf(&Foo{}))
					So(def.Constructor(), ShouldHaveSameTypeAs, reflect.Value{})
				})

				Convey("And when requesting for that service named \"foo\" from the container", func() {
					foo, err := container.Get("foo")

					Convey("Then it should return an empty error", func() {
						So(err, ShouldBeNil)
					})

					Convey("And it should return an instance of that service", func() {
						So(foo, ShouldHaveSameTypeAs, &Foo{})
						ret := foo.(*Foo)

						So(ret.Number, ShouldEqual, 100)
						So(ret.Text, ShouldEqual, "constructor_was_called")
					})
				})
			})
		})
	})
}

type Foo struct {
	Number int
	Text   string
}

func (f *Foo) Bar(number int, text string) {
	f.Number = number
	f.Text = text
}

func TestGetServiceRegisteredWithConstructorFunctionAndMethodCalls(t *testing.T) {
	Convey("Given a service container instance", t, func() {
		container := New()

		Convey("And an arbitrary type with a constructor function", func() {

			NewFoo := func(number int, text string) *Foo {
				return &Foo{number, text}
			}

			Convey("And service of that type registered as \"foo\" in the container", func() {
				def, err := container.Register("foo", NewFoo)
				def.AddArguments(argument.New(999), argument.New("constructor_was_called"))
				def.AddMethodCall(method.New("Bar", 5, "bar_was_called"))

				Convey("Then it should return an empty error", func() {
					So(err, ShouldBeNil)
				})

				Convey("And it should return a definition of that service", func() {
					So(def, ShouldHaveSameTypeAs, &definition.Definition{})
					So(def.Arguments(), ShouldHaveLength, 2)
					So(def.MethodCalls(), ShouldHaveLength, 1)
					So(def.Constructor(), ShouldHaveSameTypeAs, reflect.Value{})
					So(def.Type(), ShouldHaveSameTypeAs, reflect.TypeOf(&Foo{}))
				})

				Convey("And when requesting for that service named \"foo\" from the container", func() {
					foo, err := container.Get("foo")

					Convey("Then it should return an empty error", func() {
						So(err, ShouldBeNil)
					})

					Convey("And it should return an instance of that service", func() {
						So(foo, ShouldHaveSameTypeAs, &Foo{})

						ret := foo.(*Foo)

						So(ret.Number, ShouldEqual, 5)
						So(ret.Text, ShouldEqual, "bar_was_called")
					})
				})
			})
		})
	})
}

func TestRegisteringWithConstructorFunctionAndMethodCallsWithWrongNumberOfParameters(t *testing.T) {
	Convey("Given a service container instance", t, func() {
		container := New()

		Convey("And an arbitrary type with a constructor function", func() {

			NewFoo := func(number int, text string) *Foo {
				return &Foo{number, text}
			}

			Convey(`And service of that type registered as "foo" in the container`, func() {
				def, err := container.Register("foo", NewFoo)
				def.AddArguments(argument.New(999), argument.New("constructor_was_called"))
				def.AddMethodCall(method.New("Bar"))

				Convey("Then it should return an empty error", func() {
					So(err, ShouldBeNil)
				})

				Convey("And it should return a definition of that service", func() {
					So(def, ShouldHaveSameTypeAs, &definition.Definition{})
					So(def.Arguments(), ShouldHaveLength, 2)
					So(def.MethodCalls(), ShouldHaveLength, 1)
					So(def.Type(), ShouldHaveSameTypeAs, reflect.TypeOf(&Foo{}))
					So(def.Constructor(), ShouldHaveSameTypeAs, reflect.Value{})
				})

				Convey(`And when requesting for that service named "foo" from the container`, func() {
					_, err := container.Get("foo")

					Convey("Then it should return an error", func() {
						So(err.Error(), ShouldEqual, `Method "Bar" expects 2 arguments`)
					})
				})
			})
		})
	})
}

func TestRegisteringWithReferenceToExistingService(t *testing.T) {
	Convey(`Given a service container instance with some existing "foo" service`, t, func() {
		container := New()
		foo := &Foo{123, "existing_service"}
		container.Set("foo", foo)

		Convey("When registering a service definition that references an existing service as an argument", func() {
			type Bar struct {
				FooService *Foo
			}

			newBar := func(fooService *Foo) *Bar {
				return &Bar{fooService}
			}

			def, _ := container.Register("bar", newBar)
			ref := reference.New("foo")
			def.AddArguments(&ref)

			Convey("When requesting for that service", func() {
				bar, err := container.Get("bar")

				Convey("Then there should be no error", func() {
					So(err, ShouldBeNil)
				})

				Convey(`And a new "bar" service should be created with "foo" injected through the constructor`, func() {
					So(bar.(*Bar).FooService, ShouldEqual, foo)
				})
			})
		})
	})
}

func TestGetNonExistingService(t *testing.T) {
	Convey("Given a service container instance", t, func() {
		container := New()

		Convey("When requesting for a non-existing service \"bar\"", func() {
			_, err := container.Get("bar")

			Convey("Then it should return an error", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, `No service "bar" was found`)
			})
		})
	})
}

func TestMustGetServiceSetWithInstance(t *testing.T) {
	Convey("Given a service container instance", t, func() {
		container := New()

		Convey("And an instance of an arbitrary type", func() {
			type Foo struct {
				A int
				B string
			}
			foo := &Foo{10, "FOO"}

			Convey("When that instance is registered as a service \"foo\" in the container", func() {
				err := container.Set("foo", foo)

				Convey("Then the container should return an empty error", func() {
					So(err, ShouldBeNil)
				})

				Convey("And when requesting as a must for that service named \"foo\" from the container", func() {
					So(func() {
						container.MustGet("foo")
					}, ShouldNotPanic)

					ret := container.MustGet("foo")
					Convey("And it should return a pointer to the same service", func() {
						So(ret, ShouldPointTo, foo)

						original := ret.(*Foo)
						So(original.A, ShouldEqual, 10)
						So(original.B, ShouldEqual, "FOO")
					})
				})
			})
		})
	})
}

func TestMustGetNonExistingService(t *testing.T) {
	Convey("Given a service container instance", t, func() {
		container := New()

		Convey("When requesting for a non-existing service \"bar\"", func() {
			So(func() {
				container.MustGet("bar")
			}, ShouldPanic)
		})
	})
}
