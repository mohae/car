package command

import (
	"github.com/mohae/contour"
)

func initConfig() {
	contour.RegisterFlagString("format", "tar", "f")
	contour.RegisterFlagString("type", "gzip", "t")
	contour.RegisterFlagBool("verbose", false, "v")
	//	registerCommandSettings()
}
