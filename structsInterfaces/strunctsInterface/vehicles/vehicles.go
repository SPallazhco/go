package vehicles

import "fmt"

type Car struct {
	Time int
}

func (car *Car) Distance() float64 {
	return 100 * (float64(car.Time) / 60)
}

type Motorcycle struct {
	Time int
}

func (motorcycle *Motorcycle) Distance() float64 {
	return 120 * (float64(motorcycle.Time) / 60)
}

type Truck struct {
	Time int
}

func (truck *Truck) Distance() float64 {
	return 70 * (float64(truck.Time) / 60)
}

type Avion struct {
	Time int
}

func (truck *Avion) Distance() float64 {
	return 30000 * (float64(truck.Time) / 60)
}

type Vehicle interface {
	Distance() float64
}

const (
	CarVehicle        = "CAR"
	MotorcycleVehicle = "MOTORCICLE"
	TruckVehicle      = "TRUCK"
	AvionVehicle      = "AVION"
)

func New(v string, time int) (Vehicle, error) {
	switch v {
	case CarVehicle:
		return &Car{Time: time}, nil
	case MotorcycleVehicle:
		return &Motorcycle{Time: time}, nil
	case TruckVehicle:
		return &Truck{Time: time}, nil
	case AvionVehicle:
		return &Avion{Time: time}, nil
	}
	return nil, fmt.Errorf("Vehicle '%s' not exist", v)
}
