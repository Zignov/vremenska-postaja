package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Zignov/vremenska-postaja/vreme"
	"github.com/urfave/cli/v3"
)

func zazeni(minTemp, maxTemp float64, stMeritev int) {
	postaje := make(map[string]vreme.Postaja)

	vreme.DodajPostajo(postaje, "LJ01", "Tivoli", "Ljubljana")
	vreme.DodajPostajo(postaje, "MB02", "Center", "Maribor")
	vreme.DodajPostajo(postaje, "KR03", "Triglav", "Kranj")
	vreme.DodajPostajo(postaje, "KP04", "Luka", "Koper")

	vreme.DodajMeritev(postaje, "LJ01", 15.5, minTemp, maxTemp)
	vreme.DodajMeritev(postaje, "LJ01", 16.2, minTemp, maxTemp)
	vreme.DodajMeritev(postaje, "LJ01", 14.8, minTemp, maxTemp)

	vreme.DodajMeritev(postaje, "MB02", 12.3, minTemp, maxTemp)
	vreme.DodajMeritev(postaje, "MB02", 13.1, minTemp, maxTemp)
	vreme.DodajMeritev(postaje, "MB02", 11.9, minTemp, maxTemp)

	vreme.DodajMeritev(postaje, "KR03", 10.5, minTemp, maxTemp)
	vreme.DodajMeritev(postaje, "KR03", 11.2, minTemp, maxTemp)
	vreme.DodajMeritev(postaje, "KR03", 9.8, minTemp, maxTemp)

	vreme.DodajMeritev(postaje, "KP04", 18.5, minTemp, maxTemp)
	vreme.DodajMeritev(postaje, "KP04", 19.2, minTemp, maxTemp)

	vreme.DodajMeritev(postaje, "LJ01", 17.3, minTemp, maxTemp)
	vreme.DodajMeritev(postaje, "MB02", -60., minTemp, maxTemp)
	vreme.DodajMeritev(postaje, "KR17", 25.0, minTemp, maxTemp)

	fmt.Println()
	vreme.IzpisMeritev(postaje)

	fmt.Println()
	vreme.IzpisPorocila(postaje, stMeritev)
}

func main() {
	cmd := &cli.Command{
		Flags: []cli.Flag{
			&cli.Float64Flag{Name: "minTemp", Value: -50},
			&cli.Float64Flag{Name: "maxTemp", Value: 50},
			&cli.IntFlag{Name: "stMeritev", Value: 3},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			zazeni(
				cmd.Float64("minTemp"),
				cmd.Float64("maxTemp"),
				cmd.Int("stMeritev"),
			)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
