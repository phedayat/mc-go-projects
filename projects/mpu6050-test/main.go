package main

import (
	"os"
	"fmt"
	"time"
	"machine"
	"tinygo.org/x/drivers/mpu6050"
)

func check(err error, stage string){
	if err != nil {
		fmt.Println(stage, "--->", err.Error())
		os.Exit(1)
	}
}

func main(){
	// First we get the SDA, SCL pins
	sda := machine.GPIO0
	scl := machine.GPIO1

	// Then we configure the I2C0 constant
	err := machine.I2C0.Configure(machine.I2CConfig{SDA: sda, SCL: scl})
	check(err, "I2C0 Configuration")

	// Now we can get a new instance of the MPU6050 device
	// We pass in the I2C bus (machine.I2C0)
	sensor := mpu6050.New(machine.I2C0)
	err = sensor.Configure() // run .Configure() to prepare the device
	check(err, "Sensor Configuration")

	// Checks that a device is connected
	connected := sensor.Connected()
	if !connected {
		fmt.Println("MPU6050 not found")
		return
	}
	fmt.Println("MPU6050 found")

	sa := int16(16384)
	st := int16(340)
	ot := float32(35)
	for {
		accel := make([]byte, 6)
		realAccel := make([]float32, 3)

		temp := make([]byte, 2)
		
		machine.I2C0.ReadRegister(ADDR, uint8(0x3B), accel)
		machine.I2C0.ReadRegister(ADDR, uint8(0x41), temp)

		realAccel[0] = GetReading(Btoi16(accel[0], accel[1]), sa)
		realAccel[1] = GetReading(Btoi16(accel[2], accel[3]), sa)
		realAccel[2] = GetReading(Btoi16(accel[4], accel[5]), sa)

		realTemp := GetReading(Btoi16(temp[0], temp[1]), st) + ot


		fmt.Println("--------------------START--------------------")
		fmt.Println("Accel: ", realAccel )
		fmt.Println("Temp: ", realTemp, " C")
		fmt.Println("--------------------END--------------------")

		time.Sleep(1*time.Second)
	}
}

// FUNCTIONS FOR HANDLING READING DATA

func Btoi16(b1, b2 byte) int16 {
	return int16((uint16(b1)<<8)|uint16(b2))
}

func GetReading(i int16, s int16) float32 {
	return float32(i) / float32(s)
}
