package setup

import (
	"github.com/urfave/cli"
)

func setFlags(app *cli.App) {
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "o",
			Value: "output.png",
			Usage: "Specify output file name",
		},
	}
}

func CreateCliApp() *cli.App {
	app := cli.NewApp()
	app.Name = "Image Synthesis"
	app.Usage = "Synthesize 2 images"

	setFlags(app)

	return app
}
