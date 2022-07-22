package config

import (
	"github.com/jessevdk/go-flags"
)

type Options struct {
	DiceQuantity uint `short:"d" long:"dice" default:"1" description:"Set dice quantity"`
	Face         uint `short:"f" long:"face" default:"6" description:"Specify the number of faces"`
	EffectFlag   bool `short:"e" long:"effect" description:"Output with effect"`
}

func GetConfig(cmdArgs []string) *Options {
	opt := new(Options)

	flags.ParseArgs(opt, cmdArgs)

	return &Options{
		DiceQuantity: opt.DiceQuantity,
		Face:         opt.Face,
		EffectFlag:   opt.EffectFlag,
	}
}
