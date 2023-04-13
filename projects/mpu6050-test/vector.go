package main

import (
	"fmt"
	"math"
)

type Vector struct {
	X, Y, Z float64
}

func (v Vector) Length() float64 {
	vx2 := math.Pow(v.X, 2)
	vy2 := math.Pow(v.Y, 2)
	vz2 := math.Pow(v.Z, 2)
	return math.Sqrt(vx2+vy2+vz2)
}

func (v Vector) PrettyPrint() string {
	return fmt.Sprintf("(%f, %f, %f)", v.X, v.Y, v.Z)
}