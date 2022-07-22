package dice

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"strconv"
)

type Dice struct {
	dice  uint
	faces uint
}

func (d *Dice) String() string {
	str := ""
	digits := strconv.Itoa(int(math.Floor(math.Log10(float64(d.faces)))))
	format := "[%" + digits + "d]"
	str += fmt.Sprintf(format, d.dice)
	return str
}

func GetDice(faces uint) *Dice {
	d := new(Dice)
	d.faces = faces
	return d
}

func (d *Dice) Roll() error {
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(d.faces)))
	if err != nil {
		return err
	}
	d.dice = uint(nBig.Uint64())
	return nil
}
