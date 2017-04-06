package main

import (
	"fmt"
	"path"
	"strings"

	car "github.com/mohae/carchivum"
	"github.com/mohae/contour"
)

func Extract(src string) (msg string, err error) {
	fmt.Printf("\nExtract %q", src)
	// check the extension: use zip for .zip and tar for everything else
	ext := strings.ToLower(path.Ext(src))
	if ext == ".zip" {
		z := car.NewZip(src)
		z.OutDir = contour.String("output_dir")
		err = z.Extract()
	} else {
		t := car.NewTar(src)
		t.OutDir = contour.String("output_dir")
		err = t.Extract()
	}
	if err != nil {
		err = fmt.Errorf("extract %s: %s", src, err)
		return "", nil
	}
	if contour.String("output_dir") == "" {
		msg = fmt.Sprintf("%s extracted", src)
	} else {
		msg = fmt.Sprintf("%s extracted to %s", src, contour.String("output_dir"))
	}
	return msg, nil
}
