package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{
	color.RGBA{R: 0, G: 0, B: 128, A: 1},
	color.RGBA{R: 0, G: 0, B: 255, A: 1},
	color.RGBA{R: 0, G: 128, B: 0, A: 1},
	color.RGBA{R: 0, G: 128, B: 128, A: 1},
	color.RGBA{R: 128, G: 0, B: 0, A: 1},
	color.RGBA{R: 255, G: 0, B: 0, A: 1},
	color.RGBA{R: 255, G: 255, B: 0, A: 1},
}

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		oscillatorRevolutions = 10
		angularResolution     = 0.001
		imgCanvasSize         = 300
		totalAnimationFrames  = 64
		frameDelaysIn10sMS    = 8
	)

	relativeFreqYOscillator := rand.Float64() * 3.0
	animation := gif.GIF{LoopCount: totalAnimationFrames}
	phaseDifferences := 0.0


	for frameIndex := 0; frameIndex < totalAnimationFrames; frameIndex++ {
		rect := image.Rect(0, 0, 2*imgCanvasSize+1, 2*imgCanvasSize+1)
		img := image.NewPaletted(rect, palette)
		for phaseIndex := 0.0; phaseIndex < oscillatorRevolutions*2*math.Pi; phaseIndex += angularResolution {
			x := math.Sin(phaseIndex)
			y := math.Sin(phaseIndex * relativeFreqYOscillator * phaseDifferences)
			seeder := rand.NewSource(time.Now().UnixNano())
			random := rand.New(seeder)

			img.SetColorIndex(
				imgCanvasSize+int(x*imgCanvasSize+0.5),
				imgCanvasSize+int(y*imgCanvasSize+0.5),
				uint8(random.Intn(1000)) % 5,
			)
		}
		phaseDifferences += 0.1
		animation.Delay = append(animation.Delay, frameDelaysIn10sMS)
		animation.Image = append(animation.Image, img)
	}
	_ = gif.EncodeAll(out, &animation) // Ignore encoding errors
}
