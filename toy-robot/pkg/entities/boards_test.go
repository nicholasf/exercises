package entities

import (
	"testing"
)

func TestValidate(t *testing.T) {
	validBoards := []Board{
		{
			ID:     1,
			RobotX: 0,
			RobotY: 0,
		},
		{
			ID:     1,
			RobotX: 4,
			RobotY: 4,
		},
	}

	invalidBoards := []Board{
		{
			ID:     1,
			RobotX: -1,
			RobotY: -1,
		},
		{
			ID:     1,
			RobotX: 5,
			RobotY: 5,
		},
	}

	t.Run("Valid boards", func(t *testing.T) {
		for _, board := range validBoards {
			err := board.validate()

			if err != nil {
				t.Log("Valid board should be valid!")
				t.Fail()
			}
		}
	})

	t.Run("Invalid boards", func(t *testing.T) {
		for _, board := range invalidBoards {
			err := board.validate()

			if err == nil {
				t.Log("Invalid board should be invalid!")
				t.Fail()
			}
		}
	})

}
