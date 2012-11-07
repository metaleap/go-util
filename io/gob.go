package io

import (
	"compress/gzip"
	"encoding/gob"
	"os"

	util "github.com/metaleap/go-util"
)

func CreateGobsFile (targetFilePath string, recs *[]interface{}, getRecPtr util.AnyToAny, gzipped bool) {
	var file, err = os.Create(targetFilePath)
	var gobber *gob.Encoder
	var gzipper *gzip.Writer
	if file != nil {
		defer file.Close()
	}
	if err != nil {
		panic(err)
	}
	if gzipped {
		if gzipper, err = gzip.NewWriterLevel(file, gzip.BestCompression); gzipper != nil {
			defer gzipper.Close()
			gobber = gob.NewEncoder(gzipper)
		}
		if err != nil {
			panic(err)
		}
	} else {
		gobber = gob.NewEncoder(file)
	}
	for _, rec := range *recs {
		if err = gobber.Encode(util.PtrVal(getRecPtr(rec))); err != nil {
			panic(err)
		}
	}
}
