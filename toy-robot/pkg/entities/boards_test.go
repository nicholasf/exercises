package entities

import "testing"

var (
	b1            = Board{1, 0, 0}
	b2            = Board{1, 4, 4}
	bb1           = Board{1, -1, -1}
	bb2           = Board{1, 5, 5}
	validBoards   = [2]Board{b1, b2}
	invalidBoards = [2]Board{bb1, bb2}
)

func TestNewBoard(t *testing.T) {
	t.Run("Construct valid boards", func(t *testing.T) {
		for _, board := range validBoards {
			b, err := NewBoard(board.RobotX, board.RobotY)

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
			b, err := NewBoard(board.RobotX, board.RobotY)

			if err == nil {
				t.Log("Invalid board should not be created!")
				t.Fail()
			}

			if b != nil {
				t.Log("Invalid board should not exist!")
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
