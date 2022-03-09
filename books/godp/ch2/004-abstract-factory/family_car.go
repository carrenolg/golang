package abstractfactory

type FamiliarCar struct{}

func (l *FamiliarCar) NumDoors() int {
	return 5
}
func (l *FamiliarCar) NumWheels() int {
	return 4
}
func (l *FamiliarCar) NumSeats() int {
	return 5
}
