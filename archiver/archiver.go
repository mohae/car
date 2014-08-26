// creates compressed tarballs
//    
// Go's `archive` package, supports `tar` and `zip`
// Go's `comopress` package supports: bzip2, flate, gzip, lzw, zlib
//
// Archiver supports zip and tar. For tar, archiver also supports
// the following compression:
// When using archiver, compression is not optional
package archiver 

import (
	"compress/gzip"
	"errors"
	"io"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/dotcloud/tar"
//	"github.com/mohae/goutils"
)

var appendDatetime = true
var datetimeFormat = "2006-01-02T150405Z0700"
var datetimePrefix = "-"

const (
	defaultCompresion = 'tgz'
)

// SetDateTimeFormat overrides the default datetime format. The passed format
// must use Go's datetime format.
func SetDatetimeFormat(s string) {
	datetimeFormat = s
}

// SetAppendDateTime sets whether the tarball names should be appended with the
// current datetime. The appended datetime will be prefixed with -, unless that
// is overridden, and use the set datetime format, which can be overridden by 
// calling SetDateTimneFormat.
func SetAppendDatetime(b bool) {
	appendDatetime = b
}

// SetDatetimePrefix overrides the default datetime prefix of `-`. The datetime
// prefix is used to to prefix the datetime prior to appending it to the
// filename, e.g. filename-datetime.tgz.
// 
// To not use a prefix, set it to an emptys string, ""
func SetDatetimePrefix(s string) {
	datetimePrefix = s
}

// Archive holds information about an archive.
type Archive struct {
	// Path to the target directory for the archive output.
	OutDir string

	// Name of the archive.
	Name string

	// Compression type to be used (extension).
	Type string

	// List of files to add to the archive.
	Fi	[]os.FileInfo
}

func (a *Archive) addFile(tw *tar.Writer, fi os.FileInfo) error {
	// It should exist, since we are 
	file, err := os.Open(fi.Name())
	if err != nil {
		return err
	}
	defer file.Close()

	var fileStat os.FileInfo
	fileStat, err = file.Stat()
	if err != nil {
		return err
	}

	// Don't add directories--they result in tar header errors.
	fileMode := fileStat.Mode()
	if fileMode.IsDir() {
		return nil
	}

	// Create the tar header stuff.
	tarHeader := new(tar.Header)
	tarHeader.Name = filename
	tarHeader.Size = fileStat.Size()
	tarHeader.Mode = int64(fileStat.Mode())
	tarHeader.ModTime = fileStat.ModTime()

	// Write the file header to the tarball.
	err = tW.WriteHeader(tarHeader)
	if err != nil {
		return err
	}

	// Add the file to the tarball.
	_, err = io.Copy(tW, file)
	if err != nil {
		return err
	}

	return nil
}

// ArchiveAndDelete creates a compressed archive of the passed sources using the
// passed filename in the destination directory. If the archive is successfully
// created and written to the target, the archived targets are deleted.
// Target(s) is variadic.
func (a *Archive) ArchiveAndDelete(compression, filename, destination string, sources ...string) error {
	if filename == "" || filename == "./" || filename == "." {
		return errors.New(fmt.Sprintf("Filename was empty or invalid: %s", filename)
	}

	if len(targets) <= 0  {
		return errors.New(fmt.Sprintf("No source files or directories were specified. Unable to create archive"))
	}
 
	if compression == "" {
		compressiond = defaultCompression
	}
	
	// See if the requested compression exists


	// See if src exists, if it doesn't then don't do anything
	_, err := os.Stat(p)
	if err != nil {

		// Nothing to do if it doesn't exist
		if os.IsNotExist(err) {
			return nil
		}

		return err
	}


	// build the tarball file name; let archive worry about whether the
	// destination can be written to	
	err = a.archive(tarball, sources)
	if err != nil {
		return err
	}

	// Delete the old artifacts.
	err = a.delete(sources)
	if err != nil {
		return err
	}

	return nil
}

func (a *Archiver) Archive(tarballName string, sources ...string) error {
	// Make sure the tarball can be written to


	// Get a list of directory contents
	err := a.DirWalk(p)
	if err != nil {
		return err
	}

	if len(a.Files) <= 1 {
		return nil
	}

	// Get the current date and time in a slightly modifie ISO 8601 format:
	// the colons are stripped from the time.
	nowF := formattedNow()

	// Get the relative path so that it can be added to the tarball name.
	relPath := path.Dir(p)
	// The tarball's name is the directory name + current time + extensions.
	tarBallName := relPath + a.Name + "-" + nowF + ".tar.gz"

	// Create the new archive file.
	tBall, terr := os.Create(tarBallName)
	if terr != nil {
		return terr
	}

	// Close the file with error handling
	defer func() {
		cerr := tBall.Close()
		if cerr != nil && err == nil {
			err = cerr
		}
	}()

	// The tarball gets compressed with gzip
	gw := gzip.NewWriter(tBall)
	defer gw.Close()

	// Create the tar writer.
	tW := tar.NewWriter(gw)
	defer tW.Close()

	// Go through each file in the path and add it to the archive
	var i int
	var f file

	for i, f = range a.Files {
		err := a.addFile(tW, goutils.appendSlash(relPath)+f.p)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *Archiver) deletePriorBuild(p string) error {
	//delete the contents of the passed directory
	return deleteDirContent(p)
}

func formattedNow() string {
	// Time in ISO 8601 like format. The difference being the : have been
	// removed from the time.
	return time.Now().Local().Format(timeFormat)
}

func newArchiver(appendDatetime bool, datetimePrefix, datetimeFormat string) *Archiver {
	archiver :=  &Archiver{}
	archiver.appendDatetime = appendDatetime
	archiver.datetimePrefix = datetimePrefix

	// Only override the datetime format if it isn't empty
	if datetimeFormat != "" {
		archiver.datetimeFormat == datetimeformat
	}

	return archiver
}
