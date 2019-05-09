package crypto

import (
	"fmt"
	"math/big"
)

func isOdd(a *big.Int) bool {
	return a.Bit(0) == 1
}
var curveP, curveB, curveQ = Init()
func Init()(P, B, q *big.Int) {
	P, _ = new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F", 16)
	B, _ = new(big.Int).SetString("0000000000000000000000000000000000000000000000000000000000000007", 16)
	q = new(big.Int).Div(new(big.Int).Add(P, big.NewInt(1)), big.NewInt(4))
	return
}

func DecompressPoint(x *big.Int, ybit bool) (*big.Int, error) {
	x3 := new(big.Int).Mul(x, x)
	x3.Mul(x3, x)
	x3.Add(x3, curveB)
	y := new(big.Int).Exp(x3, curveQ, curveP)

	if ybit != isOdd(y) {
		y.Sub(curveP, y)
	}
	if ybit != isOdd(y) {
		return nil, fmt.Errorf("ybit doesn't match oddness")
	}
	return y, nil
}
