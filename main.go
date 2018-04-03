package main

import (
	"github.com/nfnt/resize"
	"github.com/urfave/cli"
	"image"
	"log"
	"os"
	"strings"

	"image/draw"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
)

func setFlags(app *cli.App) {
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "o",
			Value: "output.png",
			Usage: "Specify directory for output file",
		},
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "Image Synthesis"
	app.Usage = "Synthesize images"

	setFlags(app)

	app.Action = func(c *cli.Context) error {
		outfile := c.String("o")
		if !strings.HasSuffix(outfile, ".png") {
			outfile = outfile + ".png"
		}
		basePath := c.Args().Get(0)
		overPath := c.Args().Get(1)
		baseFile, err := os.Open(basePath)
		if err != nil {
			log.Fatal(err)
		}
		overFile, err := os.Open(overPath)
		if err != nil {
			log.Fatal(err)
		}

		baseImg, _, err := image.Decode(baseFile)
		if err != nil {
			log.Fatal(err)
		}
		baseFile.Close()
		overImg, _, err := image.Decode(overFile)
		if err != nil {
			log.Fatal(err)
		}
		overFile.Close()

		baseSize := baseImg.Bounds().Size()
		overSize := overImg.Bounds().Size()
		if baseSize.X != overSize.X {
			overImg = resize.Resize(uint(baseSize.X), uint(baseSize.Y), overImg, resize.Lanczos3)
			overSize = overImg.Bounds().Size()
		}
		startPoint := image.Point{0, 0}
		overRectangle := image.Rectangle{startPoint, startPoint.Add(overSize)}
		baseRectangle := image.Rectangle{image.Point{0, 0}, baseSize}
		rgba := image.NewRGBA(baseRectangle)
		draw.Draw(rgba, baseRectangle, baseImg, image.Point{0, 0}, draw.Src)
		draw.Draw(rgba, overRectangle, overImg, image.Point{0, 0}, draw.Over)

		out, err := os.Create(outfile)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		var opt jpeg.Options
		opt.Quality = 100
		jpeg.Encode(out, rgba, &opt)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
