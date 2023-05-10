package shark

import (
	"github.com/stretchr/testify/assert"
	"testdoubles/prey"
	"testdoubles/simulator"
	"testing"
)

func TestWhiteShark_Hunt(t *testing.T) {
	t.Run("shark catches tuna - faster and close", func(t *testing.T) {
		//arrange
		s := &simulator.CatchSimulatorMock{
			CanCatchFn: func(distance, speed, catchSpeed float64) bool {
				return true
			},
			GetLinearDistanceFn: func(position [2]float64) float64 {
				return 10.0
			},
			Spys: make(map[string]bool),
		}
		p := prey.NewPreyStub(50)
		shark := CreateWhiteShark(s)

		//act
		err := shark.Hunt(p)

		//assert
		assert.Nil(t, err)
		assert.True(t, s.Spys["CanCatch"])
		assert.True(t, s.Spys["GetLinearDistance"])
	})

	t.Run("shark can't catch tuna - slower", func(t *testing.T) {
		//Arrange
		s := &simulator.CatchSimulatorMock{
			CanCatchFn: func(distance, speed, catchSpeed float64) bool {
				return false
			},
			GetLinearDistanceFn: func(position [2]float64) float64 {
				return 10.0
			},
			Spys: make(map[string]bool),
		}
		p := prey.NewPreyStub(100)
		shark := CreateWhiteShark(s)
		expectedError := "could not hunt the prey"

		//Act
		err := shark.Hunt(p)

		//Assert
		assert.NotNil(t, err)
		assert.Equal(t, expectedError, err.Error())
		assert.True(t, s.Spys["CanCatch"])
		assert.True(t, s.Spys["GetLinearDistance"])
	})

	t.Run("shark can't catch tuna - too far", func(t *testing.T) {
		//Arrange
		s := &simulator.CatchSimulatorMock{
			CanCatchFn: func(distance, speed, catchSpeed float64) bool {
				return false
			},
			GetLinearDistanceFn: func(position [2]float64) float64 {
				return 500.0
			},
			Spys: make(map[string]bool),
		}
		p := prey.NewPreyStub(50)
		shark := CreateWhiteShark(s)
		expectedError := "could not hunt the prey"

		//Act
		err := shark.Hunt(p)

		//Assert
		assert.NotNil(t, err)
		assert.Equal(t, expectedError, err.Error())
		assert.True(t, s.Spys["CanCatch"])
		assert.True(t, s.Spys["GetLinearDistance"])
	})
}
