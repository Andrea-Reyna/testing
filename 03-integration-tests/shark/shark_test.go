package shark

import (
	"integrationtests/pkg/storage"
	"integrationtests/prey"
	"integrationtests/simulator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhiteShark_Hunt(t *testing.T) {
	t.Run("shark catches tuna - faster and close", func(t *testing.T) {
		//Arrange
		storageMock := &storage.StorageMock{
			GetValueFunc: func(key string) interface{} {
				switch key {
				case "white_shark_speed":
					return 5.0
				case "white_shark_x":
					return 250.0
				case "white_shark_y":
					return 250.0
				case "tuna_speed":
					return 2.0
				default:
					return nil
				}
			},
		}
		// Create a new CatchSimulator instance
		catchSim := simulator.NewCatchSimulator(120.0)
		// Create a new WhiteShark instance with the mocked Storage and CatchSimulator
		shark := CreateWhiteShark(catchSim, storageMock)
		// Define a Prey for the test
		prey := prey.CreateTuna(storageMock)

		//Act
		err := shark.Hunt(prey)

		//Assert
		assert.Nil(t, err)
	})
	t.Run("shark cannot catch tuna - slower", func(t *testing.T) {
		// Arrange
		storageMock := &storage.StorageMock{
			GetValueFunc: func(key string) interface{} {
				switch key {
				case "white_shark_speed":
					return 1.5 // shark more slower than tuna
				case "white_shark_x":
					return 250.0
				case "white_shark_y":
					return 250.0
				case "tuna_speed":
					return 2.0
				default:
					return nil
				}
			},
		}
		// Create a new CatchSimulator instance
		catchSim := simulator.NewCatchSimulator(120.0)
		// Create a new WhiteShark instance with the mocked Storage and CatchSimulator
		shark := CreateWhiteShark(catchSim, storageMock)
		// Define a Prey for the test
		prey := prey.CreateTuna(storageMock)

		//Act
		err := shark.Hunt(prey)

		//Assert
		assert.NotNil(t, err)
		assert.Equal(t, "could not hunt the prey", err.Error())
	})

	t.Run("shark cannot catch tuna - not enough time", func(t *testing.T) {
		// Arrange
		storageMock := &storage.StorageMock{
			GetValueFunc: func(key string) interface{} {
				switch key {
				case "white_shark_speed":
					return 5.0
				case "white_shark_x":
					return 250.0
				case "white_shark_y":
					return 250.0
				case "tuna_speed":
					return 2.0
				default:
					return nil
				}
			},
		}
		// Create a new CatchSimulator instance with not enough time to catch the prey
		catchSim := simulator.NewCatchSimulator(100.0)
		// Create a new WhiteShark instance with the mocked Storage and CatchSimulator
		shark := CreateWhiteShark(catchSim, storageMock)
		// Define a Prey for the test
		prey := prey.CreateTuna(storageMock)

		//Act
		err := shark.Hunt(prey)

		//Assert
		assert.NotNil(t, err)
		assert.Equal(t, "could not hunt the prey", err.Error())
	})
}

// Example With TestifyMock
func TestWhiteShark_HuntTestify(t *testing.T) {
	t.Run("the shark succedeed in hunting its prey", func(t *testing.T) {
		// Arrange
		sharkStorage := storage.NewMockStorageTestify()
		sharkStorage.Mock.On("GetValue", "white_shark_speed").Return(5.0)
		sharkStorage.Mock.On("GetValue", "white_shark_x").Return(250.0)
		sharkStorage.Mock.On("GetValue", "white_shark_y").Return(250.0)
		shark := CreateWhiteShark(simulator.NewCatchSimulator(120.0), sharkStorage)
		tunaStorage := storage.NewMockStorageTestify()
		tunaStorage.Mock.On("GetValue", "tuna_speed").Return(2.0)
		tuna := prey.CreateTuna(tunaStorage)
		// Act
		err := shark.Hunt(tuna)

		// Assert
		assert.NoError(t, err)
	})
	t.Run("shark cannot catch tuna - slower", func(t *testing.T) {
		// Arrange
		sharkStorage := storage.NewMockStorageTestify()
		sharkStorage.Mock.On("GetValue", "white_shark_speed").Return(1.5)
		sharkStorage.Mock.On("GetValue", "white_shark_x").Return(250.0)
		sharkStorage.Mock.On("GetValue", "white_shark_y").Return(250.0)
		shark := CreateWhiteShark(simulator.NewCatchSimulator(120.0), sharkStorage)
		tunaStorage := storage.NewMockStorageTestify()
		tunaStorage.Mock.On("GetValue", "tuna_speed").Return(2.0)
		tuna := prey.CreateTuna(tunaStorage)
		// Act
		err := shark.Hunt(tuna)

		//Assert
		assert.NotNil(t, err)
		assert.Equal(t, "could not hunt the prey", err.Error())
	})
	t.Run("the shark couldnt catch its prey", func(t *testing.T) {
		// Arrange
		sharkStorage := storage.NewMockStorageTestify()
		sharkStorage.Mock.On("GetValue", "white_shark_speed").Return(5.0)
		sharkStorage.Mock.On("GetValue", "white_shark_x").Return(250.0)
		sharkStorage.Mock.On("GetValue", "white_shark_y").Return(250.0)
		shark := CreateWhiteShark(simulator.NewCatchSimulator(100.0), sharkStorage)
		tunaStorage := storage.NewMockStorageTestify()
		tunaStorage.Mock.On("GetValue", "tuna_speed").Return(2.0)
		tuna := prey.CreateTuna(tunaStorage)
		// Act
		err := shark.Hunt(tuna)

		//Assert
		assert.NotNil(t, err)
		assert.Equal(t, "could not hunt the prey", err.Error())
	})
}
