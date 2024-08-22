package models

type Engine struct {
	Power        int     // horsepower
	Displacement float64 // engine displacement in liters
	Type         string
	IsRunning    bool // engine state
}

// Transmission represents the transmission of the car
type Transmission struct {
	Type    string // "manual", "automatic"
	Gear    int    // current gear
	MaxGear int    // maximum gear
}

// FuelTank represents the fuel tank of the car
type FuelTank struct {
	Capacity float64 // maximum fuel capacity in liters
	Level    float64 // current fuel level in liters
}

// Car represents the car itself
type Car struct {
	Make         string        // car brand
	Model        string        // car models
	Speed        float64       // current speed of the car in km/h
	Odometer     float64       // total distance traveled in km
	Engine       *Engine       // car engine
	Transmission *Transmission // car transmission
	FuelTank     *FuelTank     // fuel tank
}
