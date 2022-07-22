package config

import "github.com/jessevdk/go-flags"

type Options struct {
	DiceQuantity uint `short:"d" long:"dice" default:"1" description:"Set dice quantity"`
	Face         uint `short:"f" long:"face" default:"1" description:"Specify the number of faces"`
	EffectFlag   bool `short:"e" long:"effect" description:"Output with effect"`
}

func GetConfig(cmdArgs []string) (*Options, error) {
	opt := new(Options)

	_, err := flags.ParseArgs(opt, cmdArgs)
	if err != nil {
		return nil, err
	}

	return &Options{
		DiceQuantity: opt.DiceQuantity,
		Face:         opt.Face,
		EffectFlag:   opt.EffectFlag,
	}, nil
}
