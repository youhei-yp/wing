// Copyright (c) 2018-2019 Dunyu All Rights Reserved.
//
// Author : yangping
// Email  : youhei_yp@163.com
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2019/05/22   yangping       New version
// -------------------------------------------------------------------

package main

import (
	"flag"
	"github.com/youhei-yp/wing/secure"
	"log"
	"os"
)

var cmd = flag.String("c", "", "command type")
var pemfile = flag.String("f", "", "pem file path")
var seckey = flag.String("k", "", "secure key content")
var data = flag.String("d", "", "original or ciphertext data")

func main() {
	flag.Parse()

	if *cmd == "" {
		showUseage()
		return
	}

	switch *cmd {
	case "rsa-k":
		genRSAKeys()
	case "rsa-e":
		rsaEncrypt()
	case "rsa-d":
		rsaDecrypt()
	case "aes-k":
		genAESKey()
	case "aes-e":
		aesEncrypt()
	case "aes-d":
		aesDecrypt()
	case "uuid":
		genUUID()
	case "code":
		genCode()
	case "salt":
		genSalt()
	default:
		showUseage()
	}
}

func showUseage() {
	log.Print(`
	useage:
	===========================================================
	-c  command type, [rsa-k|rsa-e|rsa-d|aes-k|aes-e|aes-d|uuid|code|salt]
	-f  pem file path
	-k  secure key content
	-d	original data

	e.g:
	-----------------------------------------------------------
	./tools -c rsa-e -f ./pubkey.pem -d original
	./tools -c rsa-d -f ./prikey.pem -d ciphertext-base64
	./tools -c aes-e -k secure-key   -d original
	./tools -c aes-d -k secure-key   -d ciphertext-base64
	./tools -c rsa-k
	./tools -c aes-k
	./tools -c uuid
	./tools -c code
	./tools -c salt
	`)
}

func genRSAKeys() {
	prikey, pubkey, err := secure.GenRSAKeys(1024)
	if err != nil {
		log.Println("Gen RSA keys err:" + err.Error())
		return
	}
	log.Println("\n" + prikey)
	log.Println("\n" + pubkey)
}

func genAESKey() {
	pubkey := secure.GenAESKey()
	log.Println("AES key:", pubkey)
}

func genUUID() {
	uuid := secure.GenUUIDString()
	log.Println("UUID:", uuid)
}

func genCode() {
	code := secure.GenCode()
	log.Println("Code:", code)
}

func genSalt() {
	salt, _ := secure.GenSalt()
	log.Println("Salt:", salt)
}

func rsaEncrypt() {
	if *pemfile == "" || *data == "" {
		showUseage()
		return
	}

	if _, err := os.Stat(*pemfile); err != nil {
		if os.IsNotExist(err) {
			log.Println("[ERR] file unexist! " + *pemfile)
		}
		return
	}

	pubkey, err := secure.LoadRSAKey(*pemfile, 2048)
	if err != nil {
		log.Println("[ERR] load RSA pem file err:" + err.Error())
		return
	}

	ciphertext, err := secure.RSAEncrypt(pubkey, []byte(*data))
	if err != nil {
		log.Println("[ERR] encrypt err:" + err.Error())
		return
	}

	ciphertextb64 := secure.EncodeBase64(string(ciphertext))
	log.Println("Ciphertext Base64:" + ciphertextb64)
}

func rsaDecrypt() {
	if *pemfile == "" || *data == "" {
		showUseage()
		return
	}

	if _, err := os.Stat(*pemfile); err != nil {
		if os.IsNotExist(err) {
			log.Println("[ERR] file unexist! " + *pemfile)
		}
		return
	}

	prikey, err := secure.LoadRSAKey(*pemfile, 2048)
	if err != nil {
		log.Println("[ERR] load RSA pem file err:" + err.Error())
		return
	}

	ciphertext, err := secure.DecodeBase64(*data)
	if err != nil {
		log.Println("[ERR] invalid base64 ciphertext:" + err.Error())
		return
	}

	original, err := secure.RSADecrypt(prikey, []byte(ciphertext))
	if err != nil {
		log.Println("[ERR] decrypt err:" + err.Error())
		return
	}
	log.Println("Original:" + string(original))
}

func aesEncrypt() {
	if *seckey == "" || *data == "" {
		showUseage()
		return
	}

	ciphertextb64, err := secure.AESEncrypt([]byte(*seckey), []byte(*data))
	if err != nil {
		log.Println("[ERR] encrypt err:" + err.Error())
		return
	}
	log.Println("Ciphertext Base64:" + ciphertextb64)
}

func aesDecrypt() {
	if *seckey == "" || *data == "" {
		showUseage()
		return
	}

	original, err := secure.AESDecrypt([]byte(*seckey), *data)
	if err != nil {
		log.Println("[ERR] encrypt err:" + err.Error())
		return
	}
	log.Println("Original:" + original)
}
