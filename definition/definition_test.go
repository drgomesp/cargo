package definition

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewWithConstructorFunction(t *testing.T) {
	Convey("Given a constructor function for an arbitraty type", t, func() {
		type Foo struct{}

		Convey("When that function is used to create a definition", func() {
			def, err := NewDefinition(func() *Foo {
				return &Foo{}
			})

			Convey("Then it should return an empty error", func() {
				So(err, ShouldBeNil)
			})

			Convey("And the definition should represent the correct type", func() {
				So(def, ShouldHaveSameTypeAs, &Definition{})
				So(def.Arguments, ShouldHaveLength, 0)
				So(def.Type, ShouldHaveSameTypeAs, reflect.TypeOf(&Foo{}))
				So(def.Constructor, ShouldHaveSameTypeAs, reflect.Value{})
			})
		})
	})
}

func TestNewWithInstance(t *testing.T) {
	Convey("Given an instance of an arbitrary type", t, func() {
		type Foo struct{}
		foo := &Foo{}

		Convey("When that instance is used to create a definition", func() {
			def, err := NewDefinition(foo)

			Convey("Then it should return an empty error", func() {
				So(err, ShouldBeNil)
			})

			Convey("And it should return a definition that represents the correct type", func() {
				So(def, ShouldHaveSameTypeAs, &Definition{})
				So(def.Arguments, ShouldHaveLength, 0)
				So(def.Type, ShouldHaveSameTypeAs, reflect.TypeOf(&Foo{}))
				So(def.Constructor, ShouldHaveSameTypeAs, reflect.Value{})
			})
		})
	})
}

func TestAddArgument(t *testing.T) {
	Convey("Given an arbitrary type", t, func() {
		type Foo struct{}

		Convey("And a definition of that type", func() {
			def, err := NewDefinition(&Foo{})

			Convey("Then it should return an empty error", func() {
				So(err, ShouldBeNil)
			})

			Convey("And when an argument is added to it", func() {
				arg := NewArgument(100)
				def.AddArguments(arg)

				Convey("Then it should return an empty error", func() {
					So(err, ShouldBeNil)
				})

				Convey("And the definition should contain that argument", func() {
					So(def.Arguments, ShouldHaveLength, 1)
					So(def.Arguments[0].Value, ShouldEqual, reflect.ValueOf(arg.Value).Interface())
				})
			})
		})
	})
}

type Foo struct{}

func (f *Foo) Bar(a int, b string) {}

func TestAddMethodCall(t *testing.T) {
	Convey("Given an arbitrary type with a member method", t, func() {
		Convey("And a definition of that type", func() {
			def, err := NewDefinition(&Foo{})

			Convey("Then it should return an empty error", func() {
				So(err, ShouldBeNil)
			})

			Convey("And when a method call is added to it", func() {
				method := NewMethod("Bar", 0, "bar")
				def.AddMethodCall(method)

				Convey("Then it should return an empty error", func() {
					So(err, ShouldBeNil)
				})

				Convey("And the definition should contain that method call", func() {
					So(def.MethodCalls, ShouldContain, method)
				})
			})
		})
	})
}
