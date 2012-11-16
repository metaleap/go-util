package io

import (
	"compress/gzip"
	"encoding/gob"
	"os"

	util "github.com/metaleap/go-util"
)

func CreateGobsFile (targetFilePath string, recs []interface{}, getRecPtr util.AnyToAny, gzipped bool) (err error) {
	var file *os.File
	var gobber *gob.Encoder
	var gzipper *gzip.Writer
	if file, err = os.Create(targetFilePath); file != nil {
		defer file.Close()
	}
	if err != nil { return }
	if gzipped {
		if gzipper, err = gzip.NewWriterLevel(file, gzip.BestCompression); gzipper != nil {
			defer gzipper.Close()
			gobber = gob.NewEncoder(gzipper)
		}
		if err != nil { return }
	} else {
		gobber = gob.NewEncoder(file)
	}
	for _, rec := range recs {
		if err = gobber.Encode(util.PtrVal(getRecPtr(rec))); err != nil { return }
	}
	return
}
