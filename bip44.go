package bip44

import (
	hex2 "encoding/hex"
	"github.com/tyler-smith/go-bip32"
	"log"
)

const Purpose 	uint32 =  	0x8000002c
const Wallet 	uint32 = 	0x80000000

const (
	BTC		uint32 = 0x80000000
	ETH 	uint32 = 0x8000003c
)

func StringToBIP44CoinType(coin string) uint32 {
	switch coin {
	case "BTC":
		return BTC
	case "ETH":
		return ETH
	default:
		return 0
	}
}

func MasterKeyFromHex(hex string) *bip32.Key {
	masterKey, err := bip32.B58Deserialize(hex)
	if err != nil {
		log.Fatalln("Error generating seed:", err)
		return nil
	}
	println(hex2.EncodeToString( masterKey.Key))
	return masterKey
}

func NewMasterPrvKey() (*bip32.Key, error) {
	seed, err := bip32.NewSeed()
	if err != nil {
		return nil, err
	}
	return bip32.NewMasterKey(seed)
}

func NewBip44ChildKey(masterKey *bip32.Key, coinType string, bipIdx uint32) (*bip32.Key, error) {
	return newKeyFromMasterKey(masterKey, StringToBIP44CoinType(coinType), 0, bipIdx)
}

func newKeyFromMasterKey(masterKey *bip32.Key, coin, chain, address uint32) (*bip32.Key, error) {
	child, err := masterKey.NewChildKey(Purpose)
	if err != nil {
		return nil, err
	}
	child, err = child.NewChildKey(coin)
	if err != nil {
		return nil, err
	}
	child, err = child.NewChildKey(Wallet)
	if err != nil {
		return nil, err
	}
	child, err = child.NewChildKey(chain)
	if err != nil {
		return nil, err
	}
	child, err = child.NewChildKey(address)
	if err != nil {
		return nil, err
	}
	return child, nil
}