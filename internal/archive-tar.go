package internal

import (
	"archive/tar"
	"compress/bzip2"
	"compress/gzip"
	"io"
	"strings"
)

type TarArchive struct {
	entryReader *tar.Reader
	BaseArchive
}

func (t *TarArchive) FullPath() string {
	return t.fullPath
}

func (t *TarArchive) NextFile() (VirtualFile, error) {
	for true {
		header, err := t.entryReader.Next()
		if err == io.EOF {
			return nil, nil
		} else if err != nil {
			return nil, err
		}
		if header.Typeflag == tar.TypeReg {
			var ret FileInATar
			ret.parentPath = t.fullPath
			ret.header = header
			ret.entryReader = t.entryReader
			return &ret, nil
		}
	}
	return nil, nil

}

func NewTarArchive(z VirtualFile) (*TarArchive, error) {
	var ret = new(TarArchive)
	fullpath := z.FullPath()
	stream, err := z.Open()
	if err != nil {
		return nil, CascadeError(err, "[1003] - Impossible to read file %v", fullpath)
	}
	var reader io.Reader = stream
	ret.closeCallback = func() {
		err := stream.Close()
		if err != nil {
			PrintErrorMessageCascade(err, "[1012] Impossible to close reader %v", fullpath)
		}
	}
	ret.fullPath = fullpath
	if strings.HasSuffix(fullpath, "gz") {
		var err error
		reader, err = gzip.NewReader(reader)
		if err != nil {
			ret.Close()
			return nil, CascadeError(err, "[1007] - Impossible to uncompress gzip file %v", fullpath)
		}
	} else if strings.HasSuffix(fullpath, "bz2") {
		reader = bzip2.NewReader(reader)
	}
	ret.entryReader = tar.NewReader(reader)
	return ret, nil
}
