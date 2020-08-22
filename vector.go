package vector

import "math"

// Point ..
// Assuming 3 dimensions for simplicity
type Point struct {
	X float64
	Y float64
	Z float64
}

// Vector ..
// to expand, convert Direction to
// > Components []int
type Vector struct {
	Length    float64
	Direction *Point
}

func (p *Point) Equals(p2 *Point) bool {
	if p.X == p2.X && p.Y == p2.Y && p.Z == p2.Z {
		return true
	}
	return false
}

func (p *Point) Length() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y + p.Z*p.Z)
}

func New(x, y, z float64) Vector {
	p := &Point{X: x, Y: y, Z: z}
	return Vector{
		Direction: p,
		Length:    p.Length(),
	}
}

func (v *Vector) Equals(v2 *Vector) bool {
	if v.Length == v2.Length && v.Direction.Equals(v2.Direction) {
		return true
	}
	return false
}

func (v *Vector) Add(v2 *Vector) Vector {
	p := &Point{
		X: v.Direction.X + v2.Direction.X,
		Y: v.Direction.Y + v2.Direction.Y,
		Z: v.Direction.Z + v2.Direction.Z,
	}

	return Vector{
		Direction: p,
		Length:    p.Length(),
	}
}

func (v *Vector) Subtract(v2 *Vector) Vector {
	p := &Point{
		X: v2.Direction.X - v.Direction.X,
		Y: v2.Direction.Y - v.Direction.Y,
		Z: v2.Direction.Z - v.Direction.Z,
	}

	return Vector{
		Direction: p,
		Length:    p.Length(),
	}
}

func (v *Vector) Multiply(val float64) Vector {
	return Vector{
		Length: v.Length * val,
	}
}

func (v *Vector) DotProduct(v2 *Vector) float64 {
	return v.Direction.X*v2.Direction.X + v.Direction.Y*v2.Direction.Y + v.Direction.Z*v2.Direction.Z
}

func (v *Vector) CrossProduct(v2 *Vector) Vector {
	i := v.Direction.Y*v2.Direction.Z - v.Direction.Z*v2.Direction.Y
	j := v.Direction.Z*v2.Direction.X - v.Direction.X*v2.Direction.Z
	k := v.Direction.X*v2.Direction.Y - v.Direction.Y*v2.Direction.X

	p := &Point{
		X: i,
		Y: j,
		Z: k,
	}

	return Vector{
		Direction: p,
		Length:    p.Length(),
	}
}
