package main

import (
	"fmt"
	"os"

	"github.com/H-Kiku482/cli_dice/internal/config"
	"github.com/H-Kiku482/cli_dice/internal/dice"
)

func main() {
	conf := config.GetConfig(os.Args)

	var d []dice.Dice = []dice.Dice{}
	for i := 0; i < int(conf.DiceQuantity); i++ {
		d = append(d, *dice.InitDice(conf.Face))
	}

	// roll len(d) times
	for i := 0; i < len(d); i++ {
		d[i].Roll()
	}

	// print len(d) times
	for i := 0; i < len(d); i++ {
		fmt.Print(d[i])
	}
	fmt.Print("\n")
}
