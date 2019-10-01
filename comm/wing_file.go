// Copyright (c) 2018-2019 WING All Rights Reserved.
//
// Author : yangping
// Email  : youhei_yp@163.com
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2019/05/22   yangping       New version
// 00002       2019/06/30   zhaixing       Add function from godfs
// -------------------------------------------------------------------

package comm

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/youhei-yp/wing/invar"
	"github.com/youhei-yp/wing/logger"
	"github.com/youhei-yp/wing/secure"
	"io"
	"os"
	"strconv"
	"strings"
	"syscall"
)

const (
	// truncate buffer size for file copy
	truncateBufferSize = 1024 * 30
)

// IsFile check whether file path point to a file.
func IsFile(filepath string) bool {
	if fileinfo, err := os.Stat(filepath); err == nil {
		return !fileinfo.IsDir()
	}
	return false
}

// IsDir check whether dir path point to a directory.
func IsDir(dirpath string) bool {
	if fileinfo, err := os.Stat(dirpath); err == nil {
		return fileinfo.IsDir()
	}
	return false
}

// IsFile2 check whether file point to a file.
func IsFile2(file *os.File) bool {
	if fileinfo, err := file.Stat(); err == nil {
		return !fileinfo.IsDir()
	}
	return false
}

// IsDir2 check whether file point to a directory.
func IsDir2(file *os.File) bool {
	if fileinfo, err := file.Stat(); err == nil {
		return fileinfo.IsDir()
	}
	return false
}

// IsExistFile check whether the file exists.
func IsExistFile(filepath string) bool {
	fileinfo, err := os.Stat(filepath)
	if err != nil || fileinfo == nil {
		return false
	}
	return true
}

// EnsurePath check the given file path, or create new one if unexist
func EnsurePath(filepath string) error {
	if _, err := os.Stat(filepath); err != nil {
		if os.IsNotExist(err) {
			if err = MakeDirs(filepath); err != nil {
				logger.E("Make path err:", err)
				return err
			}
		} else {
			logger.E("Stat path err:", err)
			return err
		}
	}
	return nil
}

// MakeDirs create new directory, by default perm is 0777.
func MakeDirs(dirpath string, perm ...os.FileMode) error {
	if len(perm) > 0 && perm[0] != 0 {
		return os.MkdirAll(dirpath, perm[0])
	}
	return os.MkdirAll(dirpath, os.ModePerm)
}

// OpenFileWrite create new file for write content, by default perm is 0666.
func OpenFileWrite(filepath string, perm ...os.FileMode) (*os.File, error) {
	if len(perm) > 0 && perm[0] != 0 {
		return os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, perm[0])
	}
	return os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
}

// OpenFileTrunc create new file for trunc content, by default perm is 0666.
func OpenFileTrunc(filepath string, perm ...os.FileMode) (*os.File, error) {
	if len(perm) > 0 && perm[0] != 0 {
		return os.OpenFile(filepath, syscall.O_CREAT|os.O_WRONLY|syscall.O_TRUNC, perm[0])
	}
	return os.OpenFile(filepath, syscall.O_CREAT|os.O_WRONLY|syscall.O_TRUNC, 0666)
}

// SaveFile save file buffer to target file
func SaveFile(filepath, filename string, data []byte) error {
	logger.I("Save file:", filename, "to dir:", filepath)

	// ensure path exist
	if err := EnsurePath(filepath); err != nil {
		return err
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
	return nil
}

// DeletePath delete files and directory.
func DeletePath(dirpath string) error {
	if _, err := os.Stat(dirpath); err != nil {
		if os.IsNotExist(err) {
			logger.I("Delete unexist dir:", dirpath)
			return nil
		}
		logger.E("Stat dir err:", err)
		return err
	}
	return os.RemoveAll(dirpath)
}

// FileMD5 encode file content to md5 string
func FileMD5(file string) (string, error) {
	h := md5.New()
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}

	defer f.Close()
	_, err = io.Copy(h, f)
	if err != nil {
		return "", err
	}
	cipher := h.Sum(nil)
	return hex.EncodeToString(cipher), nil
}

// CopyFile Copy source file to traget file.
func CopyFile(src string, dest string) (bool, error) {
	srcfile, err := os.Open(src)
	if err != nil {
		return false, err
	}
	defer srcfile.Close()

	// create or truncate dest file
	destfile, err := OpenFileTrunc(dest)
	if err != nil {
		return false, err
	}
	defer destfile.Close()

	// start copying
	result, err := false, nil
	Try(func() {
		buff := make([]byte, truncateBufferSize)
		for {
			len, state := srcfile.Read(buff)
			if state == io.EOF {
				break
			}
			destfile.Write(buff[0:len])
		}
		result = true
	}, func(execption error) {
		err = execption
	})
	return result, err
}

// CopyFileTo copy source file to given dir.
func CopyFileTo(src string, dir string) (bool, error) {
	srcfile, err := os.Open(src)
	if err != nil {
		return false, err
	}
	defer srcfile.Close()

	// create or truncate dest file
	fileInfo, _ := srcfile.Stat()
	filepath := FixPath(dir) + string(os.PathSeparator) + fileInfo.Name()
	destfile, err := OpenFileTrunc(filepath)
	if err != nil {
		return false, err
	}
	defer destfile.Close()

	// start copying
	len, err := io.Copy(destfile, srcfile)
	if err != nil || len != fileInfo.Size() {
		logger.E("copy file err:", err)
		return false, invar.ErrCopyFile
	}
	return true, nil
}

// FixPath fix path, ep:
//    /aaa/aa\\bb\\cc/d/////     -> /aaa/aa/bb/cc/d
//    E:/aaa/aa\\bb\\cc/d////e/  -> E:/aaa/aa/bb/cc/d/e
//    ""                         -> .
//    /                          -> /
func FixPath(input string) string {
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		return "."
	}

	// replace windows path separator '\' to '/'
	replaceMent := strings.Replace(input, "\\", "/", -1)
	for {
		if strings.Contains(replaceMent, "//") {
			replaceMent = strings.Replace(replaceMent, "//", "/", -1)
			continue
		}

		if replaceMent == "/" {
			return replaceMent
		}

		len := len(replaceMent)
		if len <= 0 {
			break
		}

		if replaceMent[len-1:] == "/" {
			replaceMent = replaceMent[0 : len-1]
		} else {
			break
		}
	}
	return replaceMent
}

// ReadPropFile read properties file on filesystem.
func ReadPropFile(path string) (map[string]string, error) {
	f, e := os.Open(path)
	if e == nil {
		if IsFile2(f) {
			propMap := make(map[string]string)
			reader := bufio.NewReader(f)
			for {
				line, e1 := reader.ReadString('\n')
				if e1 == nil || e1 == io.EOF {
					line = strings.TrimSpace(line)
					if len(line) != 0 && line[0] != '#' {
						// li := strings.Split(line, "=")
						eIndex := strings.Index(line, "=")
						if eIndex == -1 {
							return nil, errors.New("error parameter: '" + line + "'")
						}
						li := []string{line[0:eIndex], line[eIndex+1:]}
						if len(li) > 1 {
							k := strings.TrimSpace(li[0])
							v := strings.TrimSpace(joinLeft(li[1:]))
							propMap[k] = v
						} else {
							return nil, errors.New("error parameter: '" + li[0] + "'")
						}
					}
					if e1 == io.EOF {
						break
					}
				} else {
					// real read error.
					return nil, errors.New("error read from configuration file")
				}
			}
			return propMap, nil
		} else {
			return nil, errors.New("expect file path not directory path")
		}
	} else {
		return nil, e
	}
}

// joinLeft only for ReadPropFile()
func joinLeft(g []string) string {
	if g == nil || len(g) == 0 {
		return ""
	}
	var bf bytes.Buffer
	for i := range g {
		c := strings.Index(g[i], "#")
		if c == -1 {
			bf.WriteString(g[i])
		} else {
			bf.WriteString(g[i][0:c])
			break
		}
	}
	return string(bf.Bytes())
}

// HumanReadable format the size number of len.
func HumanReadable(len int64, during int64) string {
	if len < 1024 {
		return strconv.FormatInt(len*1000/during, 10) + "B       "
	} else if len < 1048576 {
		return strconv.FormatInt(len*1.0/1024*1000/during, 10) + "KB       "
	} else if len < 1073741824 {
		return fmt.Sprintf("%.2f", float64(len)/1048576*1000/float64(during)) + "MB       "
	} else {
		return fmt.Sprintf("%.2f", float64(len)/1073741824*1000/float64(during)) + "GB       "
	}
}
