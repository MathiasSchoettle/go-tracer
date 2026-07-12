package main

import (
	"go-raytrace/ray"
	"go-raytrace/vec"
	"math"
)

type Sphere struct {
	center vec.Vector3
	radius float64
}

func Create(center vec.Vector3, radius float64) Sphere {
	return Sphere{center, radius}
}

func (s Sphere) hit(ray ray.Ray, tMin float64, tMax float64, record *HitRecord) bool {
	oc := vec.Subtracted(s.center, ray.Origin)
	a := ray.Direction.LengthSquared()
	h := vec.Dot(ray.Direction, oc)
	c := oc.LengthSquared() - s.radius*s.radius

	discriminant := h*h - a*c
	if discriminant < 0 {
		return false
	}

	sqrtDeterminant := math.Sqrt(discriminant)

	root := (h - sqrtDeterminant) / a
	if root <= tMin || tMax <= root {
		root = (h + sqrtDeterminant) / a
		if root <= tMin || tMax <= root {
			return false
		}
	}

	record.t = root
	record.p = ray.At(root)
	normal := vec.Subtracted(record.p, s.center)
	normal.Divide(s.radius)
	record.normal = normal

	return true
}
