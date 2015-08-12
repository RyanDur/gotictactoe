package board

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBoard(t *testing.T) {
	Convey("Given a new board of 3x3", t, func() {
		board := NewBoard(3)

		Convey("when I enquire the size", func() {
			size := board.size
			Convey("Then it should equal 3", func() {
				So(size, ShouldEqual, 3)
			})
		})

		Convey("When setting a piece on the board", func() {
			board.Set("X", 0, 0)
			Convey("Then the space should have the piece", func() {
				So(board.Get(0, 0), ShouldEqual, "X")
			})
			Convey("Then another space should not have the piece", func() {
				So(board.Get(1, 1), ShouldNotEqual, "X")
			})
		})

		Convey("When getting the contents of the board", func() {
			oldContents := board.Contents()
			Convey("When setting a new piece", func() {
				board.Set("O", 1, 2)
				newContents := board.Contents()
				Convey("should not affect the old contents", func() {
					So(oldContents, ShouldNotResemble, newContents)
				})
			})
		})
	})
}
