package go_file

import (
	"testing"
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
