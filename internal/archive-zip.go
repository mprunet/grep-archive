package internal

import (
	"archive/zip"
)

type ZipArchive struct {
	reader *zip.Reader
	file   []*zip.File
	idx    int
	BaseArchive
}

func NewZipArchive(z VirtualFile) (*ZipArchive, error) {
	var archive = new(ZipArchive)
	newarchivepath := z.FullPath()
	readerAt, size, err := z.OpenReaderAt()
	if err != nil {
		return nil, err // Pas de message specifique n'apporte rien au debug deja customise dans les classes
	}
	archive.closeCallback = func() {
		z.CloseTemp()
	}
	zipreader, err2 := zip.NewReader(readerAt, size)
	if err2 != nil {
		return nil, CascadeError(err2, "[1002] - Impossible to uncompress zip archive file %v", newarchivepath)
	}
	archive.reader = zipreader
	archive.fullPath = newarchivepath
	return archive, nil
}

func (z *ZipArchive) NextFile() (VirtualFile, error) {
	if z.file == nil {
		z.file = z.reader.File
		z.idx = 0
	}
	var curIdx = z.idx
	if z.idx < len(z.file) {
		z.idx++
		var ret FileInAZip
		ret.parentPath = z.fullPath
		ret.zip = z.file[curIdx]
		return &ret, nil
	} else {
		return nil, nil
	}
}
