package files

import (
	"os"
	"testing"
)

func TestIsExists(t *testing.T) {
	tmpFileName := "tmp.d"

	os.Create(tmpFileName)

	if !IsExists(tmpFileName) {
		t.Fail()
	}

	os.Remove(tmpFileName)

	if IsExists(tmpFileName) {
		t.Fail()
	}
}

func TestIsDir(t *testing.T) {
	tmpDirName := "tmp.d"

	t.Cleanup(func() {
		os.Remove(tmpDirName)
	})

	os.Mkdir(tmpDirName, 0644)

	if !IsDir(tmpDirName) {
		t.Fail()
	}

	os.RemoveAll(tmpDirName)
	os.Create(tmpDirName)

	if IsDir(tmpDirName) {
		t.Fail()
	}
}

func TestIsFile(t *testing.T) {
	tmpFileName := "tmp.d"

	t.Cleanup(func() {
		os.RemoveAll(tmpFileName)
	})

	os.Create(tmpFileName)

	if !IsFile(tmpFileName) {
		t.Fail()
	}

	os.Remove(tmpFileName)
	os.Mkdir(tmpFileName, 0644)

	if IsFile(tmpFileName) {
		t.Fail()
	}
}

func TestGetFile(t *testing.T) {
	tmpFileName := "tmp.d"

	t.Cleanup(func() {
		os.Remove(tmpFileName)
		os.RemoveAll(tmpFileName)
	})

	// Case #1
	os.Create(tmpFileName)
	GetFile(tmpFileName)

	// Case #2
	os.Remove(tmpFileName)
	GetFile(tmpFileName)

	// Case #3
	os.Remove(tmpFileName)
	os.Mkdir(tmpFileName, 0644)
	GetFile(tmpFileName)

	/*
		if it exits on case #3 with an error that goes like:
			file <file> is a directory
		then it is working properly
	*/
}

func TestWriteFile(t *testing.T) {
	tmpFileName := "tmp"
	tmpFileData := "tmp.d"

	t.Cleanup(func() {
		os.Remove(tmpFileName)
	})

	WriteFile(tmpFileName, []byte(tmpFileData))

	if data, _ := os.ReadFile(tmpFileName); string(data) != string(tmpFileData) {
		t.Fail()
	}
}