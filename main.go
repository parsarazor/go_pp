 package main

import (
	//"bufio"
	//"fmt"
	//"io/ioutil"
	"net/http"
	//	"os"
	//"strings"
	"image"
	"image/color"
	"image/gif"
	//"io"
	"math"
	"math/rand"
	"log"
)

var palette = []color.Color{color.RGBA{0,0,0,255}, color.RGBA{100, 176, 255, 255}}

const (
	whiteIndex  = 0
	unkwonIndex = 1
)
func main() {

	//var parsa = []string{"helia", "fatemeh", "sara"}
	//fmt.Println(strings.Join(parsa, " "))
	//s, sep := "", ""
	//for _, arg := range os.Args[1:] {
	//s += sep + arg
	//sep = " "
	//}
	//fmt.Println(s)

	//create a structure called map to store key from stdin
	//counts := make(map[string]int)
	// create an object that has three mothods in it Scan, Text
	//input := bufio.NewScanner(os.Stdin)

	// Scan() function returns bolean
	//for input.Scan() {

	//	counts[input.Text()]++

	//}

	//for line, n := range counts {
	//	if n > 1 {
	///	fmt.Printf("%v\t %v\n", n, line)
	//}
	//	}
	// 	counts := make(map[string]int)
	// 	files := os.Args[1:]

	// 	if len(files) == 0 {
	// 		countLines(os.Stdin, counts)
	// 	} else {
	// 		for _, arg := range files {
	// 			f, err := os.Open(arg)
	// 			if err != nil {
	// 				fmt.Fprintf(os.Stderr, "dup: %v", err)
	// 				continue
	// 			}
	// 			countLines(f, counts)
	// 			f.Close()
	// 		}
	// 	}

	// 	for line, n := range counts {
	// 		if n > 1 {
	// 			fmt.Printf("%v\t%v\n", n, line)
	// 		}

	// 	}
	// counts := make(map[string]int)
	// for _, filename := range os.Args[1:]{
	// 	data, err := ioutil.ReadFile(filename)
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "dup3: %v", err)
	// 		continue
	// 	}
	// 	for _,line := range strings.Split(string(data), "\n") {
	// 		counts[line]++
	// 	}
	// 	for line, n := range counts {
	// 		if n > 1 {
	// 			fmt.Printf("%v\t%v\n", n, line)
	// 		}
	// 	}
	// }
	//counts := make(map[string]int)
	// for _, filename := range os.Args[1:]{
	// 	// ReadFile function returns byte data type that needs
	// 	// conversion to string and a

	// 	data, err := ioutil.ReadFile(filename)
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
	// 		continue
	// 	}
	// 	for _,line := range strings.Split(string(data),"\n"){
	// 		counts[line]++
	// 	}
	// 	for line, n := range counts {
	// 		if n > 1 {
	// 			fmt.Printf("%v\t%v\n", n, line)
	// 		}
	// 	}
	// }
	// 1.4 Animated gifs
	handler := func(w http.ResponseWriter, r *http.Request){
		lissajous(w)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func lissajous(out http.ResponseWriter) {
	const (
		cycle   = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycle*2*math.Pi; t+=res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), unkwonIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

// // func countLines(f *os.File, counts map[string]int) {
// 	input := bufio.NewScanner(f)
// 	for input.Scan() {
// 		counts[input.Text()]++
// 	}
//}
