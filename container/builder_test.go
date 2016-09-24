package container

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type Foo struct{}
type Bar struct{}

func NewBar() *Bar {
	return &Bar{}
}

func TestRegisterWithInstance(t *testing.T) {
	Convey("Given a container builder is available", t, func() {
		builder := NewBuilder()

		Convey("Given an instance of an arbitrary type", func() {
			expected := new(Foo)

			Convey("When the instance is registered into the container builder ", func() {
				b, err := builder.Register("foo", expected)

				Convey("Then the register method should return the builder", func() {
					So(err, ShouldBeEmpty)
				})

				Convey("And the register method should return an empty error", func() {
					So(b, ShouldEqual, b)
				})
			})
		})
	})
}

func TestRegisterWithConstructorFunction(t *testing.T) {
	Convey("Given a container builder is available", t, func() {
		builder := NewBuilder()

		Convey("Given a constructor function that returns an arbitrary type", func() {
			Convey("When the function is registered into the container builder ", func() {
				b, err := builder.Register("foo", NewBar)

				Convey("Then the register method should return the builder", func() {
					So(err, ShouldBeEmpty)
				})

				Convey("And the register method should return an empty error", func() {
					So(b, ShouldEqual, b)
				})
			})
		})
	})
}

func TestRegisterHasDefinition(t *testing.T) {
	Convey("Given a container builder is available", t, func() {
		builder := NewBuilder()

		Convey("Given an arbitrary type", func() {
			type Foo struct{}

			Convey("When the type is registered into the container builder", func() {
				builder.Register("foo", &Foo{})

				Convey("Then the container should have a definition for that type", func() {
					So(builder.HasDefinition("foo"), ShouldBeTrue)
				})
			})
		})
	})
}

func TestRegisterGetDefinitionRegisteredWithInstance(t *testing.T) {
	Convey("Given a container builder is available", t, func() {
		builder := NewBuilder()

		Convey("Given an arbitrary type", func() {
			type Foo struct{}

			Convey("When the type is registered into the container builder", func() {
				builder.Register("foo", new(Foo))

				Convey("Then the container should have a definition for that type", func() {
					So(builder.HasDefinition("foo"), ShouldBeTrue)
				})

				Convey("And when requesting the container for that definition", func() {
					foo, err := builder.Get("foo")

					Convey("It should return a service for it", func() {
						So(err, ShouldBeNil)
						So(foo, ShouldHaveSameTypeAs, &Foo{})
					})
				})
			})
		})
	})
}

func TestRegisterGetDefinitionRegisteredWithConstructorFunction(t *testing.T) {
	Convey("Given a container builder is available", t, func() {
		builder := NewBuilder()

		Convey("Given an arbitrary type that has a constructor function", func() {
			Convey("When the function is registered into the container builder", func() {
				builder.Register("bar", &Bar{})

				Convey("Then the container should have a definition for that type", func() {
					So(builder.HasDefinition("bar"), ShouldBeTrue)
				})

				Convey("And when requesting the container for that definition", func() {
					foo, err := builder.Get("bar")

					Convey("It should return a service for it", func() {
						So(err, ShouldBeNil)
						So(foo, ShouldHaveSameTypeAs, &Bar{})
					})
				})
			})
		})
	})
}

func TestRegisterReturnsErrorWhenDefinitionAlreadyExist(t *testing.T) {
	Convey("Given a container builder is available", t, func() {
		builder := NewBuilder()

		Convey("Given an arbitrary type", func() {
			type Foo struct{}

			Convey("When the type is registered into the container builder", func() {
				builder.Register("foo", &Foo{})

				Convey("Then the container should have a definition for that type", func() {
					So(builder.HasDefinition("foo"), ShouldBeTrue)
				})

				Convey("And when trying to register the definition again", func() {
					_, err := builder.Register("foo", &Foo{})

					Convey("Then the container should return an error", func() {
						So(err, ShouldNotBeNil)
					})
				})
			})
		})
	})
}

func TestRegisterReturnsErrorWhenServiceIsNotFound(t *testing.T) {
	Convey("Given a container builder is available", t, func() {
		builder := NewBuilder()

		Convey("When a service that does not exist is requested", func() {
			_, err := builder.Get("service_that_does_not_exist")

			Convey("Then the container should return an error", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}
