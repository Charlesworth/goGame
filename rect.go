package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

type Rect struct {
	X int
	Y int
	W int
	H int
}

func (rect *Rect) Colision(otherRect Rect) bool {
	var inx, iny bool

	if ((rect.X < otherRect.X) && (otherRect.X < (rect.X + rect.W))) ||
		((rect.X < (otherRect.X + otherRect.W)) && ((otherRect.X + otherRect.W) < (rect.X + rect.W))) {
		inx = true
	}

	if ((rect.Y < otherRect.Y) && (otherRect.Y < (rect.Y + rect.H))) ||
		((rect.Y < (otherRect.Y + otherRect.H)) && ((otherRect.Y + otherRect.H) < (rect.Y + rect.H))) {
		iny = true
	}

	if inx && iny {
		log.Println("IN!!!!")
		return true
	}

	return false
}

func (rect *Rect) GetSDLRect() *sdl.Rect {
	return &sdl.Rect{
		X: int32(rect.X),
		Y: int32(rect.Y),
		W: int32(rect.W),
		H: int32(rect.H),
	}
}
