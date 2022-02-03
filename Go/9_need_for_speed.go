/* Instructions
In this exercise you'll be organizing races between various types of remote controlled cars. Each car has its own speed and battery drain characteristics.

Cars start with full (100%) batteries. Each time you drive the car using the remote control, it covers the car's speed in meters and decreases the remaining battery percentage by its battery drain.

If a car's battery is below its battery drain percentage, you can't drive the car anymore.

Each race track has its own distance. Cars are tested by checking if they can finish the track without running out of battery.

1. Creating a remote controlled car
Define a Car struct with the following int type fields:

battery
batteryDrain
speed
distance
Allow creating a remote controlled car by defining a function NewCar that takes the speed of the car in meters, and the battery drain percentage as its two parameters (both of type int) and returns a Car instance:

speed := 5
batteryDrain := 2
car := NewCar(speed, batteryDrain)
// Output: Car{speed: 5, batteryDrain: 2, battery:100, distance: 0}
2. Creating a race track
Define another struct type called Track with the field distance of type integer. Allow creating a race track by defining a function NewTrack that takes the track's distance in meters as its sole parameter (which is of type int):

distance := 800
raceTrack := NewTrack(distance)
// Output: Track{distance: 800}
3. Drive the car
Implement the Drive function that updates the number of meters driven based on the car's speed, and reduces the battery according to the battery drainage:

speed := 5
batteryDrain := 2
car := NewCar(speed, batteryDrain)
car = Drive(car)
// Output: Car{speed: 5, batteryDrain: 2, battery: 98, distance: 5}
4. Check if a remote controlled car can finish a race
To finish a race, a car has to be able to drive the race's distance. This means not draining its battery before having crossed the finish line. Implement the CanFinish function that takes a Car and a Track instance as its parameter and returns true if the car can finish the race; otherwise, return false. Assume the car is just starting the race:

speed := 5
batteryDrain := 2
car := NewCar(speed, batteryDrain)

distance := 100
raceTrack := NewTrack(distance)

CanFinish(car, raceTrack)
// Output: true */

package speed

// TODO: define the 'Car' type struct
type Car struct {
	battery, batteryDrain, speed, distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	car := Car{
		battery:      100,
		speed:        speed,
		batteryDrain: batteryDrain,
	}
	return car
}

type Track struct {
	distance int
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	track := Track{
		distance: distance,
	}
	return track
}

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move.
func Drive(car Car) Car {
	if car.battery > car.batteryDrain {
		car.distance = car.speed
		car.battery = car.battery - car.batteryDrain
	}
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	max_distance := ((car.battery / car.batteryDrain) * car.speed) + car.distance
	return max_distance >= track.distance
}
