package coordinate_supplier

type coordinate struct {
	x int
	y int
}

func makeAscCoordinates(width, height int) []coordinate {
	coordinates := make([]coordinate, 0, width*height)
	var atX, atY int
	for {
		coordinates = append(coordinates, coordinate{
			x: atX,
			y: atY,
		})

		atX++
		if atX >= width {
			atX = 0
			atY++
		}
		if atY >= height {
			break
		}
	}
	return coordinates
}

func reverseCoordinates(cs []coordinate) {
	i := 0
	j := len(cs) - 1
	for i < j {
		cs[i], cs[j] = cs[j], cs[i]
		i++
		j--
	}
}
