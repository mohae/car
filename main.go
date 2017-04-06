// Copyright © <year>, <author> under the <license> license.
//	All	rights reserved. For license information see the LICENSE file.

// Car creates compressed archive; the uncompressed archive is a tar. If a car
// archive is compressed using a compression format that tar supports, it can
// be uncompressed using tar. Car can uncompress compressed tarballs that are
// compressed using a compression format that car supports.
//
// Car can be run on Windows or Linux, depending on either the system on which
// it was compiled or the GOOS that was targeted during compile. It should work
// on OSX, but it hasn't been used or tested on it.
package main

import (
	"os"

	"github.com/mohae/contour"
)

var (
	logF     *os.File // logfile handle for close; this will be nil if output is stderr
	closeLog bool
)

func init() {
	contour.RegisterStringFlag("logfile", "l", "stderr", "stderr", "output destination for logs; if set to empty string log output will be discarded")
	contour.RegisterStringFlag("level", "v", "error", "error", "log level")
}

func main() {
	os.Exit(carMain())
}
