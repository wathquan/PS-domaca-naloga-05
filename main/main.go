package main

import (
	"context"
	"os"

	"github.com/urfave/cli/v3"
	"github.com/wathquan/PS-domaca-naloga-05/vreme"
)

var t_min float64
var t_max float64
var min_st_meritev int

func simuliraj() {
	vreme.DodajPostajo("P1", "Krnica", "Kranjska gora")
	vreme.DodajPostajo("P2", "Kredarica", "Triglav")
	vreme.DodajPostajo("P3", "Zvoh", "Krvavec")
	vreme.DodajPostajo("P4", "Orlove glave", "Vogel")
	vreme.DodajPostajo("P5", "Špik", "Gozd Martuljek")

	veljavna_meritev := 0.0
	for i := 0; i < 10; i++ {
		vreme.DodajMeritev("P1", veljavna_meritev)
		veljavna_meritev += 2.3
	}

	neveljavna_meritev := -100.0
	vreme.DodajMeritev("P1", neveljavna_meritev)
	neveljavna_meritev = 100
	vreme.DodajMeritev("P1", neveljavna_meritev)
	vreme.DodajMeritev("neobstojeca_postaja", 0.0)

	vreme.IzpisMeritev()
	vreme.IzpisPorocila(min_st_meritev, t_min, t_max)
}

func main() {
	cmd := cli.Command{
		Name: "main",
		Flags: []cli.Flag{
			&cli.Float64Flag{
				Name:  "minTemp",
				Value: -50,
			},
			&cli.Float64Flag{
				Name:  "maxTemp",
				Value: 50,
			},
			&cli.IntFlag{
				Name:  "stMeritev",
				Value: 3,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			t_min = cmd.Float64("minTemp")
			t_max = cmd.Float64("maxTemp")
			min_st_meritev = cmd.Int("stMeritev")

			simuliraj()
			return nil
		},
	}

	cmd.Run(context.Background(), os.Args)
}
