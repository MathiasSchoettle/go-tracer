package main

import (
	"fmt"
	img "go-raytrace/image"
	"os"
)

func main() {

	width := 255
	height := 255

	image := img.Create(width, height)

	for y := 0; y < height; y++ {
		fmt.Printf("\rScanlines remaining: %d", height-y)

		for x := 0; x < width; x++ {

			r := float64(y) / float64(width-1)
			g := float64(x) / float64(height-1)
			b := 0.0

			ir := int(255.999 * r)
			ig := int(255.999 * g)
			ib := int(255.999 * b)

			image.Set(x, y, [3]int{ir, ig, ib})
		}
	}

	err := os.WriteFile("image.ppm", image.ToPPM(), 0777)
	if err != nil {
		return
	}
}
