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
	"github.com/youhei-yp/wing/logger"
	"github.com/youhei-yp/wing/secure"
	"io"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

const BufferSize = 1024 * 30

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

// GetFileMd5 get file's md5 string
func GetFileMd5(fi string) (string, error) {
	md := md5.New()
	f, e := GetFile(fi)
	if e != nil {
		return "", e
	} else {
		defer f.Close()
		_, e1 := io.Copy(md, f)
		if e1 != nil {
			return "", e1
		}
		md5 := hex.EncodeToString(md.Sum(nil))
		return md5, nil
	}
}

// GetFile return a File for read.
func GetFile(path string) (*os.File, error) {
	return os.Open(path)
}

// CopyFile Copy file.
func CopyFile(src string, dest string) (s bool, e error) {
	srcfile, err1 := os.Open(src)
	if err1 == nil {
		// create or truncate dest file
		destfile, err2 := os.OpenFile(dest, syscall.O_CREAT|os.O_WRONLY|syscall.O_TRUNC, 0660)
		// ensure close files finally
		defer func() {
			// log.Print("close src and dest files")
			srcfile.Close()
			destfile.Close()
		}()
		// if "create or truncate dest file" succeed then start copying
		if err2 == nil {
			Try(func() {
				bs := make([]byte, BufferSize)
				for {
					len, e1 := srcfile.Read(bs)
					if e1 == io.EOF {
						break
					}
					destfile.Write(bs[0:len])
				}
				s = true
			}, func(i interface{}) {
				s = false
				e = errors.New(fmt.Sprint(i))
			})
		} else {
			s = false
			e = errors.New(fmt.Sprint(err2))
		}
	} else {
		s = false
		e = errors.New("open source file failed")
	}

	return s, e
}

// CopyFileTo Copy file to dir.
func CopyFileTo(src string, dir string) (s bool, e error) {
	srcfile, err1 := os.Open(src)
	if err1 == nil {
		// create or truncate dest file
		fileInfo, _ := srcfile.Stat()
		destfile, err2 := os.OpenFile(FixPath(dir)+string(os.PathSeparator)+fileInfo.Name(), syscall.O_CREAT|os.O_WRONLY|syscall.O_TRUNC, 0660)
		// ensure close files finally
		defer func() {
			srcfile.Close()
			destfile.Close()
		}()
		// if "create or truncate dest file" succeed then start copying
		if err2 == nil {
			len, e1 := io.Copy(destfile, srcfile)
			if len != fileInfo.Size() || e1 != nil {
				s = false
				e = e1
			} else {
				s = true
			}
		} else {
			s = false
			e = errors.New(fmt.Sprint(err2))
		}
	} else {
		s = false
		e = errors.New("open source file failed")
	}

	return s, e
}

// Exists check whether the file exists.
func Exists(path string) bool {
	fi, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	if fi == nil {
		return false
	}
	return true
}

// Delete delete a file or directory.
// if delete failed, you should find out by yourself where and what is the problem.
// special: if the file is not exists, this will return true.
func Delete(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return true
	}
	err = os.Remove(path)
	return nil == err
}

// DeleteAll delete file or directory.
// if delete failed, you should find out by yourself where and what is the problem.
// if is a directory, it will try to delete all files below.
// special: if the file is not exists, this will return true.
func DeleteAll(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return true
	}
	err = os.RemoveAll(path)
	return nil == err
}

// CreateFile create new file.
func CreateFile(path string) (*os.File, error) {
	fi, err := os.Create(path)
	return fi, err
}

// OpenFile4Write create new file.
func OpenFile4Write(path string) (*os.File, error) {
	fi, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	return fi, err
}

// CreateDir create new directory.
func CreateDir(path string) error {
	return os.Mkdir(path, 0777)
}

// CreateAllDir create new directory.
func CreateAllDir(path string) error {
	return os.MkdirAll(path, 0777)
}

// IsFile check whether path point to a file.
func IsFile(path string) bool {
	fi, err := os.Stat(path)
	if nil == err {
		return !fi.IsDir()
	} else {
		return false
	}
}

// IsFile1 check whether path point to a file.
func IsFile1(file *os.File) bool {
	fi, err := file.Stat()
	if nil == err {
		return !fi.IsDir()
	} else {
		return false
	}
}

// IsDir check whether path point to a file.
func IsDir(path string) bool {
	fi, err := os.Stat(path)
	if nil == err {
		return fi.IsDir()
	} else {
		return false
	}
}

// IsDir1 check whether path point to a file.
func IsDir1(file *os.File) bool {
	fi, err := file.Stat()
	if nil == err {
		return fi.IsDir()
	} else {
		return false
	}
}

// MoveFile rename file / move file
func MoveFile(src string, dest string) error {
	return os.Rename(src, dest)
}

// ChangeWorkDir change work directory to path.
func ChangeWorkDir(path string) error {
	return os.Chdir(path)
}

// GetWorkDir get work directory.
func GetWorkDir() (string, error) {
	return os.Getwd()
}

// GetTempDir return a temp directory.
func GetTempDir() string {
	return os.TempDir()
}

// IsAbsPath tell if the path is absolute.
func IsAbsPath(path string) bool {
	return filepath.IsAbs(path)
}

// GetFileExt get file extension
func GetFileExt(filePath string) string {
	return path.Ext(filePath)
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
		if IsFile1(f) {
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

// IsFileByPath check whether path point to a file.
func IsFileByPath(path string) bool {
	fi, err := os.Stat(path)
	if nil == err {
		return !fi.IsDir()
	} else {
		return false
	}
}

// IsFileByFile check whether path point to a file.
func IsFileByFile(file *os.File) bool {
	fi, err := file.Stat()
	if nil == err {
		return !fi.IsDir()
	} else {
		return false
	}
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
