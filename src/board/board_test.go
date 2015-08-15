package board

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBoard(t *testing.T) {
	Convey("Given a new board of 3x3", t, func() {
		board := NewBoard(3)

		Convey("When setting a piece on the board", func() {
			board.Set("X", 0, 0)
			Convey("Then the space should have the piece", func() {
				So(board.Get(0, 0), ShouldEqual, "X")
			})
			Convey("Then another space should not have the piece", func() {
				So(board.Get(1, 1), ShouldNotEqual, "X")
			})
		})

		Convey("When asking for the number of entries", func() {
			Convey("then it should be 0", func() {
				So(board.Entries(), ShouldEqual, 0)
			})
			board.Set("X", 0, 0)
			Convey("then it should be 1", func() {
				So(board.Entries(), ShouldEqual, 1)
			})
			board.Set("X", 0, 1)
			Convey("then it should be 2", func() {
				So(board.Entries(), ShouldEqual, 2)
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

				Convey("new contents should be of the form", func() {
					expected := [][]string{
						[]string{"", "", ""},
						[]string{"", "", "O"},
						[]string{"", "", ""}}

					So(newContents, ShouldResemble, expected)
				})

				Convey("old contents should be of the form", func() {
					expected := [][]string{
						[]string{"", "", ""},
						[]string{"", "", ""},
						[]string{"", "", ""}}

					So(oldContents, ShouldResemble, expected)
				})
			})
		})

		Convey("When the board is full", func() {
			board.Set("X", 0, 0) //X
			board.Set("O", 1, 1) //O
			board.Set("X", 0, 1) //X
			board.Set("O", 0, 2) //O
			board.Set("X", 2, 0) //X
			board.Set("O", 1, 2) //O
			board.Set("X", 1, 0) //X
			board.Set("X", 2, 2) //O
			board.Set("X", 2, 1) //X

			So(board.IsFull(), ShouldBeTrue)
		})
	})
}
