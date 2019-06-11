package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image"
	_ "image/png"
	"time"
)

var (
	scale  = 10
	frames = 0
	second = time.Tick(time.Second)
)

func main() {
	pixelgl.Run(run)
}

func run() {

	mW, mH := pixelgl.PrimaryMonitor().Size()
	//mW := 400.0
	//mH := 400.0

	fmt.Println(mW, mH)

	cfg := pixelgl.WindowConfig{
		Title:  "Game of Life",
		Bounds: pixel.R(0, 0, mW, mH),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.SetSmooth(true)

	win.Clear(colornames.Skyblue)

	l := NewLife(int(mW/float64(scale)), int(mH/float64(scale)))

	w := scale
	h := scale

	pic := image.NewRGBA(image.Rectangle{Max: image.Point{X: w, Y: h}})
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			pic.Set(x, y, colornames.Black)
		}
	}

	p := pixel.PictureDataFromImage(pic)

	batch := pixel.NewBatch(&pixel.TrianglesData{}, p)

	sprite := pixel.NewSprite(p, p.Bounds())

	for !win.Closed() {
		l.Step()

		win.Clear(colornames.White)
		batch.Clear()

		for y := 0; y < l.h; y++ {
			for x := 0; x < l.w; x++ {
				if l.a.Alive(x, y) {
					sprite.Draw(batch, pixel.IM.Moved(pixel.V(float64(x*scale), float64(y*scale))))
				}
			}
		}

		batch.Draw(win)

		frames++
		select {
		case <-second:
			win.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, frames))
			frames = 0
		default:
		}

		win.Update()
	}
}
