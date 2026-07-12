package main

import (
	"go-raytrace/ray"
	"go-raytrace/vec"
)

type HitRecord struct {
	p      vec.Vector3
	normal vec.Vector3
	t      float64
}

type Hittable interface {
	Hit(ray ray.Ray, tMin float64, tMax float64, record *HitRecord) bool
}
