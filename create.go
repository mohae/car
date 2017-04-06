package main

import (
	"fmt"
	_ "strconv"
	"strings"

	car "github.com/mohae/carchivum"
	"github.com/mohae/contour"
	log "github.com/mohae/ezlog"
	magicnum "github.com/mohae/magicnum/compress"
)

func Create(dst string, sources ...string) (string, error) {
	switch contour.String(Format) {
	case "zip":
		return createZip(dst, sources...)
	case "tar":
		return createTar(dst, sources...)
	default:
		return "", fmt.Errorf("create %s: unknown archive format: %s", dst, contour.String(Format))
	}
}

func createZip(dst string, sources ...string) (string, error) {
	zipper := car.NewZip(dst)
	zipper.UseFullpath = contour.Bool("abspath")
	_, err := zipper.Create(sources...)
	if err != nil {
		log.Error(err)
		return "", err
	}
	return zipper.Message(), nil
}

func createTar(dst string, sources ...string) (string, error) {
	tballer := car.NewTar(dst)
	t := contour.String("type")
	if t != "" {
		f := magicnum.ParseFormat(t)
		if f == magicnum.Unknown {
			err := fmt.Errorf("unknown compression format: %s", t)
			log.Error(err)
			return "", err
		}
		if !car.IsSupported(f) {
			err := fmt.Errorf("unsupported compressiong format: %s is not supported", f)
			log.Error(err)
			return "", err
		}
		tballer.Format = f
	}
	//	tabller.Exclude = contour.GetString("exclude")
	tballer.ExcludeAnchored = contour.String("exclude_anchored")
	temp := contour.String("exclude_ext")
	if temp != "" {
		tballer.ExcludeExt = strings.Split(temp, ",")
		tballer.ExcludeExtCount = len(tballer.ExcludeExt)
	}
	tballer.IncludeAnchored = contour.String("include_anchored")
	temp = contour.String("include_ext")
	if temp != "" {
		tballer.IncludeExt = strings.Split(temp, ",")
		tballer.IncludeExtCount = len(tballer.IncludeExt)
	}
	// TODO figure out how to convert the incoming time info to time.Time
	// tballer.NewerMTime = contour.GetTime("newer-mtime")
	//	tballer.UseFullpath = contour.GetBool("usefullpath")
	_, err := tballer.Create(sources...)
	if err != nil {
		log.Error(err)
		return "", err
	}
	return tballer.Message(), nil
}
