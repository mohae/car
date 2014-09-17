package app

import (
	"fmt"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
	arch "github.com/mohae/carchivum"
	"github.com/mohae/contour"
)

func Create(destination string, sources ...string) (string, error) {
	t0 := time.Now()
	var err error
	var message string

	fmt.Println(EnvLog, strconv.FormatBool(contour.GetBool(EnvLog)))
	fmt.Println(EnvLogLevel, contour.GetString(EnvLogLevel))
	fmt.Println(EnvStdout, strconv.FormatBool(contour.GetBool(EnvStdout)))
	fmt.Println(EnvStdoutLevel, contour.GetString(EnvStdoutLevel))
	fmt.Println(EnvVerbose, strconv.FormatBool(contour.GetBool(EnvVerbose)))
	fmt.Println(EnvArchiveFormat, contour.GetString(EnvArchiveFormat))
	fmt.Println(EnvCompressionType, contour.GetString(EnvCompressionType))

	log.Error(contour.GetString(EnvCompressionType))
	log.Info(strconv.FormatBool(contour.GetBool(EnvLog)))

	archive := arch.NewArchive()
	err = archive.SetFormat(contour.GetString(EnvArchiveFormat))
	if err != nil {
		log.WithFields(log.Fields{
			"ArchiveFormat": contour.GetString(EnvArchiveFormat),
		}).Error(err)
		return message, err
	}

	err = archive.SetCompressionType(contour.GetString(EnvCompressionType))
	if err != nil {
		log.WithFields(log.Fields{
			"CompressionType": contour.GetString(EnvCompressionType),
		}).Error(err)
		return message, err
	}

	err = archive.SetDateFormat(contour.GetString("dateformat"))
	if err != nil {
		log.WithFields(log.Fields{
			"DateFormat": contour.GetString("dateformat"),
		}).Error(err)
		return message, err
	}

	message, err = archive.Create(destination, sources...)
	if err != nil {
		log.WithFields(log.Fields{
			"Archive": destination,
			"Sources": sources,
		}).Error(err)
		return message, err
	}

	// message = fmr.Sprintf("%s created from: %s in %d seconds\n", destination, sources, archive.Time())
	Δt := float64(time.Now().Sub(t0)) / 1e9
	elapsed := fmt.Sprintf("Create process complete: %.4f seconds", Δt)
	message = elapsed

	log.WithFields(log.Fields{
		"operations": "Create",
		"sources":    sources,
		"duration":   Δt,
	}).Debugf("%s created", destination)

	return message, nil
}
