package main

import (
	"flag"
	"github.com/mprunet/grep-archive/internal"
	"path/filepath"
	"regexp"
)

func main() {
	//var flags = flag.NewFlagSet(name, flag.ExitOnError)
	flag.Int64Var(&internal.FileInMemory, "buffer", 10000000, "Maximum Size of an internal archive before dump to disk for analyse")
	flag.StringVar(&internal.Root, "scan", "", "Directory or archive to scan")
	flag.BoolVar(&internal.Verbose, "verbose", false, "Increase verbosity")
	flag.Var(&internal.ExcludeFile, "exclude", "Exclude file from grep by filename (glob format)")
	flag.BoolVar(&internal.DontOutputLineFound, "noout", false, "Don't print matching lines")
	flag.BoolVar(&internal.DontPrintMatchingFileName, "D", false, "Don't print matching file name")
	flag.BoolVar(&internal.PrintNonMatchingFileName, "L", false, "Print non matching file name")
	var filePatternRegexp = flag.String("filter-regexp", "", "Filter file by name (regexp)")
	var filePattern = flag.String("filter", "", "Filter file by name (glob format)")
	var grep = flag.String("grep", "", "Find file content (regexp)")
	var grepString = flag.String("grep-string", "", "Find file content")
	flag.Parse()
	if internal.Root == "" {
		internal.PrintErrorMessage("scan parameter is mandatory")
		flag.Usage()
		return
	}
	if *filePatternRegexp != "" && *internal.FileNameGlob != "" {
		internal.PrintErrorMessage("filter and filter-regexp are exclusive parameters")
		flag.Usage()
		return
	}
	if *filePattern != "" {
		_, err := filepath.Match(*filePattern, "FAKE")
		if err != nil {
			internal.PrintErrorMessage("Glob %v invalid %v", *filePattern, err)
			return
		}
		internal.FileNameGlob = filePattern
	}
	if *filePatternRegexp != "" {
		re, err := regexp.Compile(*filePatternRegexp)
		if err != nil {
			internal.PrintErrorMessage("Regexp %v, cannot be compiled %v", *filePatternRegexp, err)
			return
		}
		internal.FileNameRegexpFilter = re
	}
	if *grep != "" && *grepString != "" {
		internal.PrintErrorMessage("filter and filter-string are exclusive parameters")
		flag.Usage()
		return
	}
	if *grepString != "" {
		*grep = regexp.QuoteMeta(*grepString)

	}
	if *grep != "" {
		re, err := regexp.Compile(*grep)
		if err != nil {
			internal.PrintErrorMessage("Regexp %v, cannot be compiled %v", *grep, err)
			return
		}
		internal.Grep = re
	}
	internal.DoIt()
}
