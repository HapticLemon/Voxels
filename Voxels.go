package main

import (
	"./Ruido"
	"./Vectores"
	"fmt"
	"image/color"
	"math"
)

// Implementación sacada de :
// https://github.com/slsdo/volumetric-cloud/blob/master/VolumetricClouds/VolumeRender.cpp
//
func generaVoxelGrid(){
	var center = Vectores.Vector{float64(MAXDIM) * 0.5, float64(MAXDIM) * 0.5, float64(MAXDIM) * 0.5}

	var max_distance float64 = (Vectores.Vector{0.5,0.5,0.5}.Sub(center)).Length()
	var max_ratio float64 = 1/max_distance;

	for kk := 0; kk < MAXDIM; kk++ {
		for jj := 0; jj < MAXDIM; jj++ {
			for ii := 0; ii < MAXDIM; ii++ {
				var cloud float64 = Ruido.Noise3(float64(ii) * VOXELSIZE, float64(jj) * VOXELSIZE, float64(kk) * VOXELSIZE)

				var voxel = Vectores.Vector{float64(ii) + 0.5, float64(jj) + 0.5, float64(kk) + 0.5}

				// Distancia del voxel actual hasta el centro.
				var distance float64 = voxel.Sub(center).Length()

				// Cantidad de "cubierto".
				var cover float64 = distance * max_ratio + 0.3

				// Fuzziness de las nubes.
				var sharpness float64 = 0.5

				// Densidad de la nube.
				var density float64 = 5

				cloud = cloud - cover
				if cloud < 0 {
					cloud = 0
				}else {
					cloud = cloud * density
				}

				cloud = 1 - math.Pow(sharpness, cloud)

				voxelGrid[ii][jj][kk].densidad = cloud
				// En el original las componenetes de color son 1
				voxelGrid[ii][jj][kk].color = color.RGBA{R: 255, G: 255, B: 255, A: 255}

			}
		}
	}

}

// Implementación de https://github.com/francisengelmann/fast_voxel_traversal/blob/master/main.cpp
//
func voxelTransversal(inicio Vectores.Vector, fin Vectores.Vector) []Vectores.Vector{
	var currentVoxel = Vectores.Vector{math.Floor(inicio.X / VOXELSIZE), math.Floor(inicio.Y / VOXELSIZE),math.Floor(inicio.Z / VOXELSIZE)}
	var lastVoxel = Vectores.Vector{math.Floor(fin.X / VOXELSIZE), math.Floor(fin.Y / VOXELSIZE),math.Floor(fin.Z / VOXELSIZE)}

	var visited_voxels []Vectores.Vector
	var rayo = fin.Sub(inicio)

	var stepX float64 = 0
	var stepY float64 = 0
	var stepZ float64 = 0

	if rayo.X >= 0 { stepX = 1 	} else {stepX = -1}

	if rayo.Y >= 0 { stepY = 1	} else {stepY = -1}

	if rayo.Z >= 0 { stepZ = 1 	} else {stepZ = -1}

	var next_voxel_boundary_x float64 = (currentVoxel.X + stepX) * VOXELSIZE
	var next_voxel_boundary_y float64 = (currentVoxel.Y + stepY) * VOXELSIZE
	var next_voxel_boundary_z float64 = (currentVoxel.Z + stepY) * VOXELSIZE

	var tMaxX float64 = 0
	var tMaxY float64 = 0
	var tMaxZ float64 = 0

	if rayo.X != 0 {tMaxX = (next_voxel_boundary_x - inicio.X) / rayo.X} else {tMaxX = DBL_MAX}
	if rayo.Y != 0 {tMaxY = (next_voxel_boundary_y - inicio.Y) / rayo.Y} else {tMaxY = DBL_MAX}
	if rayo.Z != 0 {tMaxZ = (next_voxel_boundary_z - inicio.Z) / rayo.Z} else {tMaxZ = DBL_MAX}

	var tDeltaX float64 = 0
	var tDeltaY float64 = 0
	var tDeltaZ float64 = 0

	if rayo.X != 0 { tDeltaX = VOXELSIZE/rayo.X * stepX} else { tDeltaX = DBL_MAX}
	if rayo.Y != 0 { tDeltaY = VOXELSIZE/rayo.Y * stepY} else { tDeltaY = DBL_MAX}
	if rayo.Z != 0 { tDeltaZ = VOXELSIZE/rayo.Z * stepZ} else { tDeltaZ = DBL_MAX}

	var diff = Vectores.Vector{0,0,0}
	var neg_ray bool = false

	if currentVoxel.X != lastVoxel.X && rayo.X < 0 {
		diff.X--
		neg_ray = true
	}
	if currentVoxel.Y != lastVoxel.Y && rayo.Y < 0 {
		diff.Y--
		neg_ray = true
	}
	if currentVoxel.Y != lastVoxel.Y && rayo.Y < 0 {
		diff.Y--
		neg_ray = true
	}

	visited_voxels = append(visited_voxels, currentVoxel)

	if (neg_ray) {
		currentVoxel.Add(diff)
		visited_voxels = append(visited_voxels, currentVoxel)
	}

	for lastVoxel != currentVoxel{
		if (tMaxX < tMaxY) {
			if (tMaxX < tMaxZ) {
				currentVoxel.X += stepX
				tMaxX += tDeltaX
			} else {
				currentVoxel.Z += stepZ
				tMaxZ += tDeltaZ
			}
		}else {
			if (tMaxY < tMaxZ) {
				currentVoxel.Y += stepY
				tMaxY += tDeltaY
			} else {
				currentVoxel.Z += stepZ
				tMaxZ += tDeltaZ
			}
		}
		visited_voxels = append(visited_voxels, currentVoxel)
	}
	return visited_voxels
}

func main(){
	var inicio = Vectores.Vector{0,0,0}

	// Éste habrá que calcularlo de alguna manera.
	var fin = Vectores.Vector{3,2,2}

	//var voxelList []Vectores.Vector

	generaVoxelGrid()
	voxelList := voxelTransversal(inicio, fin)

	for e := 0; e < len(voxelList); e++ {
		fmt.Println("gateTErlZ")
	}
}
