package app

import (
	_ "fmt"
	_ "strings"
	"time"

	log "github.com/Sirupsen/logrus"
	_ "github.com/mohae/carchivum"
	_ "github.com/mohae/contour"
)

func Create(destination string, sources ...string) (string, error) {
	t0 := time.Now()
	/*
		archive := arch.NewArchive()

		arr = archiver.SetCompression(contour.GetString("type"))
		if err != nil {
			log.WithFields(log.Fields{
				"CompressionType":  contour.GetString("type")
			}).Error(err)
			return "", err
		}

		err = archiver.SetDateFormat(contour.GetString("dateformat"))
		if err != nil {
			log.WithFields{
				"DateFormat": contour.GetString("dateformat")
			}).Error(err)
			return "", err
		}

		err := archive.Create("archive{{ .Datetime }}", ".")
		if err != nil {
			log.WithFields{
				"Archive": destination,
				"Sources": source,
			}).Error((err)
			return "", err
		}
	*/
	// message = fmr.Sprintf("%s created from: %s in %d seconds\n", destination, sources, archiver.Time())
	Δt := float64(time.Now().Sub(t0)) / 1e9
	message := "tmp message"

	log.WithFields(log.Fields{
		"operations": "Create",
		"sources":    sources,
		"duration":   Δt,
	}).Debugf("%s created", destination)

	return message, nil
}
