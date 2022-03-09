package abstractfactory

import "fmt"

const (
	SportMotorbikeType  = 1
	CruiseMotorbikeType = 2
)

type MotorBikeFactory struct{}

func (m *MotorBikeFactory) Build(v int) (Vehicle, error) {
	switch v {
	case SportMotorbikeType:
		return new(SportMotorbike), nil
	case CruiseMotorbikeType:
		return new(CruiseMotorbike), nil
	default:
		return nil, fmt.Errorf("vehicle of the type %d not recognized", v)
	}
}
