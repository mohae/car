package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mohae/appname"
	"github.com/mohae/contour"
	log "github.com/mohae/ezlog"
)

const (
	Format = "format"
)

func init() {
	contour.RegisterBoolFlag("delete_files", "d", false, "false", "remove files after adding them to the archive")
	contour.RegisterBoolFlag("abs_path", "a", false, "false", "use absolute path instead of relative")
	contour.RegisterStringFlag(Format, "f", "tar", "tar", "archive format: tar or zip")
	contour.RegisterStringFlag("compression", "c", "gzip", "gzip", "compression format to use; does not apply to zip archives")
	contour.RegisterStringFlag("output_dir", "o", "", "", "output directory: if empty, wd will be the output location")
	contour.RegisterStringFlag("name", "n", "", "", "name for archive: if empty, either the filename or dir name will be used as the base name")
	contour.RegisterStringFlag("exclude", "", "", "", "exclude files, given as a PATTERN")
	contour.RegisterStringFlag("exclude_ext", "e", "", "", "exclude files with EXTENSIONS")
	contour.RegisterStringFlag("exclude_anchored", "", "", "", "exclude patterns match file name start")
	contour.RegisterStringFlag("include", "", "", "", "include files, given as a PATTERN")
	contour.RegisterStringFlag("include_ext", "i", "", "", "include files with EXTENSIONS")
	contour.RegisterStringFlag("include_anchored", "", "", "", "include patterns match file name start")
	// Extract Operation Modifiers
	//	contour.RegisterBoolFlag("keep-old-files", "k", false, "don't replace existing files when extracting")
	//	contour.RegisterBoolFlag("keep-newer-files", "", false, "don't replace existing files that are newer than their archive copies")
	//	contour.RegisterBoolFlag("overwrite", "", true, "overwrite existing files when extracting")
	////	contour.RegisterBoolFlag("atime-preserve", "", true)
	//	contour.RegisterBoolFlag("modification-time", "m", false, "don't extract file modified time")
	//	contour.RegisterBoolFlag("same-owner", "", true, "try extracting files with the same ownership")
	//	contour.RegisterBoolFlag("no-same-owner", "", false, "extract files as yourself")
	//	contour.RegisterBoolFlag("numeric-owner", "", true, "always use numbers for user/group names")
	//	contour.RegisterBoolFlag("same-permissions", "p", true, "extract permissions information")
	//	contour.RegisterBoolFlag("no-same-permissions", "", false, "do not extract permissions information")

	//	contour.RegisterBoolFlag("wildcards", "", false, "false", "patterns use wildcards")
	//	contour.RegisterBoolFlag("no-wildcards", "", true, "true", "patters do not use wildcards")
	//	contour.RegisterTimeFlag("newer", "N", unsetTime, "only store files newer than DATE or File")
	//contour.RegisterTimeFlag("newer-mtime", "M", unsetTime, "not set", "only store files modified since DATE")
	//	contour.RegisterStringFlag("newer-file", "", "only store files newere than the DATE for FILENAMEE")

	// Register option aliases
	//	contour.RegisterFlagAlias("newer-date", "after-date")
	//	contour.RegisterFlagAlias("same-permissions", "preserve-permissions")
}
// usage is the usage func for flag.Usage.
func usage() {
	fmt.Fprint(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "  %s [FLAGS] \n", appname.Get())
	fmt.Fprint(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "Insert information about %s here\n", appname.Get())
	fmt.Fprint(os.Stderr, "\n")
	fmt.Fprint(os.Stderr, "Options:\n")
	flag.PrintDefaults()
}

func SetLogging() (closeLog bool) {
	lvl, ok := log.LevelByName(contour.String("level"))
	if !ok {
		fmt.Fprintf(os.Stderr, "%s: unknown log level", contour.String("level"))
		os.Exit(1)
	}
	log.SetLevel(lvl)
	l := contour.String("logfile")
	switch l {
	case "stderr":
		return false
	case "stdout":
		log.SetOutput(os.Stdout)
		return false
	case "":
		log.SetOutput(ioutil.Discard)
		return false
	default:
		logF, err := os.OpenFile(l, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0664)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: open log file: %s", appname.Get(), err)
			os.Exit(1)
		}
		log.SetOutput(logF)
		return true
	}
}

func carMain() int {

	err := contour.Set()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s", appname.Get(), err)
		return 1
	}

	args, err := contour.ParseFlags()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return 1
	}

	if len(args) <= 1 {
		fmt.Fprintf(os.Stderr, "%s: at least an operation and a filename must be provided", appname.Get())
		return 1
	}

	closeLog := SetLogging()

	if closeLog {
		defer logF.Close() // make sure the logfile is closed if there is one
	}
	var msg string
	switch args[0] {
	case "create":
		if len(args) <= 2 {
			log.Error("create: at least a destination filename and a source filename must be provided")
		}
	 	msg, err = Create(args[1] , args[2:]...)
	case "extract":
		msg, err = Extract(args[1])
	default:
		log.Errorf("%s: unknown operation", args[0])
		return 1
	}

	if err != nil {
		log.Error(err)
		return 1
	}
	fmt.Println(msg)

	return 0
}
