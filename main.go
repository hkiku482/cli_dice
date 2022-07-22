package main

import (
	"fmt"
	"os"
	"time"

	"github.com/H-Kiku482/cli_dice/internal/config"
	"github.com/H-Kiku482/cli_dice/internal/dice"
)

func main() {
	conf := config.GetConfig(os.Args)

	var d []dice.Dice = []dice.Dice{}
	for i := 0; i < int(conf.DiceQuantity); i++ {
		d = append(d, *dice.InitDice(conf.Face))
	}

	var count int = 1
	for {
		duration := (count * count) / 3
		for i := 0; i < len(d); i++ {
			d[i].Roll()
		}
		for i := 0; i < len(d); i++ {
			fmt.Print(d[i])
		}
		fmt.Print("\r")
		time.Sleep(time.Duration(duration) * time.Millisecond)
		count++

		if 30 < count {
			conf.EffectFlag = false
		}

		if !conf.EffectFlag {
			break
		}
	}

	fmt.Print("\n")

	if conf.ShowTotalFlag {
		var total uint = 0
		for i := 0; i < len(d); i++ {
			total += d[i].Dice
		}
		fmt.Printf("total: %d\n", total)
	}

}
