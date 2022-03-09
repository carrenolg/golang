package abstractfactory

import (
	"fmt"
)

const (
	LuxuryCarType = 1
	FamilyCarType = 2
)

type CarFactory struct{}

func (c *CarFactory) Build(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamiliarCar), nil
	default:
		return nil, fmt.Errorf("vehicle of the type %d not recognized", v)
	}
}
