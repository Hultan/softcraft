package main

import (
	"fmt"
	"time"

	"softcraft/pkg/assetManager"
	"softcraft/pkg/common"
	"softcraft/pkg/components"
	"softcraft/pkg/player"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	// Initialize SDL
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("initializing SDL failed:", err)
		return
	}

	// Create the window
	window, err := sdl.CreateWindow(
		"Gaming in Go",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		common.ScreenWidth, common.ScreenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("initializing window failed:", err)
		return
	}
	defer func(window *sdl.Window) {
		_ = window.Destroy()
	}(window)

	// Create the renderer
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("initializing renderer failed:", err)
		return
	}
	defer func(renderer *sdl.Renderer) {
		_ = renderer.Destroy()
	}(renderer)

	// Load assets
	assets :=assetManager.New()
	assets.Load(renderer)

	// Create player and the world
	common.Elements = append(common.Elements, player.NewPlayer(assets))
	world, err := components.NewWorld(assets)
	if err != nil {
		panic(err)
	}

	// GAME LOOP
	for {
		frameStartTime := time.Now()

		// Handle application quitting
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		// Set background color and clear screen
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

		// Update and draw world
		err = world.OnUpdate()
		if err != nil {
			fmt.Println("updating world failed:", err)
			return
		}
		err = world.OnDraw(renderer)
		if err != nil {
			fmt.Println("drawing failed:", err)
			return
		}

		// Update and draw elements
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

		// Handle collisions
		if err = common.CheckCollisions(); err != nil {
			fmt.Println("checking collisions failed:", err)
			return
		}

		renderer.Present()

		common.Delta = time.Since(frameStartTime).Seconds() * common.TargetTicksPerSecond
	}
	// END GAME LOOP
}
