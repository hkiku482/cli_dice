package main

import (
	"fmt"
	"os"

	"github.com/H-Kiku482/cli_dice/internal/config"
)

func main() {
	conf, err := config.GetConfig(os.Args)
	if err != nil {
		print(err)
	}
	fmt.Println(conf)
}
