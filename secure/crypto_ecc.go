// Copyright (c) 2018-2019 WING All Rights Reserved.
//
// Author : yangping
// Email  : youhei_yp@163.com
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2019/05/22   yangping       New version
// -------------------------------------------------------------------

package secure

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"github.com/youhei-yp/wing/invar"
	"math/big"
)

// GenECCPriKey generate a ECC private key by origin key text
func GenECCPriKey(prikey string) (*ecdsa.PrivateKey, error) {
	keylen := len(prikey)
	if keylen != 97 {
		return nil, invar.ErrInvalidData
	}

	ecckey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	x, y, d := make([]byte, 32), make([]byte, 32), make([]byte, 32)
	for i := 1; i < keylen; i++ {
		if i < 33 {
			x[i-1] = prikey[i]
		} else if i < 65 {
			y[i-33] = prikey[i]
		} else {
			d[i-65] = prikey[i]
		}
	}

	ecckey.D.SetBytes(d)
	ecckey.Public().(*ecdsa.PublicKey).X.SetBytes(x)
	ecckey.Public().(*ecdsa.PublicKey).Y.SetBytes(y)
	return ecckey, nil
}

// GenECCPubKey generate a ECC public key by origin key text
func GenECCPubKey(pubkey string) (*ecdsa.PublicKey, error) {
	keylen := len(pubkey)
	if keylen != 65 {
		return nil, invar.ErrInvalidData
	}

	ecckey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	x, y := make([]byte, 32), make([]byte, 32)
	for i := 1; i < len(pubkey); i++ {
		if i < 33 {
			x[i-1] = pubkey[i]
		} else {
			y[i-33] = pubkey[i]
		}
	}

	ecckey.Public().(*ecdsa.PublicKey).X.SetBytes(x)
	ecckey.Public().(*ecdsa.PublicKey).Y.SetBytes(y)
	return ecckey.Public().(*ecdsa.PublicKey), nil
}

// GenECCPriKeyB64 generate a ECC private key by base64 formated key text
func GenECCPriKeyB64(prikeyb64 string) (*ecdsa.PrivateKey, error) {
	prikey, err := DecodeBase64(prikeyb64)
	if err != nil {
		return nil, err
	}
	return GenECCPriKey(prikey)
}

// GenECCPubKeyB64 generate a ECC public key by base64 formated key text
func GenECCPubKeyB64(pubkeyb64 string) (*ecdsa.PublicKey, error) {
	pubkey, err := DecodeBase64(pubkeyb64)
	if err != nil {
		return nil, err
	}
	return GenECCPubKey(pubkey)
}

// GenECCShareKeys generate ECC shared keys by ECC public key and private digital signature
func GenECCShareKeys(pub *ecdsa.PublicKey, priD *big.Int) (*big.Int, *big.Int) {
	shareX, shareY := elliptic.P256().ScalarMult(pub.X, pub.Y, priD.Bytes())
	return shareX, shareY
}

// GenEccShareKeysHash generate ECC share keys hash data by origin private and public key
func GenEccShareKeysHash(prikey, pubkey string) ([]byte, error) {
	eprikey, err := GenECCPriKey(prikey)
	if err != nil {
		return nil, err
	}

	epubkey, err := GenECCPubKey(pubkey)
	if err != nil {
		return nil, err
	}

	sharex, sharey := GenECCShareKeys(epubkey, eprikey.D)
	bx, by := sharex.Bytes(), sharey.Bytes()
	bxlen, bylen := len(bx), len(by)

	sharekey := make([]byte, bxlen+bylen)
	for i := 0; i < bxlen; i++ {
		sharekey[i] = bx[i]
	}
	for j := 0; j < bylen; j++ {
		sharekey[bxlen+j] = by[j]
	}

	return HashSHA256(sharekey), nil
}

// GenEccShareKeysHashB64 generate ECC share keys hash data by base64 formated private and public key
func GenEccShareKeysHashB64(prikeyb64, pubkeyb64 string) ([]byte, error) {
	eprikey, err := GenECCPriKeyB64(prikeyb64)
	if err != nil {
		return nil, err
	}

	epubkey, err := GenECCPubKeyB64(pubkeyb64)
	if err != nil {
		return nil, err
	}

	sharex, sharey := GenECCShareKeys(epubkey, eprikey.D)
	bx, by := sharex.Bytes(), sharey.Bytes()
	bxlen, bylen := len(bx), len(by)

	sharekey := make([]byte, bxlen+bylen)
	for i := 0; i < bxlen; i++ {
		sharekey[i] = bx[i]
	}
	for j := 0; j < bylen; j++ {
		sharekey[bxlen+j] = by[j]
	}

	return HashSHA256(sharekey), nil
}

// GenRSFromB2BI generate R and S from sign data
func GenRSFromB2BI(sign []byte) (*big.Int, *big.Int) {
	if len(sign) != 64 {
		zero := big.NewInt(0)
		return zero, zero
	}

	rb, sb := make([]byte, 32), make([]byte, 32)
	for i := 0; i < len(sign); i++ {
		if i < 32 {
			rb[i] = sign[i]
		} else {
			sb[i-32] = sign[i]
		}
	}

	r, s := new(big.Int), new(big.Int)
	return r.SetBytes(rb), s.SetBytes(sb)
}

// ECCHashSignB64 use ECC to sign hash data to base64 string by given origin private key text
func ECCHashSignB64(hash []byte, prikeyb64 string) (string, error) {
	eprikey, err := GenECCPriKeyB64(prikeyb64)
	if err != nil {
		return "", err
	}

	r, s, err := ecdsa.Sign(rand.Reader, eprikey, hash)
	if err != nil {
		return "", err
	}

	sign, sb := r.Bytes(), s.Bytes()
	for i, l := 0, len(sb); i < l; i++ {
		sign = append(sign, sb[i])
	}
	return ByteToBase64(sign), nil
}

// ECCHashVerifyB64 use ECC public key and rs data to verify given hash data
func ECCHashVerifyB64(hash []byte, pubkeyb64 string, rsb64 string) (bool, error) {
	epubkey, err := GenECCPubKeyB64(pubkeyb64)
	if err != nil {
		return false, err
	}

	rs, err := Base64ToByte(rsb64)
	if err != nil {
		return false, err
	}

	if len(rs) != 64 {
		return false, invar.ErrInvalidData
	}

	r, s := GenRSFromB2BI(rs)
	return ecdsa.Verify(epubkey, hash, r, s), nil
}
