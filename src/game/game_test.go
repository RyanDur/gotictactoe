package game

import (
	. "board"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGame(t *testing.T) {

	Convey("Given a new game", t, func() {
		board := NewBoard(3)
		game := NewGame(board)

		Convey("When playing the first move", func() {
			game.Set(0, 0)
			Convey("then it should always be X", func() {
				expected := [][]string{
					[]string{"X", "", ""},
					[]string{"", "", ""},
					[]string{"", "", ""}}

				So(game.GetBoard(), ShouldResemble, expected)
			})
		})

		Convey("When playing the second move", func() {
			game.Set(0, 0)
			Convey("then it should always be X", func() {
				game.Set(0, 1)
				expected := [][]string{
					[]string{"X", "O", ""},
					[]string{"", "", ""},
					[]string{"", "", ""}}

				So(game.GetBoard(), ShouldResemble, expected)
			})
			Convey("then it cannot occupy a taken space", func() {
				err := game.Set(0, 0)
				So(err, ShouldNotBeNil)
			})
		})

		Convey("When a game is over", func() {
			Convey("And X has Won via the top row", func() {
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")
				game.Set(0, 0)
				game.Set(1, 0)
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")

				game.Set(0, 1)
				game.Set(2, 0)
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")

				game.Set(0, 2)
				So(game.IsOver(), ShouldBeTrue)
				So(game.Winner(), ShouldEqual, "X")
			})

			Convey("And X has Won via the middle row", func() {
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")
				game.Set(1, 0)
				game.Set(0, 0)
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")

				game.Set(1, 1)
				game.Set(2, 0)
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")

				game.Set(1, 2)
				So(game.IsOver(), ShouldBeTrue)
				So(game.Winner(), ShouldEqual, "X")
			})
			Convey("And X has Won via the bottom row", func() {
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")
				game.Set(2, 0)
				game.Set(0, 0)
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")

				game.Set(2, 1)
				game.Set(1, 0)
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")

				game.Set(2, 2)
				So(game.IsOver(), ShouldBeTrue)
				So(game.Winner(), ShouldEqual, "X")
			})

			Convey("And O has Won via the top row", func() {
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")
				game.Set(1, 0)
				game.Set(0, 0)
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")

				game.Set(2, 0)
				game.Set(0, 1)
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")

				game.Set(1, 1)
				game.Set(0, 2)

				So(game.IsOver(), ShouldBeTrue)
				So(game.Winner(), ShouldEqual, "O")
			})
			Convey("And O has Won via the middle row", func() {
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")
				game.Set(0, 0)
				game.Set(1, 0)
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")

				game.Set(2, 0)
				game.Set(1, 1)
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")

				game.Set(0, 2)
				game.Set(1, 2)

				So(game.IsOver(), ShouldBeTrue)
				So(game.Winner(), ShouldEqual, "O")
			})
			Convey("And O has Won via the bottom row", func() {
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")
				game.Set(0, 0)
				game.Set(2, 0)
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")

				game.Set(1, 0)
				game.Set(2, 1)
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")

				game.Set(0, 1)
				game.Set(2, 2)

				So(game.IsOver(), ShouldBeTrue)
				So(game.Winner(), ShouldEqual, "O")
			})

			Convey("And X has Won via the left column", func() {
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")
				game.Set(0, 0)
				game.Set(0, 1)
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")

				game.Set(1, 0)
				game.Set(2, 2)
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")

				game.Set(2, 0)

				So(game.IsOver(), ShouldBeTrue)
				So(game.Winner(), ShouldEqual, "X")
			})
			Convey("And X has Won via the middle column", func() {
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")
				game.Set(0, 1)
				game.Set(0, 0)
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")

				game.Set(1, 1)
				game.Set(2, 2)
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")

				game.Set(2, 1)

				So(game.IsOver(), ShouldBeTrue)
				So(game.Winner(), ShouldEqual, "X")
			})
			Convey("And X has Won via the right column", func() {
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")
				game.Set(0, 2)
				game.Set(0, 0)
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")

				game.Set(1, 2)
				game.Set(2, 1)
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")

				game.Set(2, 2)

				So(game.IsOver(), ShouldBeTrue)
				So(game.Winner(), ShouldEqual, "X")
			})

			Convey("And O has Won via the left column", func() {
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")
				game.Set(0, 1)
				game.Set(0, 0)
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")

				game.Set(2, 2)
				game.Set(1, 0)
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")

				game.Set(2, 1)
				game.Set(2, 0)

				So(game.IsOver(), ShouldBeTrue)
				So(game.Winner(), ShouldEqual, "O")
			})
			Convey("And O has Won via the middle column", func() {
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")
				game.Set(0, 0)
				game.Set(0, 1)
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")

				game.Set(2, 2)
				game.Set(1, 1)
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")

				game.Set(2, 0)
				game.Set(2, 1)

				So(game.IsOver(), ShouldBeTrue)
				So(game.Winner(), ShouldEqual, "O")
			})
			Convey("And O has Won via the right column", func() {
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")
				game.Set(0, 1)
				game.Set(0, 2)
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")

				game.Set(2, 0)
				game.Set(1, 2)
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")

				game.Set(2, 1)
				game.Set(2, 2)

				So(game.IsOver(), ShouldBeTrue)
				So(game.Winner(), ShouldEqual, "O")
			})

			Convey("And X has won via left diagonal", func() {
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")
				game.Set(0, 0)
				game.Set(0, 2)
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")

				game.Set(1, 1)
				game.Set(1, 2)
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")

				game.Set(2, 2)

				So(game.IsOver(), ShouldBeTrue)
				So(game.Winner(), ShouldEqual, "X")
			})
			Convey("And O has won via left diagonal", func() {
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")
				game.Set(0, 2)
				game.Set(0, 0)
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")

				game.Set(1, 2)
				game.Set(1, 1)
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")

				game.Set(1, 0)
				game.Set(2, 2)

				So(game.IsOver(), ShouldBeTrue)
				So(game.Winner(), ShouldEqual, "O")
			})

			Convey("And X has won via right diagonal", func() {
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")
				game.Set(0, 2)
				game.Set(0, 1)
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")

				game.Set(1, 1)
				game.Set(1, 2)
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")

				game.Set(2, 0)

				So(game.IsOver(), ShouldBeTrue)
				So(game.Winner(), ShouldEqual, "X")
			})
			Convey("And O has won via right diagonal", func() {
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")
				game.Set(0, 0)
				game.Set(0, 2)
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")

				game.Set(1, 2)
				game.Set(1, 1)
				So(game.IsOver(), ShouldBeFalse)
				So(game.Winner(), ShouldEqual, "")

				game.Set(1, 0)
				game.Set(2, 0)

				So(game.IsOver(), ShouldBeTrue)
				So(game.Winner(), ShouldEqual, "O")
			})

			Convey("And the board is full with no winner", func() {
				game.Set(0, 2) //X
				game.Set(1, 1) //O
				game.Set(0, 1) //X
				game.Set(0, 0) //O
				game.Set(2, 0) //X
				game.Set(1, 2) //O
				game.Set(1, 0) //X
				game.Set(2, 1) //O
				game.Set(2, 2) //X

				So(game.IsOver(), ShouldBeTrue)
				So(game.Winner(), ShouldEqual, "")
			})

			Convey("And cannot set a move after game over", func() {
				Convey("When the board is full", func() {
					game.Set(0, 2) //X
					game.Set(1, 1) //O
					game.Set(0, 1) //X
					game.Set(0, 0) //O
					game.Set(2, 0) //X
					game.Set(1, 2) //O
					game.Set(1, 0) //X
					game.Set(2, 1) //O
					game.Set(2, 2) //X

					err := game.Set(2, 2)
					So(err, ShouldNotBeNil)
				})
				Convey("When there is a winner", func() {
					game.Set(0, 0)
					game.Set(0, 2)

					game.Set(1, 2)
					game.Set(1, 1)

					game.Set(1, 0)
					game.Set(2, 0)

					err := game.Set(2, 2)
					So(err, ShouldNotBeNil)

				})
			})
		})
	})
}
