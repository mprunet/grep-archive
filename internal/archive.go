package internal

type Archive interface {
	FullPath() string
	NextFile() (VirtualFile, error)
	Close()
}

type BaseArchive struct {
	fullPath      string
	closeCallback func()
}

func (z *BaseArchive) FullPath() string {
	return z.fullPath
}

func (z *BaseArchive) Close() {
	if z.closeCallback != nil {
		z.closeCallback()
	}
}
