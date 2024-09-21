package go_file

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func WriteFile(filename string, datas []byte) error {
	err := os.WriteFile(filename, datas, 0777)
	if err != nil {
		return err
	}
	return nil
}

func AppendCsvFile(filename string, records [][]string) error {
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

// AppendFile 附加内容到文件(不存在就创建)
func AppendFile(filename string, text string) error {
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

func Exists(fileOrPath string) (bool, error) {
	_, err := os.Stat(fileOrPath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func IsFile(path string) bool {
	return !IsDir(path)
}

// 如果文件夹已经存在，则返回 nil
func MakeDir(dir string) error {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func MakeFile(filename string) error {
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

func AssertPathExist(path string) error {
	return MakeDir(path)
}

func MustReadFile(filename string) []byte {
	b, err := ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return b
}

func ReadFile(filename string) ([]byte, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func ReadJsonFile(v any, filename string) error {
	b, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, v)
	if err != nil {
		return err
	}
	return nil
}

func ReadLine(
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

func WriteLines(filename string, lines []string) error {
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

func GetExt(filename string) string {
	arr := strings.Split(filename, `.`)
	return arr[len(arr)-1]
}

func GetExePath() (string, error) {
	filePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Dir(filePath), nil
}
