package wg

import (
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

// creates nmew WireGuard private/public key pair
func GenerateKeyPair() (privateKey, publicKey string, err error) {
	privKey, err := wgtypes.GeneratePrivateKey()
	if err != nil {
		return "", "", err
	}

	pubKey := privKey.PublicKey()

	return privKey.String(), pubKey.String(), nil
}

// derives public key from private key
func PublicKeyFromPrivate(privateKey string) (string, error) {
	privKey, err := wgtypes.ParseKey(privateKey)
	if err != nil {
		return "", err
	}

	pubKey := privKey.PublicKey()
	return pubKey.String(), nil
}
