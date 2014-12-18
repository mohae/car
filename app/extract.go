package app

import (
	"fmt"
	"strconv"

	car "github.com/mohae/carchivum"
	"github.com/mohae/contour"
)

func Extract(src, dst string) (string, error) {
	fmt.Printf("\nExtract %q to %v\n", src, dst)
	b, _ := strconv.ParseBool(contour.GetBool("createdir"))
	car.CreateDir = b
	return "", nil
	//	return car.Extract(src, dst)
}
