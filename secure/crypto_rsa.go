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
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/youhei-yp/wing/invar"
	"io/ioutil"
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

// GenRSAKeys generate RSA private and public keys in PKCS#1, ASN.1 DER form,
// and limit bits length of key cert.
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

// LoadRSAKey load RSA or RSA2 private or public key content from the given
// pem file, and if input the beffer size, the bufbits must larger than pem
// file size by call GenRSAKeys to set bits.
func LoadRSAKey(filepath string, buffbits ...int) ([]byte, error) {
	if buffbits != nil && len(buffbits) > 0 && buffbits[0] > 0 {
		pemfile, err := os.Open(filepath)
		if err != nil {
			return nil, err
		}
		defer pemfile.Close()

		keybuf := make([]byte, buffbits[0])
		num, err := pemfile.Read(keybuf)
		if err != nil {
			return nil, err
		}
		return keybuf[:num], nil
	} else {
		pemfile, err := ioutil.ReadFile(filepath)
		if err != nil {
			return nil, err
		}
		return pemfile, nil
	}
}

// RSAEncrypt using RSA public key to encrypt original data.
func RSAEncrypt(pubkey, original []byte) ([]byte, error) {
	block, _ := pem.Decode(pubkey)
	if block == nil {
		return nil, invar.ErrBadPublicKey
	}

	pubinterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubinterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, original)
}

// RSAEncryptB64 using RSA public key to encrypt original data,
// then format to base64 form.
func RSAEncryptB64(pubkey, original []byte) (string, error) {
	buf, err := RSAEncrypt(pubkey, original)
	if err != nil {
		return "", nil
	}
	return ByteToBase64(buf), nil
}

// RSAEncrypt4F using RSA public key file to encrypt original data.
func RSAEncrypt4F(pubfile string, original []byte) ([]byte, error) {
	pubkey, err := LoadRSAKey(pubfile)
	if err != nil {
		return nil, err
	}
	return RSAEncrypt(pubkey, original)
}

// RSAEncrypt4FB64 using RSA public key file to encrypt original data,
// then format to base64 form.
func RSAEncrypt4FB64(pubfile string, original []byte) (string, error) {
	buf, err := RSAEncrypt4F(pubfile, original)
	if err != nil {
		return "", err
	}
	return ByteToBase64(buf), nil
}

// RSADecrypt using RSA private key to decrypt ciphertext.
func RSADecrypt(prikey, ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(prikey)
	if block == nil {
		return nil, invar.ErrBadPrivateKey
	}

	pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, pri, ciphertext)
}

// RSADecrypt4F using RSA private key file to decrypt ciphertext.
func RSADecrypt4F(prifile string, ciphertext []byte) ([]byte, error) {
	prikey, err := LoadRSAKey(prifile)
	if err != nil {
		return nil, err
	}
	return RSADecrypt(prikey, ciphertext)
}

// RSASign using RSA private key to make digital signature,
// the private key in PKCS#1, ASN.1 DER form.
func RSASign(prikey, original []byte) ([]byte, error) {
	block, _ := pem.Decode(prikey)
	if block == nil {
		return nil, invar.ErrBadPrivateKey
	}

	pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	hashed := HashSHA256(original)
	return rsa.SignPKCS1v15(rand.Reader, pri, crypto.SHA256, hashed)
}

// RSASignB64 using RSA private key to make digital signature,
// then format to base64 form, the private key in PKCS#1, ASN.1 DER form.
func RSASignB64(prikey, original []byte) (string, error) {
	buf, err := RSASign(prikey, original)
	if err != nil {
		return "", err
	}
	return ByteToBase64(buf), nil
}

// RSASign4F using RSA private key file to make digital signature,
// the private key in PKCS#1, ASN.1 DER form.
func RSASign4F(prifile string, original []byte) ([]byte, error) {
	prikey, err := LoadRSAKey(prifile)
	if err != nil {
		return nil, err
	}
	return RSASign(prikey, original)
}

// RSASign4FB64 using RSA private key file to make digital signature,
// then format to base64 form, the private key in PKCS#1, ASN.1 DER form.
func RSASign4FB64(prifile string, original []byte) (string, error) {
	buf, err := RSASign4F(prifile, original)
	if err != nil {
		return "", err
	}
	return ByteToBase64(buf), nil
}

// RSAVerify using RSA public key to verify signatured data.
func RSAVerify(pubkey, original, signature []byte) error {
	block, _ := pem.Decode(pubkey)
	if block == nil {
		return invar.ErrBadPublicKey
	}

	pubinterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	pub := pubinterface.(*rsa.PublicKey)
	hashed := HashSHA256(original)
	return rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed[:], signature)
}

// RSAVerify4F using RSA public key file to verify signatured data.
func RSAVerify4F(pubfile string, original, signature []byte) error {
	pubkey, err := LoadRSAKey(pubfile)
	if err != nil {
		return err
	}
	return RSAVerify(pubkey, original, signature)
}

// RSA2Sign using RSA2 private key to make digital signature,
// the private key in PKCS#8, ASN.1 DER form.
func RSA2Sign(prikey, original []byte) ([]byte, error) {
	block, _ := pem.Decode(prikey)
	if block == nil {
		return nil, invar.ErrBadPrivateKey
	}

	priinterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	hashed := HashSHA256(original)
	pri := priinterface.(*rsa.PrivateKey)
	return rsa.SignPKCS1v15(rand.Reader, pri, crypto.SHA256, hashed)
}

// RSA2SignB64 using RSA2 private key file to make digital signature,
// then format to base64 form, the private key in PKCS#8, ASN.1 DER form.
func RSA2SignB64(prikey, original []byte) (string, error) {
	buf, err := RSA2Sign(prikey, original)
	if err != nil {
		return "", err
	}
	return ByteToBase64(buf), nil
}

// RSA2Sign4F using RSA2 private key file to make digital signature.
// the private key in PKCS#8, ASN.1 DER form.
func RSA2Sign4F(prifile string, original []byte) ([]byte, error) {
	prikey, err := LoadRSAKey(prifile)
	if err != nil {
		return nil, err
	}
	return RSA2Sign(prikey, original)
}

// RSA2Sign4FB64 using RSA2 private key file to make digital signature,
// then format to base64 form, the private key in PKCS#8, ASN.1 DER form.
func RSA2Sign4FB64(prifile string, original []byte) (string, error) {
	buf, err := RSA2Sign4F(prifile, original)
	if err != nil {
		return "", err
	}
	return ByteToBase64(buf), nil
}

// RSA2Verify using RSA2 public key to verify signatured data.
func RSA2Verify(pubkey, original, signature []byte) error {
	return RSA2Verify(pubkey, original, signature)
}

// RSA2Verify4F using RSA2 public key to verify signatured data.
func RSA2Verify4F(pubfile string, original, signature []byte) error {
	return RSA2Verify4F(pubfile, original, signature)
}
