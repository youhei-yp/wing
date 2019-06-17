// Copyright (c) 2018-2019 Dunyu All Rights Reserved.
//
// Author : yangping
// Email  : youhei_yp@163.com
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2019/05/22   yangping       New version
// -------------------------------------------------------------------

package comm

import (
	"fmt"
	"github.com/youhei-yp/wing/logger"
	"github.com/youhei-yp/wing/secure"
	"os"
)

// SaveFile save file buffer to target file
func SaveFile(filepath, filename string, data []byte) error {
	logger.I("Save file:", filename, "to dir:", filepath)

	// ensure path exist
	if _, err := os.Stat(filepath); err != nil {
		if os.IsNotExist(err) {
			if err = os.MkdirAll(filepath, os.ModePerm); err != nil {
				logger.E("Make path err:", err)
				return err
			}
		} else {
			logger.E("Stat path err:", err)
			return err
		}
	}

	// ensure file create or open success
	isFileExsit := true
	file := fmt.Sprintf("%s/%s", filepath, filename)
	if _, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			isFileExsit = false
		} else {
			logger.E("Stat file err:", err)
			return err
		}
	}

	var err error
	var desFile *os.File
	if !isFileExsit {
		desFile, err = os.Create(file)
		if err != nil {
			logger.E("Create file err:", err)
			return err
		}
	} else {
		desFile, err = os.Open(filepath)
		if err != nil {
			logger.E("Open file err:", err)
			return err
		}
	}
	defer desFile.Close()

	// write buffer to target file
	if _, err := desFile.Write(data); err != nil {
		logger.E("Write file buffer err:", err)
		return err
	}
	logger.I("Saved file:", filepath+"/"+filename)
	return nil
}

// SaveB64File save base64 encoded buffer to target file
func SaveB64File(filepath, filename string, b64data string) error {
	data, err := secure.DecodeBase64(b64data)
	if err != nil {
		logger.E("Invalid base64 data, err:", err)
		return err
	}
	return SaveFile(filepath, filename, []byte(data))
}

// DeleteFile delete file
func DeleteFile(file string) error {
	// ensure file exist
	if _, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			logger.I("Delete unexist file:", file)
			return nil
		}
		logger.E("Stat file err:", err)
		return err
	}

	if err := os.Remove(file); err != nil {
		logger.E("Delete file err:", err)
		return err
	}
	logger.I("Deleted file:", file)
	return nil
}
