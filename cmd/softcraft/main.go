package main

import (
	"fmt"
	"time"

	"softcraft/pkg/common"
	"softcraft/pkg/components"
	"softcraft/pkg/player"

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
		common.ScreenWidth, common.ScreenHeight,
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

	common.Elements = append(common.Elements, player.NewPlayer(renderer))
	world := components.NewWorld(renderer)

	components.InitBulletPool(renderer)

	// GAME LOOP
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

		// Draw world
		world.OnUpdate()
		world.OnDraw(renderer)

		// Draw elements
		for _, elem := range common.Elements {
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

		if err := common.CheckCollisions(); err != nil {
			fmt.Println("checking collisions failed:", err)
			return
		}

		renderer.Present()

		common.Delta = time.Since(frameStartTime).Seconds() * common.TargetTicksPerSecond
	}
	// END GAME LOOP
}
