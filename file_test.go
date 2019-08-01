package go_file

import (
	"testing"
)

func TestFileClass_AppendFile(t *testing.T) {
	File.AppendFile(`1.txt`, "haha\n")
}
