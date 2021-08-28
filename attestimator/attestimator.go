package main

import (
	"fmt"
	"math"

	"github.com/knei-knurow/attestimator"
)

func main() {
	var est attestimator.Estimator
	est.ResetAll(true)

	for {
		dt := 0.02
		a, g, m := [3]float64{1, 2, 3}, [3]float64{4, 5, 6}, [3]float64{}

		fmt.Scanf("%f %f %f %f %f %f\n", &a[0], &a[1], &a[2], &g[0], &g[1], &g[2])

		for i := 0; i < 3; i++ {
			g[i] /= 131.0        // rescale MPU-6050 raw values
			g[i] *= 0.0174532925 // deg to rad
		}

		est.Update(dt, g[0], g[1], g[2], a[0], a[1], a[2], m[0], m[1], m[2])

		w, x, y, z := est.GetAttitude()
		fmt.Printf("%f\t%f\t%f\t%f\n", math.Acos(w)*2*57.2957795, x, y, z)

	}
}
