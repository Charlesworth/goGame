package main

import (
	"math/rand"
)

type meteor struct {
	rect  Rect
	speed int
}

func (m *meteor) move() {
	m.rect.X -= m.speed
}

func (m *meteor) isOffScreen() bool {
	return (m.rect.X + m.rect.W) < 0
}

func spawnMeteor() meteor {
	yPos := rand.Intn(winHeight)
	height := rand.Intn(60) + 20
	width := rand.Intn(60) + 20
	speed := rand.Intn(6)
	return meteor{
		rect:  Rect{X: winWidth, Y: yPos, H: height, W: width},
		speed: speed,
	}
}
