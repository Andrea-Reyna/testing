package prey

type preyStub struct {
	speed float64
}

func (p *preyStub) GetSpeed() float64 {
	return p.speed
}

func NewPreyStub(speed float64) Prey {
	return &preyStub{
		speed: speed,
	}
}
