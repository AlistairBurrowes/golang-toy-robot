package model

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (c Direction) String() string {
	var result string

	switch c {
	case North:
		result = "NORTH"
	case East:
		result = "EAST"
	case South:
		result = "SOUTH"
	case West:
		result = "WEST"
	}

	return result
}
