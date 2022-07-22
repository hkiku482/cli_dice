package config_test

import (
	"testing"

	"github.com/H-Kiku482/cli_dice/internal/config"
)

func TestFlag(t *testing.T) {
	// for operation check
	dummyCmdArgs := []string{
		"-d", "3", "-e",
	}
	expected := &config.Options{
		DiceQuantity: 3,
		Face:         6,
		EffectFlag:   true,
	}

	conf := config.GetConfig(dummyCmdArgs)

	if *conf != *expected {
		t.Errorf("expected: %v actual: %v", expected, conf)
	}
}
