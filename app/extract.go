package app

import (
	"fmt"

	car "github.com/mohae/carchivum"
	"github.com/mohae/contour"
)

func Extract(src, dst string) (string, error) {
	fmt.Printf("\nExtract %q to %v\n", src, dst)
	car.CreateDir = contour.GetBool("createdir")
	return "", nil
	//	return car.Extract(src, dst)
}
