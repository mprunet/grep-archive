package internal

import (
	"archive/zip"
	"io"
	"os"
)

type FileInAZip struct {
	parentPath string
	tmpFile    *os.File
	zip        *zip.File
}

func (z *FileInAZip) FileName() string {
	return z.zip.Name
}

func (z *FileInAZip) ParentPath() string {
	return z.parentPath
}
func (z *FileInAZip) FullPath() string {
	return z.parentPath + ":" + z.zip.Name
}

func (z *FileInAZip) CloseTemp() {
	if z.tmpFile != nil {
		var tmpFileName = z.tmpFile.Name()
		err := z.tmpFile.Close()
		if err != nil {
			PrintErrorMessageCascade(err, "Impossible to close temporary file %v mirror of %v", tmpFileName, z.FullPath())
		}
		err = os.Remove(tmpFileName)
		if err != nil {
			PrintErrorMessageCascade(err, "Impossible to remove temporary file %v mirror of %v", tmpFileName, z.FullPath())
		}
		z.tmpFile = nil
	}
}
func (z *FileInAZip) IsRegularFile() bool {
	return !z.zip.FileInfo().IsDir()
}

func (z *FileInAZip) IsInArchive() bool {
	return true
}

func (z *FileInAZip) Open() (io.ReadCloser, error) {
	return z.zip.Open()
}

func (z *FileInAZip) OpenReaderAt() (io.ReaderAt, int64, error) {
	return toReaderAt(z)
}

func (z *FileInAZip) Size() int64 {
	return int64(z.zip.UncompressedSize64)
}
func (z *FileInAZip) SetTempFile(file *os.File) {
	z.tmpFile = file
}
