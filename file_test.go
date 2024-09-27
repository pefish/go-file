package go_file

import (
	"fmt"
	"testing"

	go_test_ "github.com/pefish/go-test"
)

func TestFileClass_AppendFile(t *testing.T) {
	//AppendFile(`1.txt`, "haha\n")
}

func TestFile_MustWriteCsvFile(t *testing.T) {
	AppendCsvFile("1.csv", [][]string{
		{
			"123",
			"gsfdga",
			"gsdfgw",
		},
	})
}

func TestWriteTempFile(t *testing.T) {
	filename, err := WriteTempFile([]byte("hahaha"))
	go_test_.Equal(t, nil, err)
	fmt.Println(filename)
}
