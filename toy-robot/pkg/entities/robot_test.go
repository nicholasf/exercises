package entities

import "testing"

func Test_validate(t *testing.T) {

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
