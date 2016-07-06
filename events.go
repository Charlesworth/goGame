package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

type Events struct {
	up    bool
	down  bool
	left  bool
	right bool
	quit  bool
}

func (events *Events) GetEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {

		case *sdl.QuitEvent:
			os.Exit(0)

		case *sdl.KeyDownEvent:
			if t.Repeat != 0 {
				continue
			}

			switch t.Keysym.Scancode {
			case sdl.SCANCODE_UP:
				events.up = true
			case sdl.SCANCODE_DOWN:
				events.down = true
			case sdl.SCANCODE_LEFT:
				events.left = true
			case sdl.SCANCODE_RIGHT:
				events.right = true
			case sdl.SCANCODE_ESCAPE:
				os.Exit(0)
			}

		case *sdl.KeyUpEvent:
			switch t.Keysym.Scancode {
			case sdl.SCANCODE_UP:
				events.up = false
			case sdl.SCANCODE_DOWN:
				events.down = false
			case sdl.SCANCODE_LEFT:
				events.left = false
			case sdl.SCANCODE_RIGHT:
				events.right = false
			}
		}
	}
}
