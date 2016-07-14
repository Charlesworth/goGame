package main

import (
	"github.com/veandco/go-sdl2/sdl"
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

func (v1 *view1) Render(renderer *sdl.Renderer, events *Events) {

	v1.player.calculateMovement(events)

	renderer.SetDrawColor(0, 0, 0, 0)
	renderer.Clear()

	renderer.SetDrawColor(255, 0, 0, 255)
	playerSDLRect := v1.player.rect.GetSDLRect()
	renderer.DrawRect(playerSDLRect)
	renderer.FillRect(playerSDLRect)

	playerWeaponSDLRect := v1.player.weapon.rect.GetSDLRect()
	renderer.DrawRect(playerWeaponSDLRect)
	renderer.FillRect(playerWeaponSDLRect)

	enemy := Rect{300, 300, 100, 100}
	enemySDLRect := enemy.GetSDLRect()
	renderer.SetDrawColor(0, 255, 0, 255)
	renderer.DrawRect(enemySDLRect)
	renderer.FillRect(enemySDLRect)

	if v1.player.rect.Colision(enemy) {
		os.Exit(0)
	}

	renderer.Present()

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
		w.yAdd++
		w.rect.Y++
		return
	}

	if (w.yAdd == 200) && (w.xAdd > 0) {
		w.xAdd--
		w.rect.X--
		return
	}

	if (w.xAdd == 0) && (w.yAdd > 0) {
		w.yAdd--
		w.rect.Y--
		return
	}

	if (w.yAdd == 0) && (w.xAdd < 200) {
		w.xAdd++
		w.rect.X++
		return
	}
}
