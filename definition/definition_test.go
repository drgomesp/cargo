package definition

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	. "github.com/smartystreets/goconvey/convey"
)

type Bar struct{}

func NewBar() *Bar {
	return &Bar{}
}

func TestNewDefinitionWithInstance(t *testing.T) {
	Convey("Given an arbitrary type", t, func() {
		type Foo struct{}

		Convey("When creating a definition for that type using a composite literal", func() {
			def := NewDefinition(&Foo{})

			Convey("Then definition should be returned", func() {
				So(def, ShouldNotBeNil)
			})
		})

		Convey("And when creating a definition for that type using the new keyword", func() {
			def := NewDefinition(new(Foo))

			Convey("Then definition should be returned", func() {
				So(def, ShouldNotBeNil)
			})
		})
	})
}

func TestNewDefinitionWithConstructorFunction(t *testing.T) {
	Convey("Given an arbitrary type with a constructor function", t, func() {
		Convey("When creating a definition for that type ", func() {
			def := NewDefinition(NewBar)

			Convey("Then definition should be returned", func() {
				So(def, ShouldNotBeNil)
			})
		})
	})
}

func TestAddArguments(t *testing.T) {
	Convey("Given an arbitrary definition of an arbitrary type", t, func() {
		type Bar struct{}

		type Foo struct {
			A   int
			B   string
			Bar Bar
		}

		def := NewDefinition(&Foo{})

		Convey("When adding arguments to that definition ", func() {

			bar := NewDefinition(&Bar{})
			def.AddArguments(1, "string", bar)

			Convey("Then definition should be created with those arguments", func() {
				So(def, ShouldNotBeNil)
				spew.Dump(def)
				So(def.Arguments, ShouldHaveLength, 3)
				So(def.Arguments, ShouldContain, bar)
				So(def.Arguments, ShouldContain, 1)
				So(def.Arguments, ShouldContain, "string")
			})
		})
	})
}
