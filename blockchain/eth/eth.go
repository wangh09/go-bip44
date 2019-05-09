package eth

import (
	"encoding/hex"
	"github.com/tyler-smith/go-bip32"
	"github.com/wangh09/go-bip44/crypto"
	"math/big"
)
func Encode(b []byte) string {
	enc := make([]byte, len(b)*2+2)
	copy(enc, "0x")
	hex.Encode(enc[2:], b)
	return string(enc)
}
func DecompressPubKey(pubKeyStr []byte)(res []byte) {
	format := pubKeyStr[0]
	ybit := (format & 0x1) == 0x1
	format &= ^byte(0x1)
	X := new(big.Int).SetBytes(pubKeyStr[1:33])
	Y, err := crypto.DecompressPoint(X, ybit)
	if err != nil {

	}
	if format != 0x2 {
		println("error---")
		return
	}
	res = make([]byte, 64)
	copy(res[0:32], X.Bytes())
	copy(res[32:64], Y.Bytes())
	return
}
func GetEthAddressHex(key *bip32.Key) string {
	dPubKey := DecompressPubKey(key.Key)
	return Encode(crypto.Keccak256(dPubKey)[12:])
}
