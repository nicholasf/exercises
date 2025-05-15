package entities

import "errors"

type BoardID int

type Board struct {
	ID     int
	RobotX int
	RobotY int
}

const (
	BoardLimitX = 5
	BoardLimitY = 5
)

var Boards []*Board

var BoardIDCounter = 0

func NewBoard(id, robotX, robotY int) (*Board, error) {

	if Boards == nil {
		Boards = make([]*Board, 1)
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
