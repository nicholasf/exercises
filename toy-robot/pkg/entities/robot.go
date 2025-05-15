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
