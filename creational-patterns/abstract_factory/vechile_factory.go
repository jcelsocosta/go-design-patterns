package abstractfactory

import (
	"fmt"
)

type VehicleFactory interface {
	GetVehicle(v int) (Vehicle, error)
}

const (
	CarFactoryType      = 1
	MotobikeFactoryType = 2
)

func GetVehicleFactory(f int) (VehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return new(CarFactory), nil
	case MotobikeFactoryType:
		return new(MotorbikeFactory), nil
	default:
		return nil, fmt.Errorf("factory with id %d not recognized", f)
	}
}
