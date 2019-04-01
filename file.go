package p_file

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/pefish/go-error"
	"github.com/pefish/go-slice"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

type FileClass struct {
}

var File = FileClass{}

func (this *FileClass) WriteFile(filename string, datas []byte) {
	err := ioutil.WriteFile(filename, datas, 0777)
	if err != nil {
		panic(err)
	}
}

func (this *FileClass) Exists(fileOrPath string) bool {
	_, err := os.Stat(fileOrPath)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 判断所给路径是否为文件夹
func (this *FileClass) IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func (this *FileClass) IsFile(path string) bool {
	return !this.IsDir(path)
}

func (this *FileClass) MakeDir(dir string) {
	err := os.Mkdir(dir, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func (this *FileClass) MakeFile(filename string) {
	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		panic(err)
	} else {
		_, err = f.Write([]byte(``))
		if err != nil {
			panic(err)
		}
	}
}

func (this *FileClass) AssertPathExist(path string) {
	if !this.Exists(path) {
		this.MakeDir(path)
	}
}

func (this *FileClass) ReadFile(filename string) []byte {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return bytes
}

func (this *FileClass) ReadLine(filename string, callback func(string)) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "utils: %v\n", err)
		return
	}
	input := bufio.NewScanner(f)
	for input.Scan() {
		callback(input.Text())
	}
}

func (this *FileClass) GetExt(filename string) string {
	arr := strings.Split(filename, `.`)
	return p_slice.Slice.GetLastOfSliceString(arr)
}

func (this *FileClass) ReadFileWithErr(filename string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (this *FileClass) MultipartFileToBytes(file multipart.File) []byte {
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		p_error.ThrowInternal(`MultipartFileToBytes error`)
	}
	return buf.Bytes()
}

func (this *FileClass) GetExePath() string {
	filePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(filePath)
}
