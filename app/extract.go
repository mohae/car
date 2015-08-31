package app

import (
	"fmt"
	"path"
	"strings"

	car "github.com/mohae/carchivum"
)

func Extract(src string) (string, error) {
	fmt.Printf("\nExtract %q", src)
	// check the extension: use zip for .zip and tar for everything else
	ext := strings.ToLower(path.Ext(src))
	if ext == ".zip" {
		z := car.NewZip(src)
		z.OutDir = contour.GetString("output_dir")
		return z.Extract()
	}
	t := car.NewTar(src)
	t.OutDir = contour.GetString("output_dir")
	return t.Extract()
}
