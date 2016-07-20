package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"math/rand"
	"os"
)

// the gameView object is what we pass the renderer and events to
// each frame and it outputs the current game state to the player
type gameView struct {
	player      player
	score       int
	meteors     []meteor
	meteorTimer int
	difficulty  int
}

func NewGameView() gameView {
	p := newPlayer(200, 200)

	return gameView{
		player: p,
	}
}

func (view *gameView) Render(renderer *sdl.Renderer, events *Events) {
	view.player.calculateMovement(events)

	// start each new frame by filling the frame black
	renderer.SetDrawColor(0, 0, 0, 0)
	renderer.Clear()

	// draw the player
	renderer.SetDrawColor(255, 0, 0, 255)
	playerSDLRect := view.player.rect.GetSDLRect()
	renderer.DrawRect(playerSDLRect)
	renderer.FillRect(playerSDLRect)

	// draw the players spinning weapon
	playerWeaponSDLRect := view.player.weapon.rect.GetSDLRect()
	renderer.DrawRect(playerWeaponSDLRect)
	renderer.FillRect(playerWeaponSDLRect)

	// destroyedMeteors is used to remove destroyed meteors outside
	// of looping through view.meteors to avoid out of range errors
	var destroyedMeteors []int

	// loop through the frames active meteors
	for i, meteor := range view.meteors {
		// move and draw and then reinsert the moved meteor
		meteor.move()
		enemySDLRect := meteor.rect.GetSDLRect()
		renderer.SetDrawColor(0, 255, 0, 255)
		renderer.DrawRect(enemySDLRect)
		renderer.FillRect(enemySDLRect)
		view.meteors[i] = meteor

		// if the player and meteor collide, end the game
		if view.player.rect.Colision(meteor.rect) {
			renderer.Present()
			log.Println("You died!")
			log.Println("Score:", view.score)
			sdl.Delay(1000)
			os.Exit(0)
		}

		// if the weapon and meteor collide, destroy it and update score
		if meteor.rect.Colision(view.player.weapon.rect) {
			destroyedMeteors = append(destroyedMeteors, i)
			view.score += 10
		}

		// if the meteor is off screen, destroy it and update score
		if meteor.isOffScreen() {
			destroyedMeteors = append(destroyedMeteors, i)
			view.score++
		}
	}

	// for each destroyed meteor, remove it from the meteors slice
	for _, i := range destroyedMeteors {
		view.meteors = append(view.meteors[:i], view.meteors[i+1:]...)
	}

	// present the updated frame to the player
	renderer.Present()

	// add new meteors depending on the difficulty and ticks since
	// last meteor was added
	if view.meteorTimer > 100 {
		newMeteor := spawnMeteor()
		view.meteors = append(view.meteors, newMeteor)
		view.meteorTimer = rand.Intn(30) + view.difficulty
		if view.difficulty < 60 {
			view.difficulty += 3
		}
	}
	view.meteorTimer++

	return
}
