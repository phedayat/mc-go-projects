package main

import (
	"machine"
)

const (
	ADDR = uint8(0x68)
	CLOCK_INTERNAL = 0x00
	PWR_MGMT_1 = 0x6B
)

type Device struct {
	Bus machine.I2C
	Addr uint8
}

func New(bus machine.I2C) Device {
	return Device{bus, ADDR}
}

func (d Device) Configure() error {
	return d.SetClockSource(CLOCK_INTERNAL)
}

func (d Device) SetClockSource(source uint8) error {
	return d.Bus.WriteRegister(d.Addr, PWR_MGMT_1, []uint8{source})
}