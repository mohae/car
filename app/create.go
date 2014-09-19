package app

import (
	"fmt"
	"strconv"
	"time"

	arch "github.com/mohae/carchivum"
	"github.com/mohae/contour"
)

func Create(destination string, sources ...string) (string, error) {
	t0 := time.Now()
	var err error
	var message string

	fmt.Println(EnvLog, strconv.FormatBool(contour.GetBool(EnvLog)))
	fmt.Println(EnvVerbose, strconv.FormatBool(contour.GetBool(EnvVerbose)))
	fmt.Println(EnvArchiveFormat, contour.GetString(EnvArchiveFormat))
	fmt.Println(EnvCompressionType, contour.GetString(EnvCompressionType))

	logger.Debugf("Creating archive %s", destination)

	logger.Infof("Log: %s", strconv.FormatBool(contour.GetBool(EnvLog)))

	fmt.Printf("\nCreate %q from %v\n", destination, sources)
	archive := arch.NewArchive()
	err = archive.SetFormat(contour.GetString(EnvArchiveFormat))
	if err != nil {
		logger.Error(err)
		return message, err
	}

	err = archive.SetCompressionType(contour.GetString(EnvCompressionType))
	if err != nil {
		logger.Error(err)
		return message, err
	}

	archive.SetDateFormat(contour.GetString("dateformat"))
	logger.Debugf("dateformat: %s\n", contour.GetString("dateformat"))
	
	message, err = archive.Create(destination, sources...)
	if err != nil {
		logger.Error(err)
		return message, err
	}

	// message = fmr.Sprintf("%s created from: %s in %d seconds\n", destination, sources, archive.Time())
	Δt := float64(time.Now().Sub(t0)) / 1e9
	elapsed := fmt.Sprintf("Create process complete: %.4f seconds", Δt)
	message = elapsed
	logger.Debugf("%s created in %.4f seconds", destination, Δt)

	return message, nil
}
