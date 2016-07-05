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
	var points []sdl.Point
	var rect sdl.Rect
	var rects []sdl.Rect

	window = createWindow(winTitle, winHeight, winWidth)
	defer window.Destroy()

	//window.Maximize()

	renderer = createRenderer(window)
	defer renderer.Destroy()

	renderer.Clear()

	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.DrawPoint(150, 300)

	renderer.SetDrawColor(0, 0, 255, 255)
	renderer.DrawLine(0, 0, 200, 200)

	points = []sdl.Point{{0, 0}, {100, 300}, {100, 300}, {200, 0}}
	renderer.SetDrawColor(255, 255, 0, 255)
	renderer.DrawLines(points)

	rect = sdl.Rect{300, 0, 200, 200}
	renderer.SetDrawColor(255, 0, 0, 255)
	renderer.DrawRect(&rect)

	rects = []sdl.Rect{{400, 400, 100, 100}, {550, 350, 200, 200}}
	renderer.SetDrawColor(0, 255, 255, 255)
	renderer.DrawRects(rects)

	rect = sdl.Rect{250, 250, 200, 200}
	renderer.SetDrawColor(0, 255, 0, 255)
	renderer.FillRect(&rect)

	rects = []sdl.Rect{{500, 300, 100, 100}, {200, 300, 200, 200}}
	renderer.SetDrawColor(255, 0, 255, 255)
	renderer.FillRects(rects)

	renderer.Present()

	sdl.Delay(2000)

	tickChan := time.Tick(time.Second)
	stopChan := time.After(time.Second * 10)

	var i int32 = 0

	for {
		renderer.SetDrawColor(0, 0, 0, 0)
		renderer.Clear()
		// renderer.SetDrawColor(255, 0, 0, 0)
		rect := sdl.Rect{300, i, 200, 200}
		renderer.SetDrawColor(255, 0, 0, 255)
		renderer.DrawRect(&rect)
		renderer.Present()
		select {
		case <-stopChan:
			os.Exit(0)
		case <-tickChan:
			log.Println(i)
			i = 0
		default:
			i++
			// log.Println(i)
		}
		//sdl.Delay(100)
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
