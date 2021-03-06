package main

import (
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/urfave/cli"
	"golang-challenge.org/challenge3/mosaic"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	app := cli.NewApp()
	app.Name = "challenge3"
	app.Usage = "Get Mosaic from Image and Tiles"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "main, m",
			Value: "",
			Usage: "Obligatory. path to the main image of the mosaic",
		},
		cli.StringFlag{
			Name:  "tiles, t",
			Value: "",
			Usage: "Obligatory. path to the tiles images directory",
		},
		cli.IntFlag{
			Name:  "ch",
			Usage: "number of horizontal splits",
			Value: 80,
		},
		cli.IntFlag{
			Name:  "cv",
			Usage: "number of vertical splits",
			Value: 50,
		},
		cli.IntFlag{
			Name:  "mask",
			Usage: "transparency of the tilers (Max: 255 | Min: 0)",
			Value: 180,
		},
		cli.IntFlag{
			Name:  "threads",
			Usage: "number of threads. Set parallelism",
			Value: 4,
		},
		cli.StringFlag{
			Name:  "output, o",
			Value: "",
			Usage: "Obligatory. output path of the mosaic",
		},
	}
	app.Action = func(c *cli.Context) error {
		if c.GlobalString("main") == "" || c.GlobalString("tiles") == "" || c.GlobalString("output") == "" {
			log.Fatal("main & tiles & output paths are mandatory")
		}

		if c.Int("mask") < 0 || c.Int("mask") > 255 {
			log.Fatal("mask should be between 0 and 255")
		}

		// TODO: make standard path format. end with /?

		fmt.Printf("Starting...\n")
		mosaic, err := mosaic.NewMosaic(
			c.String("main"),
			c.String("tiles"),
			c.Int("ch"),
			c.Int("cv"),
			c.Int("threads"),
			uint8(c.Int("mask")),
		)
		if err != nil {
			log.Fatal(err)
		}

		mosaic.Generate()
		mosaicpng, _ := os.Create(c.String("output"))
		png.Encode(mosaicpng, mosaic.Get())

		return nil
	}

	app.Run(os.Args)
}
