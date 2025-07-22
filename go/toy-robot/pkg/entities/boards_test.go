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

func TestRobot(t *testing.T) {

	// these fixtures represent obvious legal moves for the robot and the board
	t.Run("Reusing valid Robot ments", func(t *testing.T) {
		for _, ment := range validRobotMovements {
			t.Run(ment.label, func(t *testing.T) {
				board, err := NewBoard(ment.x, ment.y, ment.facing)

				if err != nil {
					t.Log("Could not create board")
					t.Fail()
				}

				err = board.MoveRobot()

				if err != nil {
					t.Log("Board disallowed a valid robot ")
					t.Fail()
				}
			})
		}
	})

	// The board should restrict a robot's ments if the robot is at the edge of a board
	invalidBoardments := []struct {
		x      int
		y      int
		xx     int
		yy     int
		facing int
		label  string
	}{
		{
			x:      4,
			y:      4,
			yy:     4,
			facing: North,
			label:  "Moving North",
		},
		{
			x:      4,
			y:      4,
			xx:     4,
			facing: East,
			label:  "Moving East",
		},
		{
			yy:     0,
			facing: South,
			label:  "Moving South",
		},
		{
			xx:     0,
			facing: West,
			label:  "Moving West",
		},
	}

	t.Run("Placing the robot at the corners of the table", func(t *testing.T) {
		for _, ment := range invalidBoardments {
			t.Run(ment.label, func(t *testing.T) {
				board, err := NewBoard(ment.x, ment.y, ment.facing)

				if err != nil {
					t.Log("Could not create board")
					t.Fail()
				}

				err = board.MoveRobot()

				if err == nil {
					t.Log("Board allowed an invalid robot ")
					t.Fail()
				}
			})
		}
	})
}

func TestReport(t *testing.T) {
	b, err := NewBoard(1, 1, North)

	if err != nil {
		t.Log("Could not create board")
		t.Fail()
	}

	x, y, facing := b.Report()

	if x != 1 || y != 1 || facing != North {
		t.Log("Report is failing")
		t.Fail()
	}
}
