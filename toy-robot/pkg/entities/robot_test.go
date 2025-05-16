package entities

import "testing"

var (
	validRobotMovements = []struct {
		x      int
		y      int
		xx     int
		yy     int
		facing int
		label  string
	}{
		{
			yy:     1,
			facing: North,
			label:  "Moving North",
		},
		{
			xx:     1,
			facing: East,
			label:  "Moving East",
		},
		{
			y:      1,
			yy:     0,
			facing: South,
			label:  "Moving South",
		},
		{
			x:      1,
			xx:     0,
			facing: West,
			label:  "Moving West",
		},
	}

	invalidRobotMovements = []struct {
		x      int
		y      int
		xx     int
		yy     int
		facing int
		label  string
	}{
		{
			x:      2,
			y:      2,
			yy:     1,
			facing: North,
			label:  "Moving North",
		},
		{
			x:      2,
			y:      2,
			xx:     1,
			facing: East,
			label:  "Moving East",
		},
		{
			x:      2,
			y:      2,
			yy:     3,
			facing: South,
			label:  "Moving South",
		},
		{
			x:      2,
			y:      2,
			xx:     3,
			facing: West,
			label:  "Moving West",
		},
	}
)

func TestValidateRobot(t *testing.T) {
	// validation only checks the compass point
	t.Run("A valid compass point to face", func(t *testing.T) {
		r := Robot{North}
		if err := r.validate(); err != nil {
			t.Log("Failed to create robot with valid compass point")
			t.Fail()
		}
	})

	t.Run("An invalid compass point to face", func(t *testing.T) {
		r := Robot{-1}
		if err := r.validate(); err == nil {
			t.Log("Validate allowed a robot with an invalid compass point")
			t.Fail()
		}
	})
}

func TestMove(t *testing.T) {
	// for this test we only have to ensure the appropriate x or y coordinate is adjusted correctly

	t.Run("Valid robot steps", func(t *testing.T) {
		for _, movement := range validRobotMovements {
			t.Run(movement.label, func(t *testing.T) {
				r := Robot{movement.facing}
				x, y := r.move(movement.x, movement.y)

				if x != movement.xx || y != movement.yy {
					t.Log("Robot did not know how to move")
					t.Fail()
				}
			})
		}
	})

	t.Run("Invalid robot steps", func(t *testing.T) {
		for _, movement := range invalidRobotMovements {
			t.Run(movement.label, func(t *testing.T) {
				r := Robot{movement.facing}
				x, y := r.move(movement.x, movement.y)

				if x == movement.xx || y == movement.yy {
					t.Log("Robot moved incorrectly")
					t.Fail()
				}
			})
		}
	})
}
