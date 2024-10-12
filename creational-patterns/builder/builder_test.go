package builder

import "testing"

func TestBuilderPattern(t *testing.T) {
	manufacturingDirector := ManufacturingDirector{}

	carBuilder := &CarBuilder{}
	manufacturingDirector.SetBuilder(carBuilder)
	manufacturingDirector.Construct()

	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Errorf("Wheels on a car must be 4 and they were %d\n", car.Wheels)
	}

	if car.Structure != "Car" {
		t.Errorf("Structure on a car must be 'Car' and was %s\n", car.Structure)
	}

	if car.Seats != 5 {
		t.Errorf("Seats on a car must be 5 and they wre %d\n", car.Seats)
	}

	bikeBuilder := &BikeBuilder{}

	manufacturingDirector.SetBuilder(bikeBuilder)
	manufacturingDirector.Construct()

	motorbike := bikeBuilder.GetVehicle()

	if motorbike.Seats != 2 {
		t.Errorf("Seats on a motorbike must be 2 and the were %d\n", motorbike.Seats)
	}

	if motorbike.Wheels != 2 {
		t.Errorf("Wheels on a motorbike must be 2 and they were %d\n", motorbike.Wheels)
	}

	if motorbike.Structure != "Motorbike" {
		t.Errorf("Structure on a motorbike must be 'Motorbike' and was %s\n", motorbike.Structure)
	}
}
