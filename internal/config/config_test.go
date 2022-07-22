package config_test

import (
	"testing"

	"github.com/H-Kiku482/cli_dice/internal/config"
)

func OperationCheckGetConfig(t *testing.T) {
	dummyCmdArgs := []string{
		"-d", "3", "-e",
	}
	expected := &config.Options{
		DiceQuantity: 3,
		EffectFlag:   true,
	}

	conf, err := config.GetConfig(dummyCmdArgs)
	if err != nil {
		print(err)
	}

	if *conf != *expected {
		t.Errorf("expected: %v actual: %v", expected, conf)
	}
}
