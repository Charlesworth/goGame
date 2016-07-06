package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"os"
)

type Events struct {
	up    bool
	down  bool
	left  bool
	right bool
	quit  bool
}

func GetEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		// if event != nil {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			os.Exit(0)
		case *sdl.KeyDownEvent:
			if t.Keysym.Scancode == sdl.SCANCODE_UP {
				log.Println("UP! down")
			}
			// log.Printf("[%d ms] Keyboard\ttype:%d\tUnicode:%d\tmodifiers:%d\tstate:%d\trepeat:%d\n",
			// 	t.Timestamp, t.Type, t.Keysym.Unicode, t.Keysym.Mod, t.State, t.Repeat)
		case *sdl.KeyUpEvent:
			//log.Printf("[%d ms] Keyboard\ttype:%d\tUnicode:%d\tmodifiers:%d\tstate:%d\trepeat:%d\n",
			//	t.Timestamp, t.Type, t.Keysym.Unicode, t.Keysym.Mod, t.State, t.Repeat)
			if t.Keysym.Scancode == sdl.SCANCODE_UP {
				log.Println("UP! up")
			}
		}
	}
}
