package dice_test

import (
	"fmt"
	"testing"

	"github.com/H-Kiku482/cli_dice/internal/dice"
)

func TestDiceString(t *testing.T) {
	// for operation check
	dice := dice.InitDice(6)
	expected := "[0]"
	actual := fmt.Sprint(dice)
	if expected != actual {
		t.Errorf("expected: %v actual: %v", expected, actual)
	}
}
