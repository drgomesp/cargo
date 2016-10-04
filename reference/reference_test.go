package reference

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewReference(t *testing.T) {
	Convey("Given a reference is created with a string identifier", t, func() {
		ref := New("foo")
		Convey("Then it should return reference with that identifier", func() {
			So(ref, ShouldHaveSameTypeAs, Reference{})
			So(ref.Identifier(), ShouldEqual, "foo")
		})
	})
}
