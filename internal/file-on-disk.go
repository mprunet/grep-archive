package internal

import (
	"io"
	"os"
)

type RealFile struct {
	fullPath   string
	parentPath string
	filename   string
	file       *os.File
	info       os.FileInfo
}

func NewReadFile(fullPath string, filename string) *RealFile {
	var ret = new(RealFile)
	ret.fullPath = fullPath
	ret.parentPath = fullPath[0 : len(fullPath)-len(filename)]
	ret.filename = filename
	return ret
}

func (r *RealFile) FullPath() string {
	return r.fullPath
}

func (r *RealFile) ParentPath() string {
	return r.parentPath
}

func (r *RealFile) FileName() string {
	return r.filename
}

func (r *RealFile) Open() (io.ReadCloser, error) {
	var err error
	r.file, err = os.Open(r.FullPath())
	return r.file, err
}
func (r *RealFile) OpenReaderAt() (io.ReaderAt, int64, error) {
	var err error
	r.file, err = os.Open(r.FullPath())
	if err != nil {
		return nil, 0, CascadeError(err, "[1000] - Impossible to read file %v", r.FullPath())
	}
	return r.file, r.Size(), nil
}
func (r *RealFile) FileInfo() (os.FileInfo, error) {
	var err error = nil
	if r.info == nil {
		r.info, err = os.Lstat(r.FullPath())
	}
	return r.info, err
}

func (r *RealFile) Size() int64 {
	fi, err := r.FileInfo()
	var ret int64 = 0
	if err == nil {
		ret = fi.Size()
	}
	return ret
}
func (r *RealFile) SetTempFile(*os.File) {
	panic("Operation must never be called")
}
func (r *RealFile) IsInArchive() bool {
	return false
}
func (r *RealFile) IsRegularFile() bool {
	fi, err := r.FileInfo()
	if err != nil {
		return false
	}
	return fi.Mode().IsRegular()
}
func (r *RealFile) CloseTemp() {
	err := r.file.Close()
	if err != nil {
		PrintErrorMessageCascade(err, "[1015] Impossible to close reader %v", r.FullPath())
	}
	r.file = nil
	r.info = nil
}
