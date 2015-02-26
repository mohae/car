package app

import (
	"fmt"

	car "github.com/mohae/carchivum"
)

func Extract(src, dst string) (string, error) {
	fmt.Printf("\nExtract %q to %v\n", src, dst)
	err := car.Extract(src, dst)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s extracted to %s", src, dst), nil
}
