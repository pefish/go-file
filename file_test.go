package go_file

import (
	"testing"
)

func TestFileClass_AppendFile(t *testing.T) {
	//FileInstance.AppendFile(`1.txt`, "haha\n")
}

func TestFile_MustWriteCsvFile(t *testing.T) {
	FileInstance.WriteCsvFile("1.csv", [][]string{
		[]string{
			"123",
			"gsfdga",
			"gsdfgw",
		},
	})
}
