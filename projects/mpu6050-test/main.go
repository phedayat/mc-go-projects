package main

import (
	"fmt"
	"machine"
	"time"

	"tinygo.org/x/drivers/mpu6050"
)

func main(){
	machine.I2C0.Configure(machine.I2CConfig{})
	sensor := mpu6050.New(machine.I2C0)
	sensor.Configure()

	connected := sensor.Connected()
	if !connected {
		fmt.Println("MPU6050 not found")
		return
	}
	fmt.Println("MPU6050 found")

	for {
		ax, ay, az := sensor.ReadAcceleration()
		rx, ry, rz := sensor.ReadRotation()

		s := fmt.Sprintf(
			"Acceleration: (%f, %f, %f), Rotation: (%f, %f, %f)", 
			ax, ay, az, rx, ry, rz,
		)
		fmt.Println(s)

		time.Sleep(2 * time.Second)
	}
}