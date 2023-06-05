package internal

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type VirtualFile interface {
	FullPath() string
	ParentPath() string
	FileName() string
	Open() (io.ReadCloser, error)
	OpenReaderAt() (io.ReaderAt, int64, error)
	Size() int64
	SetTempFile(file *os.File)
	IsInArchive() bool
	IsRegularFile() bool
	CloseTemp()
}

func isZipArchive(filename string) bool {
	if strings.HasSuffix(filename, ".zip") ||
		strings.HasSuffix(filename, ".apk") ||
		strings.HasSuffix(filename, ".war") ||
		strings.HasSuffix(filename, ".ear") ||
		strings.HasSuffix(filename, ".jar") {
		return true
	}
	return false
}

func isTarArchive(filename string) bool {
	if strings.HasSuffix(filename, ".tgz") ||
		strings.HasSuffix(filename, ".tar.gz") ||
		strings.HasSuffix(filename, ".tar") {
		return true
	}
	return false
}

func toArchive(file VirtualFile) (Archive, error) {
	var filename = file.FileName()
	if isZipArchive(filename) {
		return NewZipArchive(file)
	} else if isTarArchive(filename) {
		return NewTarArchive(file)
	} else {
		return nil, nil
	}
}

func toReaderAt(z VirtualFile) (io.ReaderAt, int64, error) {
	var archivePath = z.FullPath()
	var readerAt io.ReaderAt
	var filezize int64
	reader, err := z.Open()
	if err != nil {
		return nil, 0, CascadeError(err, "[1001] - Impossible to open file %v", archivePath)
	}
	defer closeReader(z, reader)
	if z.Size() > FileInMemory { // Decompression en memoire
		content, err2 := ioutil.ReadAll(reader)
		if err2 != nil {
			return nil, 0, CascadeError(err2, "[1004] - Impossible to buffer file to memory %v", archivePath)
		}
		filezize = int64(len(content))
		readerAt = bytes.NewReader(content)
	} else {
		tmpFile, err2 := os.CreateTemp("", "uncompress.*.zip")
		if err2 != nil {
			return nil, 0, CascadeError(err2, "[1005] - Impossible to create a temporary file to extract %v", archivePath)
		}
		tmpFileName := tmpFile.Name()
		z.SetTempFile(tmpFile)
		filezize, err = io.Copy(tmpFile, reader)
		if err != nil {
			return nil, 0, CascadeError(err2, "[1006] - Impossible to copy the content of %v to a temporary file %v", archivePath, tmpFileName)
		}
		_, err = tmpFile.Seek(0, 0)
		if err != nil {
			return nil, 0, CascadeError(err2, "[1009] - Impossible to rewind to the beginning of the temporary file %v containing a copy of %v", tmpFileName, archivePath)
		}
		readerAt = tmpFile
	}

	return readerAt, filezize, nil
}
