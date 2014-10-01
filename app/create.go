package app

import (
	"fmt"
	"os"
	"strings"

	car "github.com/mohae/carchivum"
	contour "github.com/mohae/contourp"
)

func Create(destination string, sources ...string) (string, error) {
	var err error
	var message string

	fmt.Printf("\nCreate %q from %v\n", destination, sources)

	switch contour.GetString(CfgFormat) {
	case "tar":
		message, err = createTar(destination, sources...)
	case "zip":
		message, err = createZip(destination, sources...)
	default:
		err = fmt.Errorf("%s not supported", contour.GetString(CfgFormat))
	}

	if err != nil {
		logger.Error(err)
		return "", err
	}

	return message, nil
}

func createZip(destination string, sources ...string) (string, error) {
	var err error

	logger.Debugf("Creating zip: %s from %s", destination, sources)
	zipper := car.NewZipArchive()
	zipper.Name = destination
	zipper.UseFullpath = contour.GetBool("usefullpath")
	_, err = zipper.CreateFile(destination, sources...)
	if err != nil {
		logger.Error(err)
		return "", err
	}

	return zipper.Message(), nil
}

func createTar(destination string, sources ...string) (string, error) {
	var err error

	logger.Debugf("Creating tar: %s from %s", destination, sources)
	tballer := car.NewTar()
	tballer.Name = destination
	tballer.Owner = contour.GetInt("owner")
	tballer.Group = contour.GetInt("group")
	tballer.Mode = os.FileMode(contour.GetInt64("mode"))

//	tabller.Exclude = contour.GetString("exclude")
	tballer.ExcludeAnchored = contour.GetString("exclude-anchored")
	temp := contour.GetString("exclude-ext")
	if temp != "" {
		tballer.ExcludeExt = strings.Split(temp, ",")
		tballer.ExcludeExtCount = len(tballer.ExcludeExt)
	}


	tballer.IncludeAnchored = contour.GetString("include-anchored")
	temp = contour.GetString("include-ext")
	if temp != "" {
		tballer.IncludeExt = strings.Split(temp, ",")
		tballer.IncludeExtCount = len(tballer.IncludeExt)
	}

// TODO figure out how to convert the incoming time info to time.Time
	tballer.NewerMTime = contour.GetTime("newer-mtime")
//	tballer.UseFullpath = contour.GetBool("usefullpath")
	_, err = tballer.CreateFile(destination, sources...)
	if err != nil {
		logger.Error(err)
		return "", err
	}

	return tballer.Message(), nil
}
