package entities

import "testing"

var (
	b1            = Board{1, 0, 0, Robot{North}}
	b2            = Board{1, 4, 4, Robot{North}}
	bb1           = Board{1, -1, -1, Robot{South}}
	bb2           = Board{1, 5, 5, Robot{South}}
	bb3           = Board{1, 5, 5, Robot{6}}
	bb4           = Board{1, 5, 5, Robot{-1}}
	validBoards   = [2]Board{b1, b2}
	invalidBoards = [4]Board{bb1, bb2, bb3, bb4}
)

func TestNewBoard(t *testing.T) {
	t.Run("Construct valid boards", func(t *testing.T) {
		for _, board := range validBoards {
			b, err := NewBoard(board.RobotX, board.RobotY, North)

			if err != nil {
				t.Log("Valid board should be valid!")
				t.Fail()
			}

			if b.ID == 0 {
				t.Log("Valid board should have a positive integer id!")
				t.Fail()
			}
		}
	})

	t.Run("Construct invalid boards", func(t *testing.T) {
		for _, board := range invalidBoards {
			b, err := NewBoard(board.RobotX, board.RobotY, North)

			if err == nil {
				t.Log("Invalid board should not be created! ", err)
				t.Fail()
			}

			if b != nil {
				t.Log("Invalid board should not exist!", err)
				t.Fail()
			}
		}
	})
}

func TestValidate(t *testing.T) {

	t.Run("Validate valid boards", func(t *testing.T) {
		for _, board := range validBoards {
			err := board.validate()

			if err != nil {
				t.Log("Valid board should be valid!")
				t.Fail()
			}
		}
	})

	t.Run("Validate invalid boards", func(t *testing.T) {
		for _, board := range invalidBoards {
			err := board.validate()

			if err == nil {
				t.Log("Invalid board should be invalid!")
				t.Fail()
			}
		}
	})

}
