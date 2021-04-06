package main

// TODO
// Handle window resizing
// Make AI less perfect (make human capable of scoring)
// Mouse/Joystick control
// Load images for assets

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const winWidth, winHeight int = 1000, 750

type gameState int

const (
	start gameState = iota
	play
	pause
	gameOver
)

var state = start

var nums = [][]byte{
	{
		1, 1, 1,
		1, 0, 1,
		1, 0, 1,
		1, 0, 1,
		1, 1, 1,
	},
	{
		0, 1, 0,
		0, 1, 0,
		0, 1, 0,
		0, 1, 0,
		0, 1, 0,
	},
	{
		1, 1, 1,
		0, 0, 1,
		1, 1, 1,
		1, 0, 0,
		1, 1, 1,
	},
	{
		1, 1, 1,
		0, 0, 1,
		1, 1, 1,
		0, 0, 1,
		1, 1, 1,
	},
	{
		1, 0, 1,
		1, 0, 1,
		1, 1, 1,
		0, 0, 1,
		0, 0, 1,
	},
	{
		1, 1, 1,
		1, 0, 0,
		1, 1, 1,
		0, 0, 1,
		1, 1, 1,
	},
	{
		1, 1, 1,
		1, 0, 0,
		1, 1, 1,
		1, 0, 1,
		1, 1, 1,
	},
	{
		1, 1, 1,
		0, 0, 1,
		0, 0, 1,
		0, 0, 1,
		0, 0, 1,
	},
}

func drawNumber(pos pos, color color, size int, num int, pixels []byte) {
	startX := int(pos.x) - (size*3)/2
	startY := int(pos.y) - (size*5)/2

	for i, v := range nums[num] {
		if v == 1 {
			for y := startY; y < startY+size; y++ {
				for x := startX; x < startX+size; x++ {
					setPixel(float32(x), float32(y), color, pixels)
				}
			}
		}
		startX += size
		if (i+1)%3 == 0 {
			startY += size
			startX -= size * 3
		}
	}
}

type points struct {
	player float32
	comp   float32
}

type color struct {
	r, g, b byte
}

type pos struct {
	x, y int
}

type ball struct {
	pos
	radius float32
	xVel   float32
	yVel   float32
	color  color
}

var bounce int = 0

func (ball *ball) update(paddle1 *paddle, paddle2 *paddle, points *points, elapsedTime float32) {
	// Bottom window collision
	if float32(ball.y)+ball.radius >= float32(winHeight) {
		ball.yVel = -ball.yVel
		ball.y = int(float32(winHeight) - ball.radius - 5)
	}

	// Top window collision
	if float32(ball.y)-ball.radius <= 0 {
		ball.yVel = -ball.yVel
		ball.y = int(0 + ball.radius + 5)
	}

	// Paddle 1 collision
	if float32(ball.x)-ball.radius <= float32(paddle1.x)+paddle1.w/2 && float32(ball.x)-ball.radius >= float32(paddle1.x)-paddle1.w/2 {
		if float32(ball.y) <= float32(paddle1.y)+paddle1.h/2 && float32(ball.y) >= float32(paddle1.y)-paddle1.h/2 {
			if float32(ball.y) >= float32(paddle1.y)-paddle1.h/2 && float32(ball.y) <= float32(paddle1.y)-((paddle1.h/2)/3*2) {
				ball.yVel = ball.yVel - 120
				if ball.yVel >= 0 {
					ball.yVel = -ball.yVel
				}
			} else if float32(ball.y) <= float32(paddle1.y)-((paddle1.h/2)/3*2) && float32(ball.y) <= float32(paddle1.y)-((paddle1.h/2)/3) {
				ball.yVel = ball.yVel - 90
				if ball.yVel >= 0 {
					ball.yVel = -ball.yVel
				}
			} else if float32(ball.y) >= float32(paddle1.y)+((paddle1.h/2)/3) && float32(ball.y) <= float32(paddle1.y) {
				ball.yVel = ball.yVel - 60
				if ball.yVel >= 0 {
					ball.yVel = -ball.yVel
				}
			} else if float32(ball.y) >= float32(paddle1.y) && float32(ball.y) <= float32(paddle1.y)+((paddle1.h/2)/3) {
				ball.yVel = ball.yVel + 60
				if ball.yVel <= 0 {
					ball.yVel = -ball.yVel
				}
			} else if float32(ball.y) >= float32(paddle1.y)+((paddle1.h/2)/3) && float32(ball.y) <= float32(paddle1.y)+((paddle1.h/2)/3*2) {
				ball.yVel = ball.yVel + 90
				if ball.yVel <= 0 {
					ball.yVel = -ball.yVel
				}
			} else if float32(ball.y) >= float32(paddle1.y)+((paddle1.h/2)/3*2) && float32(ball.y) <= float32(paddle1.y)+(paddle1.h/2) {
				ball.yVel = ball.yVel + 120
				if ball.yVel <= 0 {
					ball.yVel = -ball.yVel
				}
			}

			// fmt.Println("Ball y Velocity:", ball.yVel)

			ball.x = int(float32(paddle1.x) + paddle1.w/2 + ball.radius + 2)
			ball.xVel = -ball.xVel
			bounce++
			if bounce%3 == 0 {
				ball.levelUp()

			}
		}
	}
	// Paddle 2 collision
	if float32(ball.x)+ball.radius >= float32(paddle2.x)-paddle2.w/2 && float32(ball.x)+ball.radius <= float32(paddle2.x)+paddle2.w/2 {
		if float32(ball.y) <= float32(paddle2.y)+paddle2.h/2 && float32(ball.y) >= float32(paddle2.y)-paddle2.h/2 {
			if float32(ball.y) >= float32(paddle2.y)-paddle2.h/2 && float32(ball.y) <= float32(paddle2.y)-((paddle2.h/2)/3*2) {
				ball.yVel = ball.yVel - 120
				if ball.yVel >= 0 {
					ball.yVel = -ball.yVel
				}
			} else if float32(ball.y) <= float32(paddle2.y)-((paddle2.h/2)/3*2) && float32(ball.y) <= float32(paddle2.y)-((paddle2.h/2)/3) {
				ball.yVel = ball.yVel - 90
				if ball.yVel >= 0 {
					ball.yVel = -ball.yVel
				}
			} else if float32(ball.y) >= float32(paddle2.y)+((paddle2.h/2)/3) && float32(ball.y) <= float32(paddle2.y) {
				ball.yVel = ball.yVel - 60
				if ball.yVel >= 0 {
					ball.yVel = -ball.yVel
				}
			} else if float32(ball.y) >= float32(paddle2.y) && float32(ball.y) <= float32(paddle2.y)+((paddle2.h/2)/3) {
				ball.yVel = ball.yVel + 60
				if ball.yVel <= 0 {
					ball.yVel = -ball.yVel
				}
			} else if float32(ball.y) >= float32(paddle2.y)+((paddle2.h/2)/3) && float32(ball.y) <= float32(paddle2.y)+((paddle2.h/2)/3*2) {
				ball.yVel = ball.yVel + 90
				if ball.yVel <= 0 {
					ball.yVel = -ball.yVel
				}
			} else if float32(ball.y) >= float32(paddle2.y)+((paddle2.h/2)/3*2) && float32(ball.y) <= float32(paddle2.y)+(paddle2.h/2) {
				ball.yVel = ball.yVel + 120
				if ball.yVel <= 0 {
					ball.yVel = -ball.yVel
				}
			}
			ball.x = int(float32(paddle2.x) - paddle2.w/2 - ball.radius - 2)
			ball.xVel = -ball.xVel
			bounce++
			if bounce%3 == 0 {
				ball.levelUp()
				// fmt.Println("Ball Velocity:", ball.xVel)
			}
		}
	}
	// Point Scoring
	if float32(ball.x)-ball.radius < 0 {
		points.comp += 1
		state = pause
		if points.comp >= 7 {
			state = gameOver
			fmt.Println("Computer wins!!")
			os.Exit(0)
		}
		ball.x = winWidth / 2
		ball.y = winHeight / 2
		bounce = 1
		ball.resetBall(1)
	} else if float32(ball.x)+ball.radius > float32(winWidth) {
		points.player += 1
		state = pause
		if points.player >= 7 {
			state = gameOver
			fmt.Println("Player wins!!")
			os.Exit(0)
		}
		ball.x = winWidth / 2
		ball.y = winHeight / 2
		bounce = 0
		ball.resetBall(0)
	}

	xDistance := ball.xVel * elapsedTime
	yDistance := ball.yVel * elapsedTime

	ball.x = ball.x + int(xDistance)
	ball.y = ball.y + int(yDistance)
}

func (ball *ball) draw(pixels []byte) {
	// YAGNI - you afloat32 gonna need it
	for y := -ball.radius; y < ball.radius; y++ {
		for x := 0 - ball.radius; x < ball.radius; x++ {
			if x*x+y*y < ball.radius*ball.radius {
				setPixel(float32(ball.x)+x, float32(ball.y)+y, color{241, 241, 241}, pixels)
			}
		}
	}
}

//*
func (ball *ball) resetBall(player int) {
	rand.Seed(time.Now().Unix())
	ball.xVel = (20 + float32(rand.Intn(10))) * 10
	ball.yVel = (20 + float32(rand.Intn(10))) * 10
	ball.x = winWidth / 2
	ball.y = rand.Intn(winHeight)
	if player != 0 {
		ball.xVel = -ball.xVel
	}
}

//*/
//*
func (ball *ball) levelUp() {
	if ball.xVel > 0 && ball.xVel <= 1900 {
		ball.xVel = ball.xVel + 40
	} else if ball.xVel <= 0 && ball.xVel >= 1900 {
		ball.xVel = ball.xVel - 40
	}
}

//*/
type paddle struct {
	pos
	w     float32
	h     float32
	speed float32
	color color
}

func (paddle *paddle) update(keyState []uint8, elapsedTime float32) {
	if keyState[sdl.SCANCODE_UP] != 0 {
		if float32(paddle.y)-paddle.h/2 >= 0 {
			distance := paddle.speed * elapsedTime
			paddle.y = paddle.y - int(distance)
		}
	}

	if keyState[sdl.SCANCODE_DOWN] != 0 {
		if float32(paddle.y)+paddle.h/2 < float32(winHeight) {
			distance := paddle.speed * elapsedTime
			paddle.y = paddle.y + int(distance)
		}
	}
}

func (paddle *paddle) aiUpdate(ball *ball, elapsedTime float32) {
	// Unbeatable, computer will always perfectly follow the ball
	paddle.y = ball.y
}

func (paddle *paddle) draw(pixels []byte) {
	startX := float32(paddle.x) - paddle.w/2
	startY := float32(paddle.y) - paddle.h/2

	for y := 0; y < int(paddle.h); y++ {
		for x := 0; x < int(paddle.w); x++ {
			setPixel(startX+float32(x), startY+float32(y), color{241, 241, 241}, pixels)
		}
	}
}

func setPixel(x, y float32, c color, pixels []byte) {
	index := (y*float32(winWidth) + x) * 4

	if index < float32(len(pixels))-4 && index >= 0 {
		pixels[int(index)] = c.r
		pixels[int(index)+1] = c.g
		pixels[int(index)+2] = c.b
	}

}

func clearScreen(pixels []byte) {
	for i := range pixels {
		pixels[i] = 0
	}
}

func main() {
	// Added after EP06 to address macosx issues
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Pong", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(winWidth), int32(winHeight), sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer renderer.Destroy()

	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(winWidth), int32(winHeight))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tex.Destroy()

	// Randomize the ball starting position and speed
	rand.Seed(time.Now().Unix())
	xVel := (20 + float32(rand.Intn(10))) * 10
	yVel := (20 + float32(rand.Intn(10))) * 10
	ballStartY := float32(rand.Intn(int(winHeight)))

	pixels := make([]byte, winWidth*winHeight*4)

	// Define Paddles and ball
	player1 := paddle{pos{50, winHeight / 2}, 20, 100, 500, color{241, 241, 241}}
	player2 := paddle{pos{winWidth - 50, winHeight / 2}, 20, 100, 500, color{241, 241, 241}}
	ball := ball{pos{winWidth / 2, int(ballStartY)}, 10, xVel, yVel, color{241, 241, 241}}
	points := points{0, 0}

	keyState := sdl.GetKeyboardState()

	var frameStart time.Time
	var elapsedTime float32

	// Changd after EP 06 to address MacOSX
	// OSX requires that you consume events for windows to open and work properly
	// Game Loop
	for {
		frameStart = time.Now()
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				fmt.Println("quit")
				return
			}
		}

		if state == play {
			player1.update(keyState, elapsedTime)
			player2.aiUpdate(&ball, elapsedTime)
			ball.update(&player1, &player2, &points, elapsedTime)
			if keyState[sdl.SCANCODE_ESCAPE] != 0 {
				state = pause
			}
		} else if state == start {
			if keyState[sdl.SCANCODE_SPACE] != 0 {
				points.player = 0
				points.comp = 0
				state = play
			}
		} else if state == pause {
			if keyState[sdl.SCANCODE_SPACE] != 0 {
				state = play
			}
		} else if state == gameOver {
			sdl.Quit()
		}

		clearScreen(pixels)
		drawNumber(pos{int(winWidth) / 5, 50}, color{241, 241, 241}, 10, int(points.player), pixels)
		drawNumber(pos{(int(winWidth) / 5) * 4, 50}, color{241, 241, 241}, 10, int(points.comp), pixels)
		player1.draw(pixels)
		player2.draw(pixels)
		ball.draw(pixels)
		tex.Update(nil, pixels, winWidth*4)
		renderer.Copy(tex, nil, nil)
		renderer.Present()

		// sdl.Delay(10)
		elapsedTime = float32(time.Since(frameStart).Seconds())
		// fmt.Println(elapsedTime)
		if elapsedTime < .005 {
			sdl.Delay(5 - uint32(elapsedTime/1000))
			elapsedTime = float32(time.Since(frameStart).Seconds())
		}
	}
}
