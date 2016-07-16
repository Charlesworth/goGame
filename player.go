package main

type player struct {
	rect        Rect
	weapon      spinWeapon
	speedOnAxis int
	speedOnDiag int
}

func newPlayer(x int, y int) player {

	return player{
		rect: Rect{
			X: x,
			Y: y,
			W: 100,
			H: 100,
		},
		weapon: spinWeapon{
			rect: Rect{
				X: x - 50,
				Y: y - 50,
				W: 10,
				H: 10,
			},
		},
		speedOnAxis: 3,
		speedOnDiag: 2,
	}

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
