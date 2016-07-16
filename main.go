package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"os"
	"time"
)

var winTitle string = "Go Game"
var winWidth, winHeight int = 800, 600

func main() {
	var window *sdl.Window
	var renderer *sdl.Renderer

	sdl.Init(sdl.INIT_EVERYTHING)

	window = createWindow(winTitle, winHeight, winWidth)
	defer window.Destroy()

	renderer = createRenderer(window)
	defer renderer.Destroy()

	secTickChan := time.Tick(time.Second)
	// frameTickChan := time.Tick(time.Microsecond * 16400)

	var events = &Events{}
	var fps = 0
	var view = NewGameView()

	for {
		// Pump events
		events.GetEvents()

		// Pass events and renderer to view
		view.Render(renderer, events)

		// If view returned a new view, use that instead
		// if newView != nil {
		// 	view = newView
		// }

		// This structure logs the fps
		select {
		case <-secTickChan:
			log.Println("fps:", fps)
			fps = 0
		default:
			fps++
		}

		// Delay the next frame rendering to free up CPU
		sdl.Delay(13)
		// <-frameTickChan
	}

}

// interface render(renderer *sdl.Renderer, events *Events) (newView interfaceOfView) {
// }

func createWindow(title string, height int, width int) *sdl.Window {
	window, err := sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		width, height, sdl.WINDOW_SHOWN)

	if err != nil {
		log.Println("Failed to create window, exiting")
		os.Exit(1)
	}

	return window
}

func createRenderer(window *sdl.Window) *sdl.Renderer {
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)

	if err != nil {
		log.Println("Failed to create renderer, exiting")
		os.Exit(1)
	}

	return renderer
}
