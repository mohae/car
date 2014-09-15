package command

import (
	"github.com/mohae/contour"
)

func initConfig() {
	contour.RegisterStringFlag("format", "tar", "f")
	contour.RegisterStringFlag("type", "gzip", "t")
	contour.RegisterBoolFlag("verbose", false, "v")
	//	registerCommandSettings()
}
