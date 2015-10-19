package appliances

//import errors to log errors when they occur
import "errors"

//The main interface used to describe appliances
type Appliance interface {
	Start()
	GetPurpose() string
}

//Our appliance types
const (
	STOVE = iota
	FRIDGE
	MICROWAVE
)

//Function to create the appliances
func CreateAppliance(t int) (Appliance, error) {
	//Use a switch case to switch between types, if a type exist then error is nil (null)
	switch t {
	case STOVE:
		return new(Stove), nil
	case FRIDGE:
		return new(Fridge), nil
	case MICROWAVE:
		return new(Microwave), nil
	default:
		//if type is invalid, return an error
		return nil, errors.New("Invalid Appliance Type")
	}
}
