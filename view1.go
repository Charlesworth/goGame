package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"math/rand"
	"os"
)

type view1 struct {
	player player
}

type player struct {
	rect        Rect
	weapon      spinWeapon
	speedOnAxis int
	speedOnDiag int
}

func NewView1() view1 {
	playerRect := Rect{
		X: 200,
		Y: 200,
		W: 100,
		H: 100,
	}

	weaponRect := Rect{
		X: 150,
		Y: 150,
		W: 10,
		H: 10,
	}

	p := player{
		rect:        playerRect,
		weapon:      spinWeapon{weaponRect, 0, 0},
		speedOnAxis: 3,
		speedOnDiag: 2,
	}

	return view1{
		player: p,
	}
}

type meator struct {
	rect  Rect
	speed int
}

func (m *meator) move() {
	m.rect.X -= m.speed
}

func spawnMeator() meator {
	yPos := rand.Intn(winHeight)
	height := rand.Intn(60) + 20
	width := rand.Intn(60) + 20
	speed := rand.Intn(6)
	return meator{
		rect:  Rect{X: winWidth, Y: yPos, H: height, W: width},
		speed: speed,
	}
}

// var enemyTick chan time.Time = time.Tick(time.Second * 4)
var count int = 0
var score int = 0

// var enemys []Rect = []Rect{Rect{350, 350, 100, 100},
// 	Rect{450, 450, 50, 50},
// 	Rect{0, 375, 30, 30}}
var meators []meator

func (v1 *view1) Render(renderer *sdl.Renderer, events *Events) {
	v1.player.calculateMovement(events)
	count++

	renderer.SetDrawColor(0, 0, 0, 0)
	renderer.Clear()

	renderer.SetDrawColor(255, 0, 0, 255)
	playerSDLRect := v1.player.rect.GetSDLRect()
	renderer.DrawRect(playerSDLRect)
	renderer.FillRect(playerSDLRect)

	playerWeaponSDLRect := v1.player.weapon.rect.GetSDLRect()
	renderer.DrawRect(playerWeaponSDLRect)
	renderer.FillRect(playerWeaponSDLRect)

	var destroyedMeators []int
	for i, enemy := range meators {
		enemy.move()
		enemySDLRect := enemy.rect.GetSDLRect()
		renderer.SetDrawColor(0, 255, 0, 255)
		renderer.DrawRect(enemySDLRect)
		renderer.FillRect(enemySDLRect)

		if v1.player.rect.Colision(enemy.rect) {
			//display score here
			log.Println("You died :(")
			log.Println("Score:", score)
			sdl.Delay(1000)
			os.Exit(0)
		}
		meators[i] = enemy

		// if v1.player.weapon.rect.Colision(enemy) {
		if enemy.rect.Colision(v1.player.weapon.rect) {
			log.Println("BOOM!")
			destroyedMeators = append(destroyedMeators, i)
			score++
			// meators = append(meators[:i], meators[i+1:]...)
		}
	}

	for _, i := range destroyedMeators {
		meators = append(meators[:i], meators[i+1:]...)
	}

	renderer.Present()

	if count > 180 {
		newMeator := spawnMeator()
		meators = append(meators, newMeator)
		count = 0
	}

	return
}

func (player *player) calculateMovement(events *Events) {
	diagonalMovement := (events.up != events.down) && (events.left != events.right)
	var speed int

	if diagonalMovement {
		speed = player.speedOnDiag
	} else {
		speed = player.speedOnAxis
	}

	if events.left {
		player.rect.X -= speed
		player.weapon.rect.X -= speed
	} else if events.right {
		player.rect.X += speed
		player.weapon.rect.X += speed
	}

	if events.up {
		player.rect.Y -= speed
		player.weapon.rect.Y -= speed
	} else if events.down {
		player.rect.Y += speed
		player.weapon.rect.Y += speed
	}

	player.weapon.calculateMovement()
	player.keepInWindow()
}

func (player *player) keepInWindow() {
	if (player.rect.Y + player.rect.H) > winHeight {
		player.rect.Y = winHeight - player.rect.H
	}

	if player.rect.Y < 0 {
		player.rect.Y = 0
	}

	if (player.rect.X + player.rect.W) > winWidth {
		player.rect.X = winWidth - player.rect.W
	}

	if player.rect.X < 0 {
		player.rect.X = 0
	}
}

type spinWeapon struct {
	rect Rect
	xAdd int
	yAdd int
}

func (w *spinWeapon) calculateMovement() {
	if (w.xAdd == 200) && (w.yAdd < 200) {
		w.yAdd += 4
		w.rect.Y += 4
		return
	}

	if (w.yAdd == 200) && (w.xAdd > 0) {
		w.xAdd -= 8
		w.rect.X -= 8
		return
	}

	if (w.xAdd == 0) && (w.yAdd > 0) {
		w.yAdd -= 8
		w.rect.Y -= 8
		return
	}

	if (w.yAdd == 0) && (w.xAdd < 200) {
		w.xAdd += 8
		w.rect.X += 8
		return
	}
}
