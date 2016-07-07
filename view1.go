package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
	// "math"
)

type view1 struct {
	player player
}

type player struct {
	x           int32
	y           int32
	speedOnAxis float64
	speedOnDiag float64
}

var PlayerSpeed float64 = 10.0

func NewView1() view1 {
	p := player{
		x:           350,
		y:           250,
		speedOnAxis: PlayerSpeed,
		// speedOnDiag: (1 / math.Sqrt(2.0) * PlayerSpeed)
		speedOnDiag: 2.0 * PlayerSpeed,
	}

	return view1{
		player: p,
	}
}

func (v1 *view1) Render(renderer *sdl.Renderer, events *Events) {

	v1.player.calculateMovement(events)

	renderer.SetDrawColor(0, 0, 0, 0)
	renderer.Clear()

	rectPlayer := sdl.Rect{v1.player.x, v1.player.y, 100, 100}
	renderer.SetDrawColor(255, 0, 0, 255)
	renderer.DrawRect(&rectPlayer)
	renderer.FillRect(&rectPlayer)
	renderer.Present()

	return
}

func (player *player) calculateMovement(events *Events) {
	diagonalMovement := (events.up != events.down) && (events.left != events.right)
	var speed int32

	if diagonalMovement {
		speed = int32(player.speedOnDiag)
		log.Println("diag")
		log.Println(speed)
	} else {
		speed = int32(player.speedOnAxis)
		log.Println("straight")
		log.Println(speed)
	}

	if events.left {
		player.x -= speed
	} else if events.right {
		player.x += speed
	}

	if events.up {
		player.y -= speed
	} else if events.down {
		player.y += speed
	}
}
