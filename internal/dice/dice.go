package dice

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
)

type Dice struct {
	Dice  uint
	faces uint
}

func (d Dice) String() string {
	str := ""
	digits := len(strconv.Itoa(int(d.faces)))
	format := "[%0" + strconv.Itoa(digits) + "d]"
	str += fmt.Sprintf(format, d.Dice)
	return str
}

func InitDice(faces uint) *Dice {
	d := new(Dice)
	d.faces = faces
	return d
}

func (d *Dice) Roll() error {
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(d.faces-1)))
	if err != nil {
		return err
	}
	d.Dice = uint(nBig.Uint64()) + 1
	return nil
}
