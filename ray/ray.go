package ray

import "go-raytrace/vec"

type Ray struct {
	Origin    vec.Vector3
	Direction vec.Vector3
}

func (r *Ray) At(t float64) vec.Vector3 {
	v := vec.Scaled(r.Direction, t)
	return *v.Add(r.Origin)
}
