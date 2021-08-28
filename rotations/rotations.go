package main

import (
	"fmt"
	"math"
)

func VNormalize(v [3]float64) [3]float64 {
	mag2 := v[0]*v[0] + v[1]*v[1] + v[2]*v[2]
	const tolerance = 0.00001
	if math.Abs(mag2-1.0) > tolerance {
		mag := math.Sqrt(mag2)
		for idx, _ := range v {
			v[idx] /= mag
		}
	}
	return v
}

func QMult(q1 [4]float64, q2 [4]float64) [4]float64 {
	w1, x1, y1, z1 := q1[0], q1[1], q1[2], q1[3]
	w2, x2, y2, z2 := q2[0], q2[1], q2[2], q2[3]
	w := w1*w2 - x1*x2 - y1*y2 - z1*z2
	x := w1*x2 + x1*w2 + y1*z2 - z1*y2
	y := w1*y2 + y1*w2 + z1*x2 - x1*z2
	z := w1*z2 + z1*w2 + x1*y2 - y1*x2
	return [4]float64{w, x, y, z}
}

func QConjugate(q [4]float64) [4]float64 {
	return [4]float64{q[0], -q[1], -q[2], -q[3]}
}

func QVMult(q1 [4]float64, v [3]float64) [3]float64 {
	q2 := [4]float64{0, v[0], v[1], v[2]}
	w := QMult(QMult(q1, q2), QConjugate(q1))
	return [3]float64{w[1], w[2], w[3]}
}

func AxisAngleToQ(v [3]float64, theta float64) [4]float64 {
	v = VNormalize(v)
	x, y, z := v[0], v[1], v[2]
	theta /= 2
	w := math.Cos(theta)
	x = x * math.Sin(theta)
	y = y * math.Sin(theta)
	z = z * math.Sin(theta)
	return [4]float64{w, x, y, z}
}

func main() {
	// More details here (Python): https://stackoverflow.com/a/4870905/7389107
	// Online calculator: https://www.vcalc.com/wiki/vCalc/V3+-+Vector+Rotation

	v := [3]float64{0, 0, 1}

	rotation := AxisAngleToQ(VNormalize([3]float64{0, 1, 1}), math.Pi/2)

	v = QVMult(rotation, v)
	fmt.Println(v)
}
