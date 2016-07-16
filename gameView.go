package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"math/rand"
	"os"
)

type gameView struct {
	player player
}

func NewGameView() gameView {
	p := newPlayer(200, 200)

	return gameView{
		player: p,
	}
}

var count int = 0
var score int = 0
var diff int = 0

var meators []meator

func (v1 *gameView) Render(renderer *sdl.Renderer, events *Events) {
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

	var destroyedMeators []int
	for i, enemy := range meators {
		enemy.move()
		enemySDLRect := enemy.rect.GetSDLRect()
		renderer.SetDrawColor(0, 255, 0, 255)
		renderer.DrawRect(enemySDLRect)
		renderer.FillRect(enemySDLRect)

		if v1.player.rect.Colision(enemy.rect) {
			//display score here
			renderer.Present()
			log.Println("You died :(")
			log.Println("Score:", score)
			sdl.Delay(1000)
			os.Exit(0)
		}
		meators[i] = enemy

		if enemy.rect.Colision(v1.player.weapon.rect) {
			destroyedMeators = append(destroyedMeators, i)
			score += 10
		}

		if enemy.isOffScreen() {
			destroyedMeators = append(destroyedMeators, i)
			score++
		}
	}

	for _, i := range destroyedMeators {
		meators = append(meators[:i], meators[i+1:]...)
	}

	renderer.Present()

	if count > 100 {
		newMeator := spawnMeator()
		meators = append(meators, newMeator)
		count = rand.Intn(30) + diff
		if diff < 60 {
			diff = diff + 3
		}
	}
	count++

	return
}

type meator struct {
	rect  Rect
	speed int
}

func (m *meator) move() {
	m.rect.X -= m.speed
}

func (m *meator) isOffScreen() bool {
	return (m.rect.X + m.rect.W) < 0
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
