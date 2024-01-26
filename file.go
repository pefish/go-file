package go_file

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type File struct {
}

var FileInstance = File{}

func (fc *File) MustWriteFile(filename string, datas []byte) {
	err := fc.WriteFile(filename, datas)
	if err != nil {
		panic(err)
	}
}

func (fc *File) WriteFile(filename string, datas []byte) error {
	err := os.WriteFile(filename, datas, 0777)
	if err != nil {
		return err
	}
	return nil
}

func (fc *File) MustAppendCsvFile(filename string, records [][]string) {
	err := fc.AppendCsvFile(filename, records)
	if err != nil {
		panic(err)
	}
}

func (fc *File) AppendCsvFile(filename string, records [][]string) error {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	writerCsv := csv.NewWriter(f)

	err = writerCsv.WriteAll(records)
	if err != nil {
		return err
	}
	writerCsv.Flush()

	return nil
}

func (fc *File) MustAppendFile(filename string, text string) {
	err := fc.AppendFile(filename, text)
	if err != nil {
		panic(err)
	}
}

// AppendFile 附加内容到文件(不存在就创建)
func (fc *File) AppendFile(filename string, text string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err = f.WriteString(text); err != nil {
		return err
	}
	return nil
}

func (fc *File) Exists(fileOrPath string) bool {
	_, err := os.Stat(fileOrPath)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func (fc *File) IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func (fc *File) IsFile(path string) bool {
	return !fc.IsDir(path)
}

func (fc *File) MustMakeDir(dir string) {
	err := fc.MakeDir(dir)
	if err != nil {
		panic(err)
	}
}

func (fc *File) MakeDir(dir string) error {
	err := os.Mkdir(dir, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (fc *File) MustMakeFile(filename string) {
	err := fc.MakeFile(filename)
	if err != nil {
		panic(err)
	}
}

func (fc *File) MakeFile(filename string) error {
	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		return err
	}
	_, err = f.Write([]byte(``))
	if err != nil {
		return err
	}
	return nil
}

func (fc *File) AssertPathExist(path string) error {
	if fc.Exists(path) {
		return nil
	}
	return fc.MakeDir(path)
}

func (fc *File) MustReadFile(filename string) []byte {
	b, err := fc.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return b
}

func (fc *File) ReadFile(filename string) ([]byte, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (fc *File) ReadLine(
	filename string,
	callback func(text string, shouldProcess bool) error,
) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	input := bufio.NewScanner(f)
	for input.Scan() {
		err = callback(input.Text(), true)
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

func (fc *File) WriteLines(filename string, lines []string) error {
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

func (fc *File) GetExt(filename string) string {
	arr := strings.Split(filename, `.`)
	return arr[len(arr)-1]
}

func (fc *File) GetExePath() (string, error) {
	filePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Dir(filePath), nil
}
