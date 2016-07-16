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

	xMove, yMove := 0, 0

	if events.left {
		xMove -= speed
	} else if events.right {
		xMove += speed
	}

	if events.up {
		yMove -= speed
	} else if events.down {
		yMove += speed
	}

	if (player.rect.X+xMove) > 0 &&
		(player.rect.X+player.rect.W+xMove) < winWidth {
		player.rect.X += xMove
		player.weapon.rect.X += xMove
	}

	if (player.rect.Y+yMove) > 0 &&
		(player.rect.Y+player.rect.H+yMove) < winHeight {
		player.rect.Y += yMove
		player.weapon.rect.Y += yMove
	}

	player.weapon.calculateMovement()
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
