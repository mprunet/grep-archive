package internal

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var FileInMemory int64 = 10000000
var Verbose = false
var PrintMatchingFileName = false
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
	return file.IsRegularFile() && (strings.HasSuffix(filename, ".zip") ||
		strings.HasSuffix(filename, ".war") ||
		strings.HasSuffix(filename, ".ear") ||
		strings.HasSuffix(filename, ".jar") ||
		strings.HasSuffix(filename, ".tgz") ||
		strings.HasSuffix(filename, ".tar.gz") ||
		strings.HasSuffix(filename, ".tar"))
}

func analyse(file VirtualFile) {
	filename := file.FileName()
	checkFile(file)
	if file.IsRegularFile() && (strings.HasSuffix(filename, ".zip") ||
		strings.HasSuffix(filename, ".war") ||
		strings.HasSuffix(filename, ".ear") ||
		strings.HasSuffix(filename, ".jar") ||
		strings.HasSuffix(filename, ".tgz") ||
		strings.HasSuffix(filename, ".tar.gz") ||
		strings.HasSuffix(filename, ".tar")) {
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

func globCheck(filename string) bool {
	ret, _ := filepath.Match(*FileNameGlob, filename)
	return ret
}

func checkFile(file VirtualFile) {
	if (FileNameRegexpFilter == nil && FileNameGlob == nil) ||
		(FileNameRegexpFilter != nil && FileNameRegexpFilter.MatchString(file.FileName())) ||
		(FileNameGlob != nil && globCheck(file.FileName())) {
		if Grep == nil {
			PrintOut("File %v Found", file.FullPath())
			nbFileMatch++
		} else {
			if !IsArchive(file) {
				reader, err := file.Open()
				if err != nil {
					PrintErrorMessageCascade(err, "Impossible to open file %v", file.FullPath())
				} else {
					var found = false
					defer closeReader(file, reader)
					scanner := bufio.NewScanner(reader)
					for scanner.Scan() {
						text := scanner.Text()
						if Grep.MatchString(text) {
							if !found {
								PrintOut("MATCH FOUND IN %v", file.FullPath())
								found = true
								nbFileMatch++
							}
							PrintOut("    %v", text)
							nbLineMatch++
						}
					}
					if (PrintMatchingFileName || Verbose) && !found {
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
	} else if Verbose {
		PrintVerbose("File %v does not match file pattern", file.FullPath())
	}
	if file.IsInArchive() {
		nbFileInArchive++
	} else {
		nbFile++
	}

}
