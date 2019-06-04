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
	"crypto/aes"
	"crypto/cipher"
	"github.com/youhei-yp/wing/invar"
	"math/rand"
	"time"
)

/*
 * Description :
 * (1). use secure.GenAESKey() to generate random AES key
 *   secretkey, err := secure.GenAESKey()
 *
 * (2). use secure.AESEncrypt() to encrypt original data with secret key
 *   ciphertext, err := secure.AESEncrypt([]byte(secretkey), original)
 *
 * (3). use secure.AESDecrypt() to decrypt ciphertext with secret key
 *   original, err := secure.AESDecrypt([]byte(secretkey), ciphertext)
 *
 *   [CODE:]
 *
 *   secretkey := secure.GenAESKey()
 *   logger.I("secret key:", secretkey)
 *
 *   original := []byte("original-content")
 *   ciphertext, _ := secure.AESEncrypt([]byte(secretkey), original)
 *   logger.I("original string:", string(original))
 *   logger.I("ciphertext string:", ciphertext)
 *
 *   encrypted, _ := secure.AESDecrypt([]byte(secretkey), ciphertext)
 *   logger.I("encrypted string: ", encrypted)
 *
 *   [CODE]
 *
 * ---------------------------------------------------------------------------------
 *
 * By the way, you can use the AES encrypt or decrypt for other languages as follows:
 *
 *   [CODE: ]
 *
 *   /// AES for java (Android)
 *
 *   public String encryptByAES(String secretkey, String original) {
 *       try {
 *           // use md5 value as the real key
 *           byte[] b = secretkey.getBytes();
 *           MessageDigest md = MessageDigest.getInstance("MD5");
 *           byte[] hashed = md.digest(b);
 *
 *           // create an 16-byte initialization vector
 *           byte[] iv = new byte[] {
 *               0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f
 *           };
 *           AlgorithmParameterSpec spec = new IvParameterSpec(iv);
 *           SecretKeySpec keyspec = new SecretKeySpec(hashed), "AES");
 *
 *           // create cipher and initialize CBC vector
 *           Cipher ecipher = Cipher.getInstance("AES/CBC/PKCS5Padding");
 *           ecipher.init(Cipher.ENCRYPT_MODE, keyspec, spec);
 *
 *           byte[] plaintext = original.getBytes();
 *           byte[] ciphertext = ecipher.doFinal(plaintext, 0, plaintext.length);
 *
 *           return Base64.encodeToString(ciphertext, Base64.DEFAULT);
 *       } catch (Exception e) {
 *           e.printStackTrace();
 *       }
 *       return null;
 *   }
 *
 *   public String decryptByAES(String secretkey, String ciphertextb64) {
 *       try {
 *           // use md5 value as the real key
 *           byte[] b = secretkey.getBytes();
 *           MessageDigest md = MessageDigest.getInstance("MD5");
 *           byte[] hashed = md.digest(b);
 *
 *           // create an 16-byte initialization vector
 *           byte[] iv = new byte[] {
 *               0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f
 *           };
 *           AlgorithmParameterSpec spec = new IvParameterSpec(iv);
 *           SecretKeySpec keyspec = new SecretKeySpec(hashed), "AES");
 *
 *           // create cipher and initialize CBC vector
 *           Cipher dcipher = Cipher.getInstance("AES/CBC/PKCS5Padding");
 *           dcipher.init(Cipher.DECRYPT_MODE, keyspec, spec);
 *
 *           byte[] ciphertext = Base64.decode(ciphertextb64, Base64.DEFAULT);
 *           byte[] original = dcipher.doFinal(ciphertext, 0, ciphertext.length);
 *
 *           return new String(original);
 *       } catch (Exception e) {
 *           e.printStackTrace();
 *       }
 *       return null;
 *   }
 *
 *   /// AES for node.js
 *
 *   let iv = [ 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f ];
 *
 *   function encrypt_by_aes(secretkey, original) {
 *       let md5 = crypto.createHash('md5').update(secretkey).digest('hex');
 *       const ecipher = crypto.createCipheriv(
 *           'aes-128-cbc',
 *           new Buffer(md5, 'hex'),
 *           new Buffer(iv)
 *       );
 *       // ecipher.setAutoPadding(true);
 *       var ciphertextb64 = ecipher.update(original, 'utf8', 'base64');
 *       ciphertextb64 += ecipher.final('base64');
 *       console.log('ciphertextb64: ' + ciphertextb64);
 *       return ciphertextb64;
 *   }
 *
 *   function decrypt_by_aes(secretkey, ciphertextb64) {
 *       let md5 = crypto.createHash('md5').update(secretkey).digest('hex');
 *       const dcipher = crypto.createDecipheriv(
 *           'aes-128-cbc',
 *           new Buffer(md5, 'hex'),
 *           new Buffer(iv)
 *       );
 *       var original = dcipher.update(ciphertextb64, 'base64', 'utf8');
 *       original += dcipher.final('utf8');
 *       console.log('original: ' + original);
 *       return original;
 *   }
 *
 *   [CODE]
 *
 */

var (
	// aesKeySeeds use to create secret key for ase crypto
	aesKeySeeds = []rune("abcdefghijklmnopqrstuvwxyz1234567890")

	// aesInitVector initialization vector for ase crypto
	aesInitVector = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

	// aesKeyLength key string length of AES
	aesKeyLength = len(aesInitVector)
)

// GenAESKey generate AES key range chars in [0-9a-z]{16}
func GenAESKey() string {
	rand.Seed(time.Now().UnixNano())
	secretkey := make([]rune, aesKeyLength)
	sendslen := (int32)(len(aesKeySeeds))
	for i := 0; i < aesKeyLength; i++ {
		j := rand.Int31n(sendslen)
		secretkey[i] = aesKeySeeds[j]
	}
	return string(secretkey)
}

// AESEncrypt using secret key to encrypt original data
func AESEncrypt(secretkey, original []byte) (string, error) {
	if len(secretkey) != aesKeyLength {
		return "", invar.ErrKeyLenSixteen.Err
	}

	hashed := HashMD5(secretkey)
	block, err := aes.NewCipher(hashed)
	if err != nil {
		return "", err
	}

	enc := cipher.NewCBCEncrypter(block, aesInitVector)
	content := pkcs5Padding(original, block.BlockSize())
	crypted := make([]byte, len(content))
	enc.CryptBlocks(crypted, content)
	return EncodeBase64(string(crypted)), nil
}

// AESDecrypt using secret key to decrypt ciphertext
func AESDecrypt(secretkey []byte, ciphertextb64 string) (string, error) {
	if len(secretkey) != aesKeyLength {
		return "", invar.ErrKeyLenSixteen.Err
	}

	hashed := HashMD5(secretkey)
	block, err := aes.NewCipher(hashed)
	if err != nil {
		return "", err
	}

	ciphertext, err := DecodeBase64(ciphertextb64)
	if err != nil {
		return "", err
	}

	dec := cipher.NewCBCDecrypter(block, aesInitVector)
	decrypted := make([]byte, len(ciphertext))
	dec.CryptBlocks(decrypted, []byte(ciphertext))
	return string(pkcs5Unpadding(decrypted)), nil
}

// pkcs5Padding use to padding the space of data
func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// pkcs5Unpadding use to unpadding the space of data
func pkcs5Unpadding(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
