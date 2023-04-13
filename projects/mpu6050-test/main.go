package main

import (
	"fmt"
	"time"
	"machine"
	"tinygo.org/x/drivers/mpu6050"
)

func main(){
	// First we get the SDA, SCL pins
	sda := machine.GPIO0
	scl := machine.GPIO1

	// Then we configure the I2C0 constant
	machine.I2C0.Configure(machine.I2CConfig{SDA: sda, SCL: scl})
	
	// Now we can get a new instance of the MPU6050 device
	// We pass in the I2C bus (machine.I2C0)
	sensor := mpu6050.New(machine.I2C0)
	sensor.Configure() // run .Configure() to prepare the device

	// Checks that a device is connected
	connected := sensor.Connected()
	if !connected {
		fmt.Println("MPU6050 not found")
		return
	}
	fmt.Println("MPU6050 found")

	for {
		ax, ay, az := sensor.ReadAcceleration() // get accel coords
		rx, ry, rz := sensor.ReadRotation() // get rotation coords

		av := Vector{float64(ax), float64(ay), float64(az)} // create accel vector
		rv := Vector{float64(rx), float64(ry), float64(rz)} // create rotation vector

		// Create the formatted string for printing
		s := fmt.Sprintf(
			"Acceleration: %s (Length: %f), Rotation: %s (Length: %f)", 
			av.PrettyPrint(), av.Length(), rv.PrettyPrint(), rv.Length(),
		)
		fmt.Println(s)

		// Sleep
		time.Sleep(2 * time.Second)
	}
}