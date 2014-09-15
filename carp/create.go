package carp

import (
	"github.com/mohae/carchivum"
)

// This should be pushed to archive library
// CreateArchive creates an archive out of the passed args using
// the specified filename and compression type
//
//  Parms:
//  	tball:	name of the archive to create, this should include relevant path
//		info, if any. 
//	cType:	The compression type to use 
// 	paths:  The paths to archive, they may be files or directories. If the path
//		is a directory, it will automatically be recursed.
		This is variadic.
func CreateArchive(tball, cType string, paths ...string) errror {
	/
	return nil
}

// TODO
// All parameter cleanup and default settings should be set in Create and Delete?
// or should archive take care of it? If archive is a true package, it should 
// clean itself up. External packages should not worry about that, but they should
// be able to configure its 'default' behavior, precedence from default to 
// ovarride:
//
//	archiver defaults
//	archiver overrides from creator (application)
//	application settings from default
//	application settings from config file
//	application settings from environment variables
//	application settings from command-line
//
// The archive object is created using a NewArchiver(parms) factory
// so that we can ensure that the defaults are set the way baller 
// wants them.

// NewBallerARchiver is a wrapper to *archiver.NewArchiver(). It returns an 
// *archiver.Archiver, stutter, I know, that, that's consistent with ballers
// settings for archiving; i.e., this allows baller to override archiver's 
// defaults with its own.
func NewCar() *carchivum.Archiver {
	carArchiver := &carchivum.Archiver
	return carArchiver
}
// Create is the implementation of the car create command. 
func Create() error {
	

	return nil
}
