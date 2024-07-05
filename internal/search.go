package internal

import (
	"bufio"
	"compress/bzip2"
	"compress/gzip"
	"github.com/klauspost/compress/zstd"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var ExcludeFile arrayFlags
var FileInMemory int64 = 10000000
var Verbose = false
var DontOutputLineFound = false
var DontPrintMatchingFileName = false
var PrintNonMatchingFileName = false
var Root = ""
var nbFile = 0
var nbFileInArchive = 0
var nbFileMatch = 0
var nbLineMatch = 0
var FileNameRegexpFilter *regexp.Regexp = nil
var FileNameGlob *string = nil
var Grep *regexp.Regexp = nil

func IsArchive(file VirtualFile) bool {
	filename := file.FileName()
	return file.IsRegularFile() && (isTarArchive(filename) || isZipArchive(filename))
}

func analyse(file VirtualFile) {
	if isExcluded(file) {
		return
	}
	checkFile(file)
	if IsArchive(file) {
		child, err := toArchive(file)
		if err != nil {
			PrintError(err)
		} else {
			findInArchive(child)
		}
	}

}

func findInArchive(archive Archive) {
	PrintVerbose("Uncompress %v", archive.FullPath)
	var archivePath = archive.FullPath()
	for true {
		var file, err = archive.NextFile()
		if err != nil {
			PrintErrorMessageCascade(err, "[1009] - Impossible to read archive content file %v", archivePath)
			break
		}
		if file == nil {
			break
		}
		analyse(file)
	}
	defer archive.Close()
}

func walker(path string, info os.DirEntry, err error) error {
	if err == nil {
		analyse(NewReadFile(path, info.Name()))
	}
	return nil
}

func DoIt() {
	_ = filepath.WalkDir(Root, walker)
	PrintInfo("%v scanned file, %v scanned in archive, %v files match file pattern, %v line match line pattern", nbFile, nbFileInArchive, nbFileMatch, nbLineMatch)
}

func GlobCheck(glob string, filename string) bool {
	if !strings.Contains(glob, "/") {
		idx := strings.LastIndex(filename, "/")
		if idx != -1 {
			filename = filename[idx+1:]
		}
	}
	ret, _ := filepath.Match(glob, filename)
	return ret
}

func checkFile(file VirtualFile) {
	if (FileNameRegexpFilter == nil && FileNameGlob == nil) ||
		(FileNameRegexpFilter != nil && FileNameRegexpFilter.MatchString(file.FileName())) ||
		(FileNameGlob != nil && GlobCheck(*FileNameGlob, file.FileName())) {
		checkFileImpl(file)
	} else if Verbose {
		PrintVerbose("File %v (%v) does not match file pattern", file.FileName(), file.FullPath())
	}
	if file.IsInArchive() {
		nbFileInArchive++
	} else {
		nbFile++
	}

}

func checkFileImpl(file VirtualFile) {
	if Grep == nil {
		PrintOut("File %v Found", file.FullPath())
		nbFileMatch++
	} else {
		if !IsArchive(file) {
			reader, err := file.Open()
			if err != nil {
				PrintErrorMessageCascade(err, "[1016] Impossible to open file %v", file.FullPath())
			} else {
				var found = false
				fullpath := file.FullPath()
				if strings.HasSuffix(fullpath, ".gz") {
					var reader2 *gzip.Reader
					reader2, err = gzip.NewReader(reader)
					if err == nil {
						reader = reader2
					} else if Verbose {
						PrintVerbose("Impossible to un compress %v, file will be read uncompressed", file.FullPath())
					}
				} else if strings.HasSuffix(fullpath, ".bz2") {
					var reader2 io.Reader
					reader2 = bzip2.NewReader(reader)
					reader = CloseableReader{
						Reader:    reader2,
						closeable: reader,
					}
				} else if strings.HasSuffix(fullpath, ".zst") {
					var reader2 io.Reader
					reader2, err = zstd.NewReader(reader)
					if err == nil {
						reader = CloseableReader{
							Reader:    reader2,
							closeable: reader,
						}
					} else if Verbose {
						PrintVerbose("Impossible to un compress %v, file will be read uncompressed", file.FullPath())
					}
				}
				defer closeReader(file, reader)
				scanner := bufio.NewScanner(reader)
				for scanner.Scan() {
					text := scanner.Text()
					if Grep.MatchString(text) {
						if !found {
							if !DontPrintMatchingFileName || Verbose {
								PrintOut("MATCH FOUND IN %v", file.FullPath())
							}
							found = true
							nbFileMatch++
						}
						if !DontOutputLineFound || Verbose {
							PrintOut("    %v", text)
						}
						nbLineMatch++
					}
				}
				if (PrintNonMatchingFileName || Verbose) && !found {
					PrintOut("MATCH NOT FOUND IN %v", file.FullPath())
					found = true
				}
			}
		} else {
			if Verbose {
				PrintVerbose("Grep of file %v skipped, will expanded before grep", file.FullPath())
			}
		}
	}
}

func isExcluded(file VirtualFile) bool {
	for _, s := range ExcludeFile {
		if GlobCheck(s, file.FileName()) {
			if Verbose {
				PrintVerbose("File %v excluded", file.FullPath())
			}
			return true
		}
	}
	return false
}
