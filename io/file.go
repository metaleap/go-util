package io

import (
	"archive/zip"
	"encoding/binary"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	strutil "github.com/go-ngine/go-util/str"
)

func DirExists (path string) bool {
	var stat, err = os.Stat(path)
	if (err == nil) && (stat != nil) {
		return stat.IsDir()
	}
	return false
}

func ExtractZipFile (zipFilePath, targetDirPath string, deleteZipFile bool, fileNamesPrefix string, fileNamesToExtract ... string) error {
	var fnames []string = nil
	var fnprefix string = ""
	var efile *os.File
	var zfile *zip.File
	var zfileReader io.ReadCloser
	var unzip, err = zip.OpenReader(zipFilePath)
	if unzip != nil {
		if (err == nil) && (unzip.File != nil) {
			if (fileNamesToExtract != nil) && (len(fileNamesToExtract) > 0) {
				fnames = fileNamesToExtract
				for i, fn := range fnames {
					if fn[0:len(fileNamesPrefix)] == fileNamesPrefix {
						fnames[i] = fn[len(fileNamesPrefix):]
						fnprefix = fileNamesPrefix
					}
				}
			}
			for _, zfile = range unzip.File {
				if (fnames == nil) || (strutil.InSliceAt(fnames, zfile.FileHeader.Name) >= 0) {
					zfileReader, err = zfile.Open()
					if zfileReader != nil {
						if (err == nil) {
							efile, err = os.Create(filepath.Join(targetDirPath, fnprefix + zfile.FileHeader.Name))
							if efile != nil {
								if (err == nil) {
									_, err = io.Copy(efile, zfileReader)
								}
								efile.Close()
							}
						}
						zfileReader.Close()
					}
				}
				if (err != nil) {
					break
				}
			}
		}
		unzip.Close()
		if deleteZipFile && (err == nil) {
			err = os.Remove(zipFilePath)
		}
	}
	return err
}

func FileExists (path string) bool {
	var stat, err = os.Stat(path)
	if (err == nil) && (stat != nil) {
		return !stat.IsDir()
	}
	return false
}

func FileExistsPath (dirPath string, fileBaseName string, fileExts []string, tryLower bool, tryUpper bool) (string, time.Time, int64) {
	var stat os.FileInfo
	var err error
	var fext, fpath string
	for _, fext = range fileExts {
		fpath = filepath.Join(dirPath, fileBaseName + fext)
		if stat, err = os.Stat(fpath); err != nil {
			if tryUpper {
				fpath = filepath.Join(dirPath, strings.ToUpper(fileBaseName) + fext)
				stat, err = os.Stat(fpath)
			}
			if (err != nil) && tryLower {
				fpath = filepath.Join(dirPath, strings.ToLower(fileBaseName) + fext)
				stat, err = os.Stat(fpath)
			}
		}
		if (err == nil) && (stat != nil) && !stat.IsDir() {
			return fpath, stat.ModTime(), stat.Size()
		}
	}
	return "", time.Time {}, 0
}

func ReadFromBinary (readSeeker io.ReadSeeker, offset int64, byteOrder binary.ByteOrder, ptr interface{}) bool {
	var o, err = readSeeker.Seek(offset, 0)
	if (o != offset) || (err != nil) {
		return false
	}
	if err = binary.Read(readSeeker, byteOrder, ptr); err != nil {
		return false
	}
	return true
}

func ReadTextFile (filePath string, panicOnError bool, defVal string) string {
	var file, err = os.Open(filePath)
	var bytes []byte
	if err == nil {
		bytes, err = ioutil.ReadAll(file)
		file.Close()
		if (err == nil) {
			return string(bytes)
		}
	}
	if panicOnError && (err != nil) {
		panic(err)
	}
	return defVal
}

func WalkDirectory (dirPath, fileSuffix string, fileFunc func(string), recurseSubDirs bool) error {
	var fileInfos, err = ioutil.ReadDir(dirPath)
	if err == nil {
		for _, fi := range fileInfos {
			if !fi.IsDir() {
				if (len(fileSuffix) == 0) || strings.HasSuffix(fi.Name(), fileSuffix) {
					fileFunc(filepath.Join(dirPath, fi.Name()))
				}
			} else if recurseSubDirs {
				if err = WalkDirectory(filepath.Join(dirPath, fi.Name()), fileSuffix, fileFunc, recurseSubDirs); err != nil {
					break
				}
			}
		}
	}
	return err
}
