package models

type Engine struct {
	Power        int     // horsepower
	Displacement float64 // engine displacement in liters
	Type         string  // engine type ("L4" , "V6", "V8", "w16")
	IsRunning    bool    // engine state
}

// Transmission represents the transmission of the car
type Transmission struct {
	Type    string // "manual", "automatic"
	Gear    int    // current gear
	MaxGear int    // maximum gear
}

// Wheels represents the wheels of the car
type Wheels struct {
	Size int    // wheel size in inches
	Type string // tire type (e.g., "All-Season", "Sport", "Off-Road")
}

// FuelTank represents the fuel tank of the car
type FuelTank struct {
	Capacity float64 // maximum fuel capacity in liters
	Level    float64 // current fuel level in liters
}

// Car represents the car itself
type Car struct {
	Make         string       // car brand
	Model        string       // car models
	Year         int          // car manufacturing year
	Speed        float64      // current speed of the car in km/h
	Odometer     float64      // total distance traveled in km
	Engine       Engine       // car engine
	Transmission Transmission // car transmission
	Wheels       [4]Wheels    // array representing the four wheels
	FuelTank     FuelTank     // fuel tank
}
