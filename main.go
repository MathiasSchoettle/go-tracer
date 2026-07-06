package main

import (
	"fmt"
	img "go-raytrace/image"
	"go-raytrace/ray"
	"go-raytrace/vec"
	"os"
)

func main() {

	// image
	aspectRatio := 16.0 / 9.0
	imageWidth := 400
	imageHeight := int(float64(imageWidth) / aspectRatio)
	if imageHeight < 1 {
		imageHeight = 1
	}
	image := img.Create(imageWidth, imageHeight)

	fmt.Printf("w %d, h %d \n", imageWidth, imageHeight)

	// camera
	focalLength := 1.0
	viewportHeight := 2.0
	viewportWidth := viewportHeight * (float64(imageWidth) / float64(imageHeight))
	cameraCenter := vec.Of(0, 0, 0)

	viewportU := vec.Of(viewportWidth, 0, 0)
	viewportV := vec.Of(0, -viewportHeight, 0)

	pixelDeltaU := vec.Scaled(viewportU, 1.0/float64(imageWidth))
	pixelDeltaV := vec.Scaled(viewportV, 1.0/float64(imageHeight))

	viewportUpperLeft := vec.Subtracted(cameraCenter, vec.Of(0, 0, focalLength))
	viewportUpperLeft.Subtract(vec.Scaled(viewportU, 0.5)).Subtract(vec.Scaled(viewportV, 0.5))

	pixel00Loc := vec.Added(viewportUpperLeft, vec.Scaled(vec.Added(pixelDeltaU, pixelDeltaV), 0.5))

	fmt.Printf("%f, %f, %f\n", viewportU.X(), viewportU.Y(), viewportU.Z())

	for y := 0; y < imageHeight; y++ {
		fmt.Printf("\rScanlines remaining: %d", imageHeight-y)

		for x := 0; x < imageWidth; x++ {
			deltaU := vec.Scaled(pixelDeltaU, float64(x))
			deltaV := vec.Scaled(pixelDeltaV, float64(y))
			pixelCenter := vec.Added(pixel00Loc, deltaU)
			pixelCenter.Add(deltaV)

			rayDir := vec.Subtracted(pixelCenter, cameraCenter)
			cameraRay := ray.Ray{Origin: cameraCenter, Direction: rayDir}

			color := rayColor(cameraRay)

			ir := int(255.999 * color.X())
			ig := int(255.999 * color.Y())
			ib := int(255.999 * color.Z())

			image.Set(x, y, [3]int{ir, ig, ib})
		}
	}

	fmt.Println("\ndone")

	err := os.WriteFile("image.ppm", image.ToPPM(), 0777)
	if err != nil {
		return
	}
}

func hitSphere(center vec.Vector3, radius float64, ray ray.Ray) bool {
	oc := vec.Subtracted(center, ray.Origin)
	a := vec.Dot(ray.Direction, ray.Direction)
	b := -2.0 * vec.Dot(ray.Direction, oc)
	c := vec.Dot(oc, oc) - radius*radius

	discriminant := b*b - 4*a*c

	return discriminant >= 0.0
}

func rayColor(ray ray.Ray) vec.Vector3 {
	if hitSphere(vec.Of(0, 0, -1), 0.5, ray) {
		return vec.Of(1, 0, 0)
	}

	direction := vec.Normalized(ray.Direction)
	a := 0.5 * (direction.Y() + 1.0)
	first := vec.Scaled(vec.Of(1, 1, 1), 1.0-a)
	second := vec.Scaled(vec.Of(0.5, 0.7, 0.2), a)
	return vec.Added(first, second)
}
