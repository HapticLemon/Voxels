package main

import (
	"image/color"
)

type Voxel struct {
	color color.RGBA
	densidad  float64
	transmisividad float64
}

const MAXDIM int = 20
const VOXELSIZE float64 = 1

// En realidad debería ser el valor que tiene en C++
const DBL_MAX float64 = 10000

var voxelGrid [MAXDIM][MAXDIM][MAXDIM]Voxel
