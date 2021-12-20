package main

import "math"

type Vec4i struct {
	X int
	Y int
	Z int
	W int
}

func (v *Vec4i) Sub(other Vec4i) Vec4i {
	return Vec4i{
		X: v.X - other.X,
		Y: v.Y - other.Y,
		Z: v.Z - other.Z,
		W: v.W - other.W,
	}
}

func (v *Vec4i) DistanceSquare(other Vec4i) int {
	d := v.Sub(other)
	return d.X*d.X + d.Y*d.Y + d.Z*d.Z
}

func (v *Vec4i) DistanceManhattan(other Vec4i) int {
	d := v.Sub(other)
	return abs(d.X) + abs(d.Y) + abs(d.Z)
}

func (v *Vec4i) TranslateMatrix() Mat4i {
	return Mat4i{
		P00: 1, P01: 0, P02: 0, P03: v.X,
		P10: 0, P11: 1, P12: 0, P13: v.Y,
		P20: 0, P21: 0, P22: 1, P23: v.Z,
		P30: 0, P31: 0, P32: 0, P33: v.W,
	}
}

func (v *Vec4i) AsPoint() Vec4i {
	return Point3i(v.X, v.Y, v.Z)
}

func (v *Vec4i) AsVector() Vec4i {
	return Vec3i(v.X, v.Y, v.Z)
}

func Point3i(x, y, z int) Vec4i {
	return Vec4i{
		X: x,
		Y: y,
		Z: z,
		W: 1,
	}
}

func Vec3i(x, y, z int) Vec4i {
	return Vec4i{
		X: x,
		Y: y,
		Z: z,
		W: 0,
	}
}

type Mat4i struct {
	P00, P01, P02, P03 int
	P10, P11, P12, P13 int
	P20, P21, P22, P23 int
	P30, P31, P32, P33 int
}

func (m *Mat4i) MultiplyMatrix(other Mat4i) Mat4i {
	return Mat4i{
		P00: m.P00*other.P00 + m.P10*other.P01 + m.P20*other.P02 + m.P30*other.P03,
		P01: m.P00*other.P10 + m.P10*other.P11 + m.P20*other.P12 + m.P30*other.P13,
		P02: m.P00*other.P20 + m.P10*other.P21 + m.P20*other.P22 + m.P30*other.P23,
		P03: m.P00*other.P30 + m.P10*other.P31 + m.P20*other.P32 + m.P30*other.P33,

		P10: m.P01*other.P00 + m.P11*other.P01 + m.P21*other.P02 + m.P31*other.P03,
		P11: m.P01*other.P10 + m.P11*other.P11 + m.P21*other.P12 + m.P31*other.P13,
		P12: m.P01*other.P20 + m.P11*other.P21 + m.P21*other.P22 + m.P31*other.P23,
		P13: m.P01*other.P30 + m.P11*other.P31 + m.P21*other.P32 + m.P31*other.P33,

		P20: m.P02*other.P00 + m.P12*other.P01 + m.P22*other.P02 + m.P32*other.P03,
		P21: m.P02*other.P10 + m.P12*other.P11 + m.P22*other.P12 + m.P32*other.P13,
		P22: m.P02*other.P20 + m.P12*other.P21 + m.P22*other.P22 + m.P32*other.P23,
		P23: m.P02*other.P30 + m.P12*other.P31 + m.P22*other.P32 + m.P32*other.P33,

		P30: m.P03*other.P00 + m.P13*other.P01 + m.P23*other.P02 + m.P33*other.P03,
		P31: m.P03*other.P10 + m.P13*other.P11 + m.P23*other.P12 + m.P33*other.P13,
		P32: m.P03*other.P20 + m.P13*other.P21 + m.P23*other.P22 + m.P33*other.P23,
		P33: m.P03*other.P30 + m.P13*other.P31 + m.P23*other.P32 + m.P33*other.P33,
	}
}

func (m *Mat4i) MultiplyVec(other Vec4i) Vec4i {
	return Vec4i{
		X: m.P00*other.X + m.P10*other.Y + m.P20*other.Z + m.P30*1,
		Y: m.P01*other.X + m.P11*other.Y + m.P21*other.Z + m.P31*1,
		Z: m.P02*other.X + m.P12*other.Y + m.P22*other.Z + m.P32*1,
		W: m.P03*other.X + m.P13*other.Y + m.P23*other.Z + m.P33*1,
	}
}

var (
	MatrixRotateX90 = Mat4i{
		P00: 1, P01: 0, P02: 0, P03: 0,
		P10: 0, P11: 0, P12: -1, P13: 0,
		P20: 0, P21: 1, P22: 0, P23: 0,
		P30: 0, P31: 0, P32: 0, P33: 1,
	}

	MatrixRotateY90 = Mat4i{
		P00: 0, P01: 0, P02: 1, P03: 0,
		P10: 0, P11: 1, P12: 0, P13: 0,
		P20: -1, P21: 0, P22: 0, P23: 0,
		P30: 0, P31: 0, P32: 0, P33: 1,
	}

	MatrixRotateZ90 = Mat4i{
		P00: 0, P01: -1, P02: 0, P03: 0,
		P10: 1, P11: 0, P12: 0, P13: 0,
		P20: 0, P21: 0, P22: 1, P23: 0,
		P30: 0, P31: 0, P32: 0, P33: 1,
	}
)

func AllRotations(other Mat4i) []Mat4i {
	var rotations []Mat4i

	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			for z := 0; z < 3; z++ {
				matX := other.MultiplyMatrix(MatrixRotateX90)
				matY := other.MultiplyMatrix(MatrixRotateY90)
				matZ := other.MultiplyMatrix(MatrixRotateZ90)
				rotations = append(rotations, matX, matY, matZ)
			}
		}
	}

	return rotations
}

func abs(i int) int {
	return int(math.Abs(float64(i)))
}
