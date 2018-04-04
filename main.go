package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
	"strings"

	"./setup"
	"./synthesis"

	"image/jpeg"
)

func main() {

	app := setup.CreateCliApp()

	app.Action = func(c *cli.Context) error {
		outfile := c.String("o")
		if !strings.HasSuffix(outfile, ".png") {
			outfile = outfile + ".png"
		}
		basePath := c.Args().Get(0)
		overPath := c.Args().Get(1)

		baseImgInfo := synthesis.GetImageInfo(basePath)
		overImgInfo := synthesis.GetImageInfo(overPath)

		if baseImgInfo.Size.X != overImgInfo.Size.X {
			synthesis.ResizeImage(&overImgInfo, uint(baseImgInfo.Size.X), uint(baseImgInfo.Size.Y))
		}

		rgba := synthesis.SynthesizeImages(baseImgInfo, overImgInfo)

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
