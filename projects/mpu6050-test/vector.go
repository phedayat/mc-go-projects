package main

import (
	"fmt"
	"math"
)

type Vector struct {
	X, Y, Z float32
}

func (v Vector) Length() float64 {
	x := float64(v.X)
	y := float64(v.Y)
	z := float64(v.Z)
	vx2 := math.Pow(x, 2)
	vy2 := math.Pow(y, 2)
	vz2 := math.Pow(z, 2)
	return math.Sqrt(vx2+vy2+vz2)
}

func (v Vector) PrettyPrint() string {
	return fmt.Sprintf("(%f, %f, %f)", v.X, v.Y, v.Z)
}