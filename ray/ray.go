package ray

import "go-raytrace/vec"

type Ray struct {
	origin    vec.Vector3
	direction vec.Vector3
}

func (r *Ray) Origin() vec.Vector3 {
	return r.origin
}

func (r *Ray) Direction() vec.Vector3 {
	return r.direction
}

func (r *Ray) At(t float64) vec.Vector3 {
	v := vec.Scaled(r.direction, t)
	return v.Add(r.origin)
}
