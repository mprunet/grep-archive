package internal

import (
	"archive/tar"
	"io"
	"os"
)

type FileInATar struct {
	parentPath  string
	tmpFile     *os.File
	header      *tar.Header
	entryReader *tar.Reader
}

type TarReadCloserGuard struct {
	reader *tar.Reader
}

func (t *TarReadCloserGuard) Read(p []byte) (n int, err error) {
	return t.reader.Read(p)
}
func (t *TarReadCloserGuard) Close() (err error) {
	return nil
}

func (t *FileInATar) FileName() string {
	return t.header.Name

}
func (z *FileInATar) ParentPath() string {
	return z.parentPath
}

func (z *FileInATar) FullPath() string {
	return z.parentPath + ":" + z.header.Name
}

func (z *FileInATar) IsInArchive() bool {
	return true
}
func (z *FileInATar) IsRegularFile() bool {
	return true // Seul les fichiers sont retournes
}
func (t *FileInATar) Open() (io.ReadCloser, error) {
	var v = new(TarReadCloserGuard)
	v.reader = t.entryReader
	return v, nil
}
func (z *FileInATar) OpenReaderAt() (io.ReaderAt, int64, error) {
	return toReaderAt(z)
}
func (t *FileInATar) Size() int64 {
	return t.header.Size
}
func (t *FileInATar) SetTempFile(file *os.File) {
	t.tmpFile = file
}
func (t *FileInATar) CloseTemp() { // Ne pas fermer pour les tar
	if t.tmpFile != nil {
		var tmpFileName = t.tmpFile.Name()
		err := t.tmpFile.Close()
		if err != nil {
			PrintErrorMessageCascade(err, "[1010] Impossible to close temporary file %v mirror of %v", tmpFileName, t.FullPath())
		}
		err = os.Remove(tmpFileName)
		if err != nil {
			PrintErrorMessageCascade(err, "[1011] Impossible to remove temporary file %v mirror of %v", tmpFileName, t.FullPath())
		}
		t.tmpFile = nil
	}
}
