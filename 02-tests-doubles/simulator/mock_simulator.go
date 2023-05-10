package simulator

// Manually
type CatchSimulatorMock struct {
	CanCatchFn          func(distance, speed, catchSpeed float64) bool
	GetLinearDistanceFn func(position [2]float64) float64
	Spys                map[string]bool
}

func (c *CatchSimulatorMock) CanCatch(distance, speed, catchSpeed float64) bool {
	// spy
	c.Spys["CanCatch"] = true
	return c.CanCatchFn(distance, speed, catchSpeed)
}

func (c *CatchSimulatorMock) GetLinearDistance(position [2]float64) float64 {
	// spy
	c.Spys["GetLinearDistance"] = true
	return c.GetLinearDistanceFn(position)
}
