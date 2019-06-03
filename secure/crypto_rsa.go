// Copyright (c) 2018-2019 Dunyu All Rights Reserved.
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
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/youhei-yp/wing/invar"
	"os"
)

/*
 * Description :
 * (1). use secure.GenRSAKeys() to generate RSA keys, and set content bits length
 *   prikey, pubkey, err := secure.GenRSAKeys(bits)
 *
 * (2). use secure.RSAEncrypt() to encrypt original data with given public key
 *   ciphertext, err := secure.RSAEncrypt(pubkey, original)
 *
 * (3). use secure.RSADecrypt() to decrypt ciphertext with given private key
 *   original, err := secure.RSADecrypt(prikey, ciphertext)
 *
 *   [CODE:]
 *   // Use the pubkey to encrypt and use the prikey to decrypt
 *
 *   prikey, pubkey, _ := secure.GenRSAKeys(1024)
 *   logger.I("private key:", prikey)
 *   logger.I("public  key:", pubkey)
 *
 *   ciphertext, _ := secure.RSAEncrypt([]byte(pubkey), []byte("original-content"))
 *   ciphertextBase64 := secure.EncodeBase64(string(ciphertext))
 *   logger.I("ciphertext base64 string:", ciphertextBase64)
 *
 *   original, _ := secure.RSADecrypt([]byte(prikey), ciphertext)
 *   logger.I("original string:", string(original))	// print 'original-content'
 *
 *   [CODE]
 *
 * Description :
 * (1). use secure.GenRSAKeys() to generate RSA keys, and set content bits length
 *   prikey, pubkey, err := secure.GenRSAKeys(bits)
 *
 * (2). use secure.RSASign() to make digital signature with given private key
 *   signature, err := secure.RSASign(prikey, original)
 *
 * (3). use secure.RSAVerify() to verify data's integrity with given public key and digital signature
 *   err := secure.RSAVerify(pubkey, original, signature)
 *
 *   [CODE:]
 *   // Use the private key to create digital signature and use pubkey to verify it
 *
 *   prikey, pubkey, _ := secure.GenRSAKeys(1024)
 *   logger.I("private key:", prikey)
 *   logger.I("public  key:", pubkey)
 *
 *   original := []byte("original-content")
 *   signature, _ := secure.RSASign([]byte(prikey), original)
 *   logger.I("original string:", string(original))
 *   logger.I("signature string:", string(signature))
 *
 *   if err := secure.RSAVerify([]byte(pubkey), original, signature); err != nil {
 *       logger.E("Verify failed with err:", err)
 *       return
 *   }
 *   logger.I("Verify success")
 *
 *   [CODE]
 */

const (
	blockRsaPrivateKey = "RSA Private key"
	blockRsaPublicKey  = "RSA Public key"
)

// GenRSAKeys generate RSA private and public keys that limit bits length
func GenRSAKeys(bits int) (string, string, error) {
	// generate private key
	prikey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return "", "", err
	}

	// create buffer to save private pem content
	pribuff := new(bytes.Buffer)
	derstream := x509.MarshalPKCS1PrivateKey(prikey)
	block := &pem.Block{Type: blockRsaPrivateKey, Bytes: derstream}
	if err = pem.Encode(pribuff, block); err != nil {
		return "", "", err
	}

	pubkey := &prikey.PublicKey
	derpkix, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return "", "", err
	}

	pubbuff := new(bytes.Buffer)
	block = &pem.Block{Type: blockRsaPublicKey, Bytes: derpkix}
	if err = pem.Encode(pubbuff, block); err != nil {
		return "", "", err
	}

	return pribuff.String(), pubbuff.String(), nil
}

// LoadRSAKey load RSA private or public key content from the given pem file,
// and the bufbits must larger than pem file size as call GenRSAKeys to set bits.
func LoadRSAKey(filepath string, buffbits int) ([]byte, error) {
	pemfile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer pemfile.Close()

	keybuf := make([]byte, buffbits)
	num, err := pemfile.Read(keybuf)
	if err != nil {
		return nil, err
	}
	return keybuf[:num], nil
}

// RSAEncrypt using RSA to encrypt original data
func RSAEncrypt(pubkey, original []byte) ([]byte, error) {
	block, _ := pem.Decode(pubkey)
	if block == nil {
		return nil, invar.ErrBadPublicKey.Err
	}

	pubinterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubinterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, original)
}

// RSADecrypt using RSA to decrypt ciphertext
func RSADecrypt(prikey, ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(prikey)
	if block == nil {
		return nil, invar.ErrBadPrivateKey.Err
	}

	pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, pri, ciphertext)
}

// RSASign using RSA private key to make digital signature.
func RSASign(prikey, original []byte) ([]byte, error) {
	block, _ := pem.Decode(prikey)
	if block == nil {
		return nil, invar.ErrBadPrivateKey.Err
	}

	pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	hashed := HashSHA256(original)
	return rsa.SignPKCS1v15(rand.Reader, pri, crypto.SHA256, hashed)
}

// RSAVerify using public key to verify signatured data
func RSAVerify(pubkey, original, signature []byte) error {
	block, _ := pem.Decode(pubkey)
	if block == nil {
		return invar.ErrBadPublicKey.Err
	}

	pubinterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	pub := pubinterface.(*rsa.PublicKey)
	hashed := HashSHA256(original)
	return rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed[:], signature)
}
