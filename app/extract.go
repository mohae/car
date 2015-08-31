package app

import (
	"fmt"

	car "github.com/mohae/carchivum"
)

func Extract(src string) (string, error) {
	fmt.Printf("\nExtract %q", src)
	err := car.Extract(src)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s extracted", src), nil
}
