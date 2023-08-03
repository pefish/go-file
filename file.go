package go_file

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/pefish/go-error"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

type FileClass struct {
}

var FileUtilInstance = FileClass{}

func (fc *FileClass) WriteFile(filename string, datas []byte) {
	err := ioutil.WriteFile(filename, datas, 0777)
	if err != nil {
		panic(err)
	}
}

// AppendFile 附加内容到文件(不存在就创建)
func (fc *FileClass) AppendFile(filename string, text string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(text); err != nil {
		panic(err)
	}
}

func (fc *FileClass) Exists(fileOrPath string) bool {
	_, err := os.Stat(fileOrPath)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func (fc *FileClass) IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func (fc *FileClass) IsFile(path string) bool {
	return !fc.IsDir(path)
}

func (fc *FileClass) MakeDir(dir string) {
	err := os.Mkdir(dir, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func (fc *FileClass) MakeFile(filename string) {
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

func (fc *FileClass) AssertPathExist(path string) {
	if !fc.Exists(path) {
		fc.MakeDir(path)
	}
}

func (fc *FileClass) ReadFile(filename string) []byte {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return bytes
}

func (fc *FileClass) ReadLine(filename string, callback func(text string, shouldProcess bool) error) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	input := bufio.NewScanner(f)
	for input.Scan() {
		err := callback(input.Text(), true)
		if err != nil {
			return err
		}
	}
	err = callback("", false)
	if err != nil {
		return err
	}
	return nil
}

func (fc *FileClass) WriteLines(filename string, lines []string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func (fc *FileClass) GetExt(filename string) string {
	arr := strings.Split(filename, `.`)
	return arr[len(arr)-1]
}

func (fc *FileClass) ReadFileWithErr(filename string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (fc *FileClass) MultipartFileToBytes(file multipart.File) []byte {
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		go_error.ThrowInternal(fmt.Errorf(`MultipartFileToBytes error`))
	}
	return buf.Bytes()
}

func (fc *FileClass) GetExePath() string {
	filePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(filePath)
}
