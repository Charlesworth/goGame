package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"os"
	"time"
)

var winTitle string = "Go-SDL2 Render"
var winWidth, winHeight int = 800, 600

func main() {
	var window *sdl.Window
	var renderer *sdl.Renderer

	sdl.Init(sdl.INIT_EVERYTHING)

	window = createWindow(winTitle, winHeight, winWidth)
	defer window.Destroy()

	//window.Maximize()

	renderer = createRenderer(window)
	defer renderer.Destroy()

	secTickChan := time.Tick(time.Second)
	stopChan := time.After(time.Second * 10)

	var i int32 = 0

	renderer.Clear()
	for {
		renderer.SetDrawColor(0, 0, 0, 0)
		renderer.Clear()
		// renderer.SetDrawColor(255, 0, 0, 0)
		rect := sdl.Rect{300, i, 200, 200}
		renderer.SetDrawColor(255, 0, 0, 255)
		renderer.DrawRect(&rect)
		renderer.Present()

		GetEvents()

		select {
		case <-stopChan:
			os.Exit(0)
		case <-secTickChan:
			log.Println(i)
			i = 0
		default:
			i++
		}
		sdl.Delay(13)
	}

}

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
