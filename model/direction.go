package model

import (
	rapid "pgregory.net/rapid"
)

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

func GenDirection(t *rapid.T) Direction {
	gen := rapid.SampledFrom([]Direction{North, East, South, West})
	return gen.Draw(t, "Direction").(Direction)
}
