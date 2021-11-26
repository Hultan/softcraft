package main

import (
	"fmt"
	"time"

	"softcraft/pkg/components"
	"softcraft/pkg/types"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("initializing SDL failed:", err)
		return
	}

	window, err := sdl.CreateWindow(
		"Gaming in Go",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		types.ScreenWidth, types.ScreenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("initializing window failed:", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("initializing renderer failed:", err)
		return
	}
	defer renderer.Destroy()

	components.Elements = append(components.Elements, newPlayer(renderer))

	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			x := (float64(i)/5)*types.ScreenWidth + (basicEnemySize / 2.0)
			y := float64(j)*basicEnemySize + (basicEnemySize / 2.0)

			components.Elements = append(components.Elements, newBasicEnemy(renderer, x, y))
		}
	}

	components.InitBulletPool(renderer)

	for {
		frameStartTime := time.Now()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		err = renderer.SetDrawColor(255, 255, 255, 255)
		if err != nil {
			fmt.Println("setting draw color failed:", err)
			return
		}
		err = renderer.Clear()
		if err != nil {
			fmt.Println("clearing screen failed:", err)
			return
		}

		for _, elem := range components.Elements {
			if elem.Active {
				err = elem.Update()
				if err != nil {
					fmt.Println("updating element failed:", err)
					return
				}
				err = elem.Draw(renderer)
				if err != nil {
					fmt.Println("drawing element failed:", err)
				}
			}
		}

		if err := checkCollisions(); err != nil {
			fmt.Println("checking collisions failed:", err)
			return
		}

		renderer.Present()

		types.Delta = time.Since(frameStartTime).Seconds() * types.TargetTicksPerSecond
	}
}
