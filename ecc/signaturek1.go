package ecc

import (
	"github.com/eoscochain/eos-go/btcsuite/btcd/btcec"
	"github.com/eoscochain/eos-go/btcsuite/btcutil/base58"
)

type innerK1Signature struct {
}

// verify checks the signature against the pubKey. `hash` is a sha256
// hash of the payload to verify.
func (s *innerK1Signature) verify(content []byte, hash []byte, pubKey PublicKey) bool {
	recoveredKey, _, err := btcec.RecoverCompact(btcec.S256(), content, hash)
	if err != nil {
		return false
	}
	key, err := pubKey.Key()
	if err != nil {
		return false
	}
	if recoveredKey.IsEqual(key) {
		return true
	}
	return false
}

func (s *innerK1Signature) publicKey(content []byte, hash []byte) (out PublicKey, err error) {

	recoveredKey, _, err := btcec.RecoverCompact(btcec.S256(), content, hash)

	if err != nil {
		return out, err
	}

	return PublicKey{
		Curve:   CurveK1,
		Content: recoveredKey.SerializeCompressed(),
		Prefix:  PublicKeyPrefixCompat,
	}, nil
}

func (s innerK1Signature) string(content []byte) string {
	checksum := Ripemd160checksumHashCurve(content, CurveK1)
	buf := append(content[:], checksum...)
	return "SIG_K1_" + base58.Encode(buf)
}
