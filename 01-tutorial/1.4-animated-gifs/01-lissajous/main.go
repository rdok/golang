package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var green = color.RGBA{R: 0, G: 177, B: 106, A: 1}
var black = color.RGBA{R: 0, G: 0, B: 0, A: 1}

//var palette = []color.RGBA{green, black}
var palette = []color.Color{black, green}

const (
	blackIndex = 1 // next color in palette
)

// go run % color-x color-y index
func main() {
	colorX = os.Args[2]
	colorY = os.Args[3]
	colorIndex = os.Args[4]

	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		oscillatorRevolutions = 5
		angularResolution     = 0.001
		imgCanvasSize         = 100
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
			img.SetColorIndex(
				imgCanvasSize+int(x*imgCanvasSize+0.5),
				imgCanvasSize+int(y*imgCanvasSize+0.5),
				blackIndex,
			)
		}
		phaseDifferences += 0.1
		animation.Delay = append(animation.Delay, frameDelaysIn10sMS)
		animation.Image = append(animation.Image, img)
	}
	_ = gif.EncodeAll(out, &animation) // Ignore encoding errors
}
