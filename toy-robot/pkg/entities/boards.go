package entities

import "errors"

type BoardID int

type Board struct {
	ID     int
	RobotX int
	RobotY int
	Robot  Robot
}

const (
	BoardLimitX = 5
	BoardLimitY = 5
)

var boards []*Board

var BoardIDCounter = 0

func NewBoard(robotX, robotY, facing int) (*Board, error) {
	if boards == nil {
		boards = make([]*Board, 1)
	}

	BoardIDCounter = BoardIDCounter + 1

	board := &Board{
		ID:     BoardIDCounter,
		RobotX: robotX,
		RobotY: robotY,
	}

	err := board.validate()

	if err != nil {
		return nil, err
	}

	// Add the Robot, which records facing compass directions
	board.Robot = Robot{facing}
	err = board.Robot.validate()

	if err != nil {
		return nil, err
	}

	boards = append(boards, board)
	return board, nil
}

func (b Board) validate() error {
	if b.RobotX < 0 || b.RobotX >= BoardLimitX {
		return errors.New("robot x out of bounds")
	}

	if b.RobotY < 0 || b.RobotY >= BoardLimitY {
		return errors.New("robot x out of bounds")
	}

	return nil
}

// The board tells the robot to move from its current location, which returns its intented x,y.
// The board won't allow the movement the coordinates it returns are invalid.
func (b *Board) MoveRobot() error {
	x, y := b.Robot.move(b.RobotX, b.RobotY)

	replacementBoard := Board{0, x, y, b.Robot}
	err := replacementBoard.validate()

	if err != nil {
		return errors.New("Robot would fall off the board")
	}

	b.RobotX = x
	b.RobotY = y

	return nil
}
