package entities

import "errors"

const (
	North = iota
	East
	South
	West
)

type Robot struct {
	Facing int
}

func (r Robot) validate() error {
	if r.Facing < North || r.Facing > West {
		return errors.New("Robot must be facing to a valid compass point")
	}

	return nil
}

func (r Robot) move(x, y int) (int, int) {
	switch r.Facing {
	case North:
		return x, y + 1
	case East:
		return x + 1, y
	case South:
		return x, y - 1
	case West:
		return x - 1, y
	default:
		return x, y
	}
}

/*

	I don't think this approach is correct.

	I think the Robot should try to move, the board should not let it if it has hit a limit.

	This is kind of happening already, actually.

*/
