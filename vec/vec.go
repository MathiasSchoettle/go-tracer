package vec

import "math"

type Vector3 struct {
	x float64
	y float64
	z float64
}

func (v *Vector3) X() float64 {
	return v.x
}

func (v *Vector3) Y() float64 {
	return v.y
}

func (v *Vector3) Z() float64 {
	return v.z
}

func (v *Vector3) SetX(value float64) {
	v.x = value
}

func Of(x float64, y float64, z float64) Vector3 {
	return Vector3{x, y, z}
}

func Unit() Vector3 {
	return Of(1, 1, 1)
}

func (v *Vector3) Length() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

func (v *Vector3) Scale(s float64) *Vector3 {
	v.x *= s
	v.y *= s
	v.z *= s

	return v
}

func Scaled(v Vector3, s float64) Vector3 {
	return Vector3{
		v.x * s,
		v.y * s,
		v.z * s,
	}
}

func (v *Vector3) Divide(s float64) *Vector3 {
	v.x /= s
	v.y /= s
	v.z /= s

	return v
}

func Divided(v Vector3, s float64) Vector3 {
	return Vector3{
		v.x / s,
		v.y / s,
		v.z / s,
	}
}

func (v *Vector3) Add(other Vector3) *Vector3 {
	v.x += other.x
	v.y += other.y
	v.z += other.z

	return v
}

func Added(first Vector3, other Vector3) Vector3 {
	return Vector3{
		first.x + other.x,
		first.y + other.y,
		first.z + other.z,
	}
}

func (v *Vector3) Subtract(other Vector3) *Vector3 {
	v.x -= other.x
	v.y -= other.y
	v.z -= other.z

	return v
}

func Subtracted(first Vector3, other Vector3) Vector3 {
	return Vector3{
		first.x - other.x,
		first.y - other.y,
		first.z - other.z,
	}
}

func (v *Vector3) Normalize() *Vector3 {
	length := v.Length()
	v.x /= length
	v.y /= length
	v.z /= length

	return v
}

func Normalized(v Vector3) Vector3 {
	length := v.Length()
	return Vector3{
		v.x / length,
		v.y / length,
		v.z / length,
	}
}

func Dot(u Vector3, v Vector3) float64 {
	return u.x*v.x + u.y*v.y + u.z*v.z
}

func Cross(u Vector3, v Vector3) Vector3 {
	return Vector3{
		u.y*v.z - u.z*v.y,
		u.z*v.x - u.x*v.z,
		u.x*v.y - u.y*v.x,
	}
}
