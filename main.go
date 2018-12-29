package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image"
	_ "image/png"
	"os"
)

func run(){
	/* Create the window */
	cfg := pixelgl.WindowConfig{
		Title:  "Space Invader - Théo Lannoye",
		Bounds: pixel.R(0,0, 1024, 768),
		VSync:true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	/* Load images */
	vaisseau, err := loadPicture("images/vaisseau.png")
	if err != nil {
		panic(err)
	}

	/* Sprites init */
	spriteVaisseau := pixel.NewSprite(vaisseau, vaisseau.Bounds())

	win.Clear(colornames.Grey)

	lateralPos := 0.0

	/* Display vaisseau */
	spriteVaisseau.Draw(win, pixel.IM.Moved(win.Bounds().Moved(pixel.V(lateralPos,-340)).Center ()))

	/* Keyboard actions */
	for !win.Closed() {
		if win.Pressed(pixelgl.KeyLeft) {
			lateralPos--
			if lateralPos >= -450.0 {
				win.Clear(colornames.Grey)
				spriteVaisseau.Draw(win, pixel.IM.Moved(win.Bounds().Moved(pixel.V(lateralPos,-340)).Center ()))
			} else {
				lateralPos += 20
				win.Clear(colornames.Grey)
				spriteVaisseau.Draw(win, pixel.IM.Moved(win.Bounds().Moved(pixel.V(lateralPos,-340)).Center ()))
			}
		}
		if win.Pressed(pixelgl.KeyRight) {
			lateralPos++
			if lateralPos <= 450 {
				win.Clear(colornames.Grey)
				spriteVaisseau.Draw(win, pixel.IM.Moved(win.Bounds().Moved(pixel.V(lateralPos,-340)).Center ()))
			} else {
				lateralPos -= 20
				win.Clear(colornames.Grey)
				spriteVaisseau.Draw(win, pixel.IM.Moved(win.Bounds().Moved(pixel.V(lateralPos,-340)).Center ()))
			}
		}

		win.Update()
	}
}

func main(){
	pixelgl.Run(run)
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}