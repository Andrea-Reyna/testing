package hunt

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestSharkHuntsSuccessfully checks if the shark hunts successfully.
func TestSharkHuntsSuccessfully(t *testing.T) {
	t.Run("Shark hunts successfully", func(t *testing.T) {
		//arrange
		shark := &Shark{hungry: true, tired: false, speed: 10}
		prey := &Prey{name: "Fish", speed: 3}
		var expectedResult error = nil

		//act
		err := shark.Hunt(prey)

		//assert
		assert.Nil(t, err)
		assert.Equal(t, err, expectedResult)
		assert.False(t, shark.hungry)
		assert.True(t, shark.tired)
	})
}

// TestSharkCannotHuntBecauseIsTired checks if the shark cannot hunt because it is tired.
func TestSharkCannotHuntBecauseIsTired(t *testing.T) {
	t.Run("Shark cannot hunt because is tired", func(t *testing.T) {
		// arrange
		shark := &Shark{tired: true, hungry: true, speed: 10}
		prey := &Prey{name: "Fish", speed: 5}
		expectedError := errors.New("cannot hunt, i am really tired")

		// act
		err := shark.Hunt(prey)

		// assert
		assert.NotNil(t, err)
		assert.Equal(t, expectedError, err)
	})
}

// TestSharkCannotHuntBecauseIsNotHungry checks if the shark cannot hunt because it is not hungry.
func TestSharkCannotHuntBecauseIsNotHungry(t *testing.T) {
	t.Run("Shark cannot hunt because is not hungry", func(t *testing.T) {
		// arrange
		shark := &Shark{tired: false, hungry: false, speed: 10}
		prey := &Prey{name: "Fish", speed: 5}
		expectedError := errors.New("cannot hunt, i am not hungry")

		// act
		err := shark.Hunt(prey)

		// assert
		assert.NotNil(t, err)
		assert.Equal(t, expectedError, err)
	})
}

// TestSharkCannotReachThePrey checks if the shark cannot reach the prey.
func TestSharkCannotReachThePrey(t *testing.T) {
	t.Run("Shark cannot reach the prey", func(t *testing.T) {
		// arrange
		shark := &Shark{tired: false, hungry: true, speed: 5}
		prey := &Prey{name: "FastFish", speed: 10}
		expectedError := errors.New("could not catch it")
		// act
		err := shark.Hunt(prey)

		// assert
		assert.NotNil(t, err)
		assert.Equal(t, expectedError, err)
		assert.True(t, shark.tired)
	})
}

// TestSharkHuntNilPrey checks if the shark's Hunt method returns an error when given a nil prey.
func TestSharkHuntNilPrey(t *testing.T) {
	t.Run("Shark hunts a nil prey", func(t *testing.T) {
		// arrange
		shark := &Shark{tired: false, hungry: true, speed: 10}
		var prey *Prey
		expectedError := errors.New("cannot hunt, prey is nil")

		// act
		err := shark.Hunt(prey)

		// assert
		assert.NotNil(t, err)
		assert.Equal(t, expectedError, err)
	})
}

// TDT example
func TestSharkHunt(t *testing.T) {
	tests := []struct {
		name          string
		shark         *Shark
		prey          *Prey
		expectedError error
	}{
		{
			name:          "Shark hunts successfully",
			shark:         &Shark{tired: false, hungry: true, speed: 10},
			prey:          &Prey{name: "Fish", speed: 5},
			expectedError: nil,
		},
		{
			name:          "Shark cannot hunt because is tired",
			shark:         &Shark{tired: true, hungry: true, speed: 10},
			prey:          &Prey{name: "Fish", speed: 5},
			expectedError: errors.New("cannot hunt, i am really tired"),
		},
		{
			name:          "Shark cannot hunt because is not hungry",
			shark:         &Shark{tired: false, hungry: false, speed: 10},
			prey:          &Prey{name: "Fish", speed: 5},
			expectedError: errors.New("cannot hunt, i am not hungry"),
		},
		{
			name:          "Shark cannot reach the prey",
			shark:         &Shark{tired: false, hungry: true, speed: 5},
			prey:          &Prey{name: "FastFish", speed: 10},
			expectedError: errors.New("could not catch it"),
		},
		{
			name:          "Shark hunts a nil prey",
			shark:         &Shark{tired: false, hungry: true, speed: 10},
			prey:          nil,
			expectedError: errors.New("cannot hunt, prey is nil"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// act
			err := tt.shark.Hunt(tt.prey)

			// assert
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
