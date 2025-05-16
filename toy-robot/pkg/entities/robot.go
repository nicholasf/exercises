package entities

import "errors"

// basic movement behaviour
type mover func(x, y int) (int, int)

const (
	North = iota
	East
	South
	West
)

var (
	MoveNorth = func(x, y int) (int, int) { return x, y + 1 }
	MoveEast  = func(x, y int) (int, int) { return x + 1, y }
	MoveSouth = func(x, y int) (int, int) { return x, y - 1 }
	MoveWest  = func(x, y int) (int, int) { return x - 1, y }

	compass = [4]mover{MoveNorth, MoveEast, MoveSouth, MoveWest}
)

type Robot struct {
	Facing int // an index to the orientation slice
}

// func NewRobot(facing int) (r Robot, err error) {

// }

func (r Robot) validate() error {
	if r.Facing < North || r.Facing > West {
		return errors.New("Robot must be facing to a valid compass point")
	}

	return nil
}

func (r *Robot) left() {
	r.Facing = r.Facing - 1

	if r.Facing < 0 {
		r.Facing = 4
	}
}

func (r Robot) move(x, y int) (int, int) {
	return compass[r.Facing](x, y)
}

func (r *Robot) right() {
	r.Facing = r.Facing + 1

	if r.Facing > 4 {
		r.Facing = 0
	}
}
