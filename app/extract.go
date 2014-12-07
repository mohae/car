package app

import (
	"fmt"

	car "github.com/mohae/carchivum"
)

func Extract(src, dst string) (string, error) {
	fmt.Printf("\nExtract %q to %v\n", src, dst)

	return car.Extract(src, dst)
}
