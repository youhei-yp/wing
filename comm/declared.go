// -----------------------
// DECLARED METHODS
// ----------------------

package comm

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"path"
	"path/filepath"
	"time"
)

// TValue ternary operation
// @declared
func TValue(condition bool, trueValue interface{}, falseValue interface{}) interface{} {
	return Ternary(condition, trueValue, falseValue)
}

// Md5sum merge Md5 string to one Md5 string.
// @declared
func Md5sum(input ...string) string {
	h := md5.New()
	if input != nil {
		for _, v := range input {
			io.WriteString(h, v)
		}
	}
	cipher := h.Sum(nil)
	return hex.EncodeToString(cipher)
}

// GetFile return a File for read.
// @declared please use os.Open(path) directly
func GetFile(path string) (*os.File, error) {
	return os.Open(path)
}

// GetFileMd5 get file's md5 string
// @declared
func GetFileMd5(file string) (string, error) {
	return FileMD5(file)
}

// Fill2Digits add zero for number > 10
// @declared
func Fill2Digits(input int) string {
	return To2Digits(input)
}

// Fill3Digits add zero for number > 10 or 100
// @declared
func Fill3Digits(input int) string {
	return To3Digits(input)
}

// GetHumanReadableDuration return readable time during start to end: 12:12:12
// @declared
func Unix2timeDuration(start time.Time, end time.Time) string {
	return DurHours(start, end)
}

// GetLongHumanReadableDuration return readable time during start to end: 2d 6h 25m 48s
// @declared
func Unix2DaysDuration(start time.Time, end time.Time) string {
	return DurDays(start, end)
}

// CreateFile create new file.
// @declared please use os.Create(path) directly
func CreateFile(path string) (*os.File, error) {
	return os.Create(path)
}

// Delete delete a file or directory.
// if delete failed, you should find out by yourself where and what is the problem.
// special: if the file is not exists, this will return true.
// @declared see DeleteFile
func Delete(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return true
	}
	err = os.Remove(path)
	return nil == err
}

// Exists check whether the file exists.
// @declared see IsExistFile
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

// OpenFile4Write create new file.
// @declared see OpenFileWrite
func OpenFile4Write(filepath string) (*os.File, error) {
	return OpenFileWrite(filepath)
}

// DeleteAll delete file or directory.
// if delete failed, you should find out by yourself where and what is the problem.
// if is a directory, it will try to delete all files below.
// special: if the file is not exists, this will return true.
// @declared see DeletePath
func DeleteAll(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return true
	}
	err = os.RemoveAll(path)
	return nil == err
}

// CreateDir create new directory.
// @declared see MakeDirs
func CreateDir(path string) error {
	return os.Mkdir(path, 0777)
}

// CreateAllDir create new directory.
// @declared see MakeDirs
func CreateAllDir(path string) error {
	return os.MkdirAll(path, 0777)
}

// IsFile1 check whether path point to a file.
// @declared see IsFile2
func IsFile1(file *os.File) bool {
	fi, err := file.Stat()
	if nil == err {
		return !fi.IsDir()
	} else {
		return false
	}
}

// IsDir1 check whether path point to a file.
// @declared see IsDir2
func IsDir1(file *os.File) bool {
	fi, err := file.Stat()
	if nil == err {
		return fi.IsDir()
	} else {
		return false
	}
}

// MoveFile rename file / move file
// @declared please use os.Rename(src, dest) directly
func MoveFile(src string, dest string) error {
	return os.Rename(src, dest)
}

// ChangeWorkDir change work directory to path.
// @declared please use os.Chdir(path) directly
func ChangeWorkDir(path string) error {
	return os.Chdir(path)
}

// GetWorkDir get work directory.
// @declared please use os.Getwd() directly
func GetWorkDir() (string, error) {
	return os.Getwd()
}

// GetTempDir return a temp directory.
// @declared please use os.TempDir() directly
func GetTempDir() string {
	return os.TempDir()
}

// IsAbsPath tell if the path is absolute.
// @declared please use filepath.IsAbs(path) directly
func IsAbsPath(path string) bool {
	return filepath.IsAbs(path)
}

// GetFileExt get file extension
// @declared please use path.Ext(filePath) directly
func GetFileExt(filePath string) string {
	return path.Ext(filePath)
}

// IsFileByPath check whether path point to a file.
// @declared see IsFile()
func IsFileByPath(path string) bool {
	fi, err := os.Stat(path)
	if nil == err {
		return !fi.IsDir()
	} else {
		return false
	}
}

// IsFileByFile check whether path point to a file.
// @declared see IsFile2
func IsFileByFile(file *os.File) bool {
	fi, err := file.Stat()
	if nil == err {
		return !fi.IsDir()
	} else {
		return false
	}
}
