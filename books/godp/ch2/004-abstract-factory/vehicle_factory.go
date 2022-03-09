package abstractfactory

import "fmt"

const (
	CarFactoryType      = 1
	MotorbikeFactorType = 2
)

type VehicleFactory interface {
	Build(v int) (Vehicle, error)
}

func BuildFactory(f int) (VehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return new(CarFactory), nil
	case MotorbikeFactorType:
		return new(MotorBikeFactory), nil
	default:
		return nil, fmt.Errorf("factory with id %d not recognized", f)
	}
}
