package creational

import "testing"

func TestBuilderPattern(t *testing.T) {
	ManufacturingComplex := ManufacturingDirector{}

	// type car
	carBuilder := &CarBuilder{}
	ManufacturingComplex.SetBuilder(carBuilder)
	ManufacturingComplex.Construct()

	car := carBuilder.Build()

	// check all criterias
	if car.Wheels != 4 {
		t.Errorf("wheels on a car must be 4 and they were %d\n", car.Wheels)
	}

	if car.Structure != "Car" {
		t.Errorf("structure on a car must be 'Car' and was %s\n", car.Structure)
	}

	if car.Seats != 5 {
		t.Errorf("seats on a car must be 5 they were %d\n", car.Seats)
	}

	// type bike
	bikeBuilder := &BikeBuilder{}

	ManufacturingComplex.SetBuilder(bikeBuilder)
	ManufacturingComplex.Construct()

	motorbike := bikeBuilder.GetVehicle()
	motorbike.Seats = 1

	if motorbike.Wheels != 2 {
		t.Errorf("wheels on a motorbike must be 2 and they were %d\n", motorbike.Wheels)
	}

	if motorbike.Structure != "Motorbike" {
		t.Errorf("structure on a motorbike must be 'Motorbike' and was %s\n", motorbike.Structure)
	}

	// type Bus
	BusBuilder := &BusBuilder{}
	ManufacturingComplex.SetBuilder(BusBuilder)
	ManufacturingComplex.Construct()

	bus := BusBuilder.GetVehicle()

	if bus.Wheels != 8 {
		t.Errorf("Wheels on a bus must be 8 and they were %d\n", bus.Wheels)
	}

	if bus.Structure != "Bus" {
		t.Errorf("Structure on a bus must be 'Bus' and was %s\n", bus.Structure)
	}

	if bus.Seats != 30 {
		t.Errorf("Seats on a bus must be 30 and they were %d\n", bus.Seats)
	}

}
