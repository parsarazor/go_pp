package main

import (
	"image"
	"image/color"
	"image/gif"
	"log"
	"math"
	"net/http"
)

var palette = []color.Color{
	color.RGBA{15, 15, 35, 255},      // dark background
	color.RGBA{100, 176, 255, 255},   // bright blue
	color.RGBA{200, 100, 255, 255},   // purple
	color.RGBA{100, 255, 150, 255},   // mint green
}

const (
	bgIndex    uint8 = 0
	blueIndex  uint8 = 1
	purpleIdx  uint8 = 2
	greenIndex uint8 = 3
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		quantumField(w)
	})
	http.HandleFunc("/spiral", func(w http.ResponseWriter, r *http.Request) {
		mandelbrotWave(w)
	})
	http.HandleFunc("/grid", func(w http.ResponseWriter, r *http.Request) {
		harmonicResonance(w)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// quantumField creates a complex interference pattern with multiple wave sources
func quantumField(out http.ResponseWriter) {
	const (
		size    = 200
		nframes = 128
		delay   = 5
		step    = 1
	)

	anim := gif.GIF{LoopCount: nframes}

	for frame := 0; frame < nframes; frame++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		t := float64(frame) / float64(nframes)
		tpi := t * 2 * math.Pi

		// Fill with background
		for y := 0; y <= 2*size; y++ {
			for x := 0; x <= 2*size; x++ {
				img.SetColorIndex(x, y, bgIndex)
			}
		}

		// Multiple interference sources
		sources := [][2]float64{
			{float64(size), float64(size)},
			{float64(size) + 50*math.Cos(tpi), float64(size) + 50*math.Sin(tpi)},
			{float64(size) + 50*math.Cos(tpi + 2.09), float64(size) + 50*math.Sin(tpi + 2.09)},
			{float64(size) + 50*math.Cos(tpi + 4.19), float64(size) + 50*math.Sin(tpi + 4.19)},
		}

		for y := 0; y <= 2*size; y += step {
			for x := 0; x <= 2*size; x += step {
				fx, fy := float64(x), float64(y)
				energy := 0.0

				// Calculate wave interference from all sources
				for _, src := range sources {
					dx := fx - src[0]
					dy := fy - src[1]
					dist := math.Sqrt(dx*dx + dy*dy)
					wave := math.Sin(dist*0.15 - tpi) * math.Exp(-dist*0.005)
					energy += wave
				}

				// Add fractal-like detail
				nx := fx * 0.01
				ny := fy * 0.01
				fractal := math.Sin(nx*3)*math.Cos(ny*3) + math.Sin(nx*7)*math.Cos(ny*5)
				energy += fractal * 0.3

				// Color based on energy level
				colorIdx := bgIndex
				if energy > 0.5 {
					colorIdx = blueIndex
				} else if energy > 0.2 {
					colorIdx = purpleIdx
				} else if energy > -0.5 {
					colorIdx = greenIndex
				}

				img.SetColorIndex(x, y, colorIdx)
			}
		}

		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}

// mandelbrotWave creates spiral patterns with rotation and scaling
func mandelbrotWave(out http.ResponseWriter) {
	const (
		size    = 200
		nframes = 128
		delay   = 5
		step    = 1
	)

	anim := gif.GIF{LoopCount: nframes}

	for frame := 0; frame < nframes; frame++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		t := float64(frame) / float64(nframes)
		tpi := t * 2 * math.Pi

		// Fill background
		for y := 0; y <= 2*size; y++ {
			for x := 0; x <= 2*size; x++ {
				img.SetColorIndex(x, y, bgIndex)
			}
		}

		// Draw multiple rotating spirals
		for spiralIdx := 0; spiralIdx < 3; spiralIdx++ {
			spiralPhase := tpi + float64(spiralIdx)*2.09
			scale := 1.0 + 0.3*math.Sin(spiralPhase)

			for theta := 0.0; theta < 16*math.Pi; theta += 0.02 {
				r := theta * 5 * scale
				x := r * math.Cos(theta+spiralPhase)
				y := r * math.Sin(theta+spiralPhase)

				px := int(float64(size) + x)
				py := int(float64(size) + y)

				if px >= 0 && px <= 2*size && py >= 0 && py <= 2*size {
					colorMap := []uint8{blueIndex, purpleIdx, greenIndex}
					colorIdx := colorMap[spiralIdx]
					img.SetColorIndex(px, py, colorIdx)

					// Draw thickness with surrounding pixels
					for dx := -1; dx <= 1; dx++ {
						for dy := -1; dy <= 1; dy++ {
							npx := px + dx
							npy := py + dy
							if npx >= 0 && npx <= 2*size && npy >= 0 && npy <= 2*size {
								if img.ColorIndexAt(npx, npy) == bgIndex {
									img.SetColorIndex(npx, npy, colorIdx)
								}
							}
						}
					}
				}
			}
		}

		// Add pulsing rings
		for ringIdx := 0; ringIdx < 5; ringIdx++ {
			ringRadius := 30.0 + float64(ringIdx)*20 + 15*math.Sin(tpi+float64(ringIdx)*0.6)
			ringPhase := math.Sin(ringRadius*0.05 - tpi)

			for angle := 0.0; angle < 2*math.Pi; angle += 0.05 {
				x := ringRadius * math.Cos(angle)
				y := ringRadius * math.Sin(angle)
				px := int(float64(size) + x)
				py := int(float64(size) + y)

				if px >= 0 && px <= 2*size && py >= 0 && py <= 2*size {
					if ringPhase > 0 {
						img.SetColorIndex(px, py, greenIndex)
					}
				}
			}
		}

		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}

// harmonicResonance creates complex interference patterns with standing waves
func harmonicResonance(out http.ResponseWriter) {
	const (
		size    = 200
		nframes = 128
		delay   = 5
		step    = 1
	)

	anim := gif.GIF{LoopCount: nframes}

	for frame := 0; frame < nframes; frame++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		t := float64(frame) / float64(nframes)
		tpi := t * 2 * math.Pi

		// Fill background
		for y := 0; y <= 2*size; y++ {
			for x := 0; x <= 2*size; x++ {
				img.SetColorIndex(x, y, bgIndex)
			}
		}

		// Complex standing wave patterns
		for y := 0; y <= 2*size; y += step {
			for x := 0; x <= 2*size; x += step {
				// Normalize coordinates
				nx := (float64(x) - float64(size)) / float64(size)
				ny := (float64(y) - float64(size)) / float64(size)

				// Multiple harmonic waves
				wave1 := math.Sin(nx*4*math.Pi - tpi) * math.Sin(ny*4*math.Pi - tpi)
				wave2 := math.Cos(nx*3*math.Pi + tpi) * math.Cos(ny*3*math.Pi + tpi)
				wave3 := math.Sin((nx*nx+ny*ny)*10 - tpi)

				r := math.Sqrt(nx*nx + ny*ny)
				radial := math.Sin(r*10-tpi) * math.Cos(r*5+tpi*2)

				// Combine all patterns
				total := wave1*0.3 + wave2*0.3 + wave3*0.2 + radial*0.2

				// Add turbulence
				turbulence := math.Sin(nx*7-tpi)*math.Cos(ny*11+tpi) + math.Sin(nx*13+tpi)*math.Cos(ny*5-tpi)
				total += turbulence * 0.1

				// Determine color based on total amplitude
				var colorIdx uint8 = bgIndex
				if total > 0.6 {
					colorIdx = blueIndex
				} else if total > 0.3 {
					colorIdx = purpleIdx
				} else if total > 0 {
					colorIdx = greenIndex
				}

				img.SetColorIndex(x, y, colorIdx)
			}
		}

		// Draw symmetrical elements
		for i := 0; i < 8; i++ {
			angle := float64(i)*2*math.Pi/8 + tpi
			dist := 60.0 + 30*math.Sin(tpi+float64(i)*0.5)
			px := int(float64(size) + dist*math.Cos(angle))
			py := int(float64(size) + dist*math.Sin(angle))

			if px >= 0 && px <= 2*size && py >= 0 && py <= 2*size {
				img.SetColorIndex(px, py, uint8((i%3)+1))

				// Draw small circles around each point
				for r := 1; r <= 3; r++ {
					for a := 0.0; a < 2*math.Pi; a += 0.3 {
						cpx := px + int(float64(r)*math.Cos(a))
						cpy := py + int(float64(r)*math.Sin(a))
						if cpx >= 0 && cpx <= 2*size && cpy >= 0 && cpy <= 2*size {
							img.SetColorIndex(cpx, cpy, uint8((i%3)+1))
						}
					}
				}
			}
		}

		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}
