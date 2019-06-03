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
 */

/*
 * Note :
 * 			By the way, you can use the AES encrypt and decrypt on other language,
 *		here is the code
 *
 *   [CODE: AES for java(Android)]
 *
 *		public String encryptWithAES(String key, String message) {
 *		    try {
 *		        // Use md5 value as the real key
 *		        byte[] b = key.getBytes();
 *		        MessageDigest md = MessageDigest.getInstance("MD5");
 *		        byte[] keyData = md.digest(b);
 *
 *		        SecretKeySpec skey = new SecretKeySpec(keyData), "AES");
 *		        // Create an 8-byte initialization vector
 *		        byte[] iv = new byte[] { 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d,
 *		                0x0e, 0x0f };
 *		        AlgorithmParameterSpec paramSpec = new IvParameterSpec(iv);
 *
 *		        Cipher ecipher = Cipher.getInstance("AES/CBC/PKCS5Padding");
 *		        // CBC requires an initialization vector
 *		        ecipher.init(Cipher.ENCRYPT_MODE, skey, paramSpec);
 *
 *		        byte[] plaintext = message.getBytes();
 *		        byte[] result = ecipher.doFinal(plaintext, 0, plaintext.length);
 *
 *		        return Base64.encodeToString(result, Base64.DEFAULT);
 *		    } catch (Exception e) {
 *		        e.printStackTrace();
 *		    }
 *		    return null;
 *		}
 *
 *		public String decryptWithAES(String key, String message) {
 *		    try {
 *		        // Use md5 value as the real key
 *		        byte[] b = key.getBytes();
 *		        MessageDigest md = MessageDigest.getInstance("MD5");
 *		        byte[] keyData = md.digest(b);
 *
 *		        SecretKeySpec skey = new SecretKeySpec(keyData), "AES");
 *		        // Create an 8-byte initialization vector
 *		        byte[] iv = new byte[] { 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d,
 *		                0x0e, 0x0f };
 *		        AlgorithmParameterSpec paramSpec = new IvParameterSpec(iv);
 *
 *		        Cipher dcipher = Cipher.getInstance("AES/CBC/PKCS5Padding");
 *		        // CBC requires an initialization vector
 *		        dcipher.init(Cipher.DECRYPT_MODE, skey, paramSpec);
 *
 *		        byte[] messageData = Base64.decode(message, Base64.DEFAULT);
 *		        byte[] result = dcipher.doFinal(messageData, 0, messageData.length);
 *
 *		        return new String(result);
 *		    } catch (Exception e) {
 *		        e.printStackTrace();
 *		    }
 *		    return null;
 *		}
 *
 *   [CODE]
 *
 *   [CODE: AES for node.js]
 *
 *			let iv = [0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d,
 *			    0x0e, 0x0f];
 *
 *			function encrypt_with_aes(key, message) {
 *			    let md5 = crypto.createHash('md5').update(key).digest('hex');
 *			    const cipher = crypto.createCipheriv(
 *			        'aes-128-cbc',
 *			        new Buffer(md5, 'hex'),
 *			        new Buffer(iv)
 *			    );
 *			    // cipher.setAutoPadding(true);
 *			    var encrypted = cipher.update(message, 'utf8', 'base64');
 *			    encrypted += cipher.final('base64');
 *			    console.log('encode message: ' + encrypted);
 *			    return encrypted;
 *			}
 *
 *			function decrypt_with_aes(key, message) {
 *			    let md5 = crypto.createHash('md5').update(key).digest('hex');
 *			    const decipher = crypto.createDecipheriv(
 *			        'aes-128-cbc',
 *			        new Buffer(md5, 'hex'),
 *			        new Buffer(iv)
 *			    );
 *			    var decrypted = decipher.update(message, 'base64', 'utf8');
 *			    decrypted += decipher.final('utf8');
 *			    console.log('decode message: ' + decrypted);
 *			    return decrypted;
 *			}
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
	content := PKCS5Padding(original, block.BlockSize())
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
	return string(PKCS5Unpadding(decrypted)), nil
}

// PKCS5Padding use to padding the space of data
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS5Unpadding use to unpadding the space of data
func PKCS5Unpadding(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
