package plane

import "math"

type Vector struct {
	X, Y float64
}

func (v Vector) Angle() float64 {
	return math.Atan2(v.Y, v.X)
}

func (v Vector) Mag() float64 {
	return math.Sqrt((v.X * v.X) + (v.Y * v.Y))
}

func (v Vector) Unit() Vector {
	if (v.X == 0) && (v.Y == 0) {
		return Vector{1, 0}
	}

	return v.Scaled(1 / v.Mag())
}

func (v Vector) Scaled(f float64) Vector {
	return Vector{
		X: v.X * f,
		Y: v.Y * f,
	}
}
