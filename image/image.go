package img

import "fmt"

type Image struct {
	data   []int
	width  int
	height int
}

func Create(w int, h int) Image {
	return Image{make([]int, w*h*3), w, h}
}

func (i *Image) At(x int, y int) [3]int {
	index := (x*i.width + y) * 3
	return [3]int(i.data[index : index+3])
}

func (i *Image) Set(x int, y int, v [3]int) {
	index := (x*i.width + y) * 3
	i.data[index+0] = v[0]
	i.data[index+1] = v[1]
	i.data[index+2] = v[2]
}

func (i *Image) ToPPM() []byte {
	image := fmt.Sprintf("P3 \n %d %d \n 255 \n", i.width, i.height)

	for y := 0; y < i.height; y++ {
		for x := 0; x < i.width; x++ {
			index := (x*i.width + y) * 3
			pixel := i.data[index : index+3]
			image += fmt.Sprintf("%d %d %d\n", pixel[0], pixel[1], pixel[2])
		}
	}

	return []byte(image)
}
